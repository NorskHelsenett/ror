package dex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/cmd/auth/msauthconnections"
	"github.com/NorskHelsenett/ror/internal/models/vaultmodels"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"

	"github.com/dexidp/dex/api/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	grpcCreds "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func HandleClusterCreated(ctx context.Context, event *amqp.Delivery) error {
	var message messagebuscontracts.ClusterCreatedEvent
	err := json.Unmarshal(event.Body, &message)
	if err != nil {
		rlog.Error("could not unmarshal message", err)
		return fmt.Errorf("could not convert message to json")
	}

	dexHost := viper.GetString(configconsts.DEX_HOST)
	dexGrpcPort := viper.GetString(configconsts.DEX_GRPC_PORT)

	hostandPort := fmt.Sprintf("%s:%s", dexHost, dexGrpcPort)
	if len(dexGrpcPort) == 0 {
		hostandPort = dexHost
	}
	rlog.Info("dex host and port", rlog.String("host", dexHost), rlog.String("port", dexGrpcPort))
	caPath := viper.GetString(configconsts.DEX_CERT_FILEPATH)

	dexClient, err := newDexClient(hostandPort, caPath)
	if err != nil {
		return fmt.Errorf("could not create a dex client, host and port: %s", hostandPort)
	}

	secret := stringhelper.RandomString(16, stringhelper.StringTypeAlphaNum)
	if len(secret) < 1 {
		err := errors.New("could not create clientsecret, secret is empty")
		rlog.Error("error creating client secret", err)
		return err
	}

	clusterNameAndWorkspaceName := fmt.Sprintf("%s.%s", message.ClusterName, message.WorkspaceName)
	req := &api.CreateClientReq{
		Client: &api.Client{
			Id:     clusterNameAndWorkspaceName,
			Name:   clusterNameAndWorkspaceName,
			Secret: secret,
			RedirectUris: []string{
				fmt.Sprintf("https://argo.%s.sky.nhn.no/auth/callback", clusterNameAndWorkspaceName),
				fmt.Sprintf("https://grafana.%s.sky.nhn.no/login/generic_oauth", clusterNameAndWorkspaceName),
			},
		},
	}

	rlog.Debug("request to dex", rlog.Any("request", req))
	response, err := dexClient.CreateClient(ctx, req)
	if err != nil {
		rlog.Error("failed creating oauth2 client", err)
		return err
	}

	alreadyExists := response.AlreadyExists
	// what is going on here??
	rlog.Info(fmt.Sprintf("client secret created in dex (%s): %t, already exist: %t", message.ClusterId, !alreadyExists, alreadyExists))

	secretPath := fmt.Sprintf("%s%s", viper.GetString(configconsts.DEX_VAULT_PATH), message.ClusterId)

	existingSecret, err := msauthconnections.VaultClient.GetSecret(secretPath)
	if err != nil && !strings.Contains(err.Error(), "404") {
		fmt.Println(err.Error())
		rlog.Error("error getting existing secret", err)
		return err
	}

	secretjson, err := json.Marshal(existingSecret)
	if err != nil {
		rlog.Error("error marshaling existing secret", err)
		return err
	}

	var secretModel vaultmodels.VaultDexModel
	err = json.Unmarshal(secretjson, &secretModel)
	if err != nil {
		rlog.Error("error unmarshaling secret", err)
		return err
	}

	if alreadyExists && len(secretModel.Data.DexSecret) > 0 {
		return nil
	}

	secretModel.Data.DexSecret = secret
	secretByteArray, err := json.Marshal(secretModel)
	if err != nil {
		rlog.Error("error marshaling secret", err)
		return err
	}

	resultOk, err := msauthconnections.VaultClient.SetSecret(secretPath, secretByteArray)
	if err != nil {
		rlog.Error("error setting secret", err)
		return fmt.Errorf("failed creating oauth2 client: %v", err)
	}

	if resultOk {
		return nil
	}

	rlog.Error("error setting secret, final call", err)
	return errors.New("could not set secret for clusterid")
}

func newDexClient(hostAndPort string, caPath string) (api.DexClient, error) {
	dexTls := viper.GetBool(configconsts.DEX_TLS)
	var dialOptions grpc.DialOption
	if dexTls {
		creds, err := grpcCreds.NewClientTLSFromFile(caPath, "")
		if err != nil {
			rlog.Error("Could not create new client tls from file", err)
			return nil, fmt.Errorf("load dex cert: %v", err)
		}
		dialOptions = grpc.WithTransportCredentials(creds)
	} else {
		dialOptions = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	conn, err := grpc.Dial(hostAndPort,
		dialOptions,
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		rlog.Error("Could not dial grpc endpoint", err)
		return nil, fmt.Errorf("dial: %v", err)
	}
	return api.NewDexClient(conn), nil
}
