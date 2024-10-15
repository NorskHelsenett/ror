package v1alpha1

import (
	"context"
	"github.com/NorskHelsenett/ror/cmd/operator/api/v1alpha1"
	appsv1alpha1 "github.com/NorskHelsenett/ror/cmd/operator/api/v1alpha1"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

var (
	resourceName string = "tasks"
)

func (c *TasksConfigV1Alpha1Client) TaskConfigs(ctx context.Context, namespace string) TasksConfigInterface {
	return &taskConfigclient{
		client: c.restClient,
		ns:     namespace,
		ctx:    ctx,
	}
}

type TasksConfigV1Alpha1Client struct {
	restClient rest.Interface
}

type TasksConfigInterface interface {
	Get(name string) (*appsv1alpha1.Task, error)
	List(meta_v1.ListOptions) (*appsv1alpha1.TaskList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)

	Create(obj *appsv1alpha1.Task) (*appsv1alpha1.Task, error)
	Update(obj *appsv1alpha1.Task) (*appsv1alpha1.Task, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
}

type taskConfigclient struct {
	client rest.Interface
	ns     string
	ctx    context.Context
}

func (c *taskConfigclient) Create(task *v1alpha1.Task) (*v1alpha1.Task, error) {
	result := v1alpha1.Task{}
	request := c.client.
		Post().
		Namespace(c.ns).
		Resource(resourceName).
		Body(task)
	response := request.Do(c.ctx)

	err := response.Into(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *taskConfigclient) Update(obj *appsv1alpha1.Task) (*appsv1alpha1.Task, error) {
	result := &appsv1alpha1.Task{}
	request := c.client.Put().
		Namespace(c.ns).
		Resource(resourceName).
		Body(obj).
		Name(obj.Name)
	response := request.Do(c.ctx)
	err := response.Into(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *taskConfigclient) Delete(name string, options *meta_v1.DeleteOptions) error {
	request := c.client.Delete().
		Namespace(c.ns).Resource(resourceName).
		Name(name).
		Body(options)
	err := request.Do(c.ctx).Error()
	return err
}

func (c *taskConfigclient) Get(name string) (*appsv1alpha1.Task, error) {
	result := &appsv1alpha1.Task{}
	request := c.client.Get().
		Namespace(c.ns).
		Resource(resourceName).
		Name(name)
	response := request.Do(c.ctx)
	err := response.Into(result)
	return result, err
}

func (c *taskConfigclient) List(opts meta_v1.ListOptions) (*v1alpha1.TaskList, error) {
	result := v1alpha1.TaskList{}
	err := c.client.
		Get().
		Namespace(c.ns).
		Resource(resourceName).
		Do(c.ctx).
		Into(&result)

	return &result, err
}

func (c *taskConfigclient) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.
		Get().
		Namespace(c.ns).
		Resource(resourceName).
		//VersionedParams(&opts, scheme.ParameterCodec).
		Watch(c.ctx)
}
