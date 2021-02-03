// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"github.com/tigera/api/pkg/client/clientset_generated/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type ProjectcalicoInterface interface {
	RESTClient() rest.Interface
	GlobalAlertsGetter
	GlobalAlertTemplatesGetter
	GlobalReportTypesGetter
	LicenseKeysGetter
	ManagedClustersGetter
}

// ProjectcalicoClient is used to interact with features provided by the projectcalico.org group.
type ProjectcalicoClient struct {
	restClient rest.Interface
}

func (c *ProjectcalicoClient) GlobalAlerts() GlobalAlertInterface {
	return newGlobalAlerts(c)
}

func (c *ProjectcalicoClient) GlobalAlertTemplates() GlobalAlertTemplateInterface {
	return newGlobalAlertTemplates(c)
}

func (c *ProjectcalicoClient) GlobalReportTypes(namespace string) GlobalReportTypeInterface {
	return newGlobalReportTypes(c, namespace)
}

func (c *ProjectcalicoClient) LicenseKeys(namespace string) LicenseKeyInterface {
	return newLicenseKeys(c, namespace)
}

func (c *ProjectcalicoClient) ManagedClusters() ManagedClusterInterface {
	return newManagedClusters(c)
}

// NewForConfig creates a new ProjectcalicoClient for the given config.
func NewForConfig(c *rest.Config) (*ProjectcalicoClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ProjectcalicoClient{client}, nil
}

// NewForConfigOrDie creates a new ProjectcalicoClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ProjectcalicoClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ProjectcalicoClient for the given RESTClient.
func New(c rest.Interface) *ProjectcalicoClient {
	return &ProjectcalicoClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("projectcalico.org")[0].Group {
		gv := scheme.Scheme.PrioritizedVersionsForGroup("projectcalico.org")[0]
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs

	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ProjectcalicoClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
