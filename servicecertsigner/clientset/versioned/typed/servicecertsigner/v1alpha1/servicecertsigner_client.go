// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/openshift/api/servicecertsigner/v1alpha1"
	"github.com/openshift/client-go/servicecertsigner/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type ServicecertsignerV1alpha1Interface interface {
	RESTClient() rest.Interface
	ServiceCertSignerOperatorConfigsGetter
}

// ServicecertsignerV1alpha1Client is used to interact with features provided by the servicecertsigner.config.openshift.io group.
type ServicecertsignerV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ServicecertsignerV1alpha1Client) ServiceCertSignerOperatorConfigs() ServiceCertSignerOperatorConfigInterface {
	return newServiceCertSignerOperatorConfigs(c)
}

// NewForConfig creates a new ServicecertsignerV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ServicecertsignerV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ServicecertsignerV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ServicecertsignerV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ServicecertsignerV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ServicecertsignerV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ServicecertsignerV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ServicecertsignerV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ServicecertsignerV1alpha1Client {
	return &ServicecertsignerV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ServicecertsignerV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
