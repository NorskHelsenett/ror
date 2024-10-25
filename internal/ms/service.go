package ms

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"
	"github.com/NorskHelsenett/ror/pkg/telemetry/trace"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

// HandlerFunc defines the signature of the callback function called whenever there is a new delivery on the rabbitmq
// queue.
type HandlerFunc func(ctx *Context) error

type listener struct {
	Workqueue string
	Autoack   bool
	Handler   HandlerFunc
}

// Service is a containing struct containing our RabbitMQ credentials and connections. It also manages the internal
// listener on which to perform a callback when a queue gets a new delivery.
// It also keeps track of other practical things such as our token for communicating with vault.
type Service struct {
	Role string

	RmqCredentials *rabbitMQCredentials

	RmqChannel    *amqp.Channel
	RmqConnection *amqp.Connection

	Listeners map[string]*listener

	Wg *sync.WaitGroup
}

// New initializes a new service for us performing our RabbitMQ connection, vault communication to extract our token.
func New(role string, vaultClient *vaultclient.VaultClient) (*Service, error) {
	service := new(Service)

	service.Role = role

	service.Listeners = make(map[string]*listener)

	rmqCreds, err := newGetRabbitMQCreds(role, vaultClient)
	if err != nil {
		return nil, err
	}
	service.RmqCredentials = rmqCreds

	if err := service.newRabbitMQConnection(); err != nil {
		return nil, err
	}

	service.Wg = new(sync.WaitGroup)

	return service, nil
}
func (s *Service) setupListener(rabbitlistener *listener) error {
	q, err := s.RmqChannel.QueueDeclare(
		rabbitlistener.Workqueue, // name
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	if err != nil {
		return fmt.Errorf("could not declare queue: %s: %v", rabbitlistener.Workqueue, err)
	}

	messages, err := s.RmqChannel.Consume(
		q.Name,                 // queue
		"",                     // consumer
		rabbitlistener.Autoack, // auto-ack
		false,                  // exclusive
		false,                  // no-local
		false,                  // no-wait
		nil,                    // args
	)
	if err != nil {
		return fmt.Errorf("could not register a consumer on workqueue: %s: %v", rabbitlistener.Workqueue, err)
	}

	s.Wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for delivery := range messages {
			// Loop variables captured by 'func' literals in 'go' statements might have unexpected values
			clone := delivery
			go func() {
				tr := otel.Tracer("event")
				c, span := tr.Start(trace.ExtractAMQPHeaders(context.Background(), clone.Headers), rabbitlistener.Workqueue)
				defer span.End()

				ctx, err := newContext(c, clone, s.RmqChannel, s.Role)
				if err != nil {
					rlog.Error("could not initialize new ms.Context", err)
					return
				}

				if err := rabbitlistener.Handler(ctx); err != nil {
					rlog.Error("could not handle delivery on queue", err, rlog.String("queue", rabbitlistener.Workqueue))
				}
			}()
		}

	}(s.Wg)

	return nil
}

func (s *Service) newRabbitMQConnection() error {
	connectionString := s.RmqCredentials.String()

	c := make(chan *amqp.Error)
	go func() {
		err := <-c
		rlog.Error("Reconnect Rabbitmq", err)

		for _, listener := range s.Listeners {
			if err := s.setupListener(listener); err != nil {
				rlog.Error("could not setup listener", err, rlog.String("listener", listener.Workqueue))
			}
		}

		time.Sleep(5 * time.Second)
	}()

	connection, err := amqp.Dial(connectionString)
	if err != nil {
		rlog.Error("cannot connect", err, rlog.String("Host", s.RmqCredentials.Host))
		return fmt.Errorf("cannot connect to rabbitmq url: %s", s.RmqCredentials.Host)
	}

	s.RmqConnection = connection
	s.RmqConnection.NotifyClose(c)

	if err != nil {
		return fmt.Errorf("could not connect with rabbitmq: %v", err)
	}

	channel, err := connection.Channel()
	if err != nil {
		panic("Failed to open rabbitmq channel")
	}

	s.RmqChannel = channel

	rlog.Info("Connected to RabbitMQ")

	return nil
}

// Listen initializes a RabbitMQ queue and listens for new deliveries on said queue. When a new delivery is discovered
// it calls the callback function. This is similar to how http frameworks like Gin And Echo do it.
func (s *Service) Listen(workqueue string, autoack bool, handler HandlerFunc) error {
	l := new(listener)

	l.Autoack = autoack
	l.Workqueue = workqueue
	l.Handler = handler

	if err := s.setupListener(l); err != nil {
		return err
	}

	s.Listeners[workqueue] = l

	return nil
}

func (s *Service) healthEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Healthy"))
}

func (s *Service) Wait() error {
	s.Wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		stop := make(chan struct{})
		trace.ConnectTracer(stop, s.Role, viper.GetString(configconsts.OPENTELEMETRY_COLLECTOR_ENDPOINT))
	}(s.Wg)

	s.Wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		serveAddress := ":8080"
		if viper.GetString(configconsts.MS_HTTP_BIND_PORT) != "" {
			serveAddress = fmt.Sprintf(":%s", viper.GetString(configconsts.MS_HTTP_BIND_PORT))
		}

		if viper.GetString(configconsts.MS_HTTP_PORT) != "" {
			serveAddress = ":" + viper.GetString(configconsts.MS_HTTP_PORT)
		}

		httpServer := &http.Server{
			Addr:              serveAddress,
			Handler:           otelhttp.NewHandler(http.HandlerFunc(s.healthEndpoint), "/health"),
			ReadHeaderTimeout: 0,
		}

		if err := httpServer.ListenAndServe(); err != nil {
			rlog.Fatal("could not start health server", err)
		}
	}(s.Wg)

	s.Wg.Wait()

	return nil
}
