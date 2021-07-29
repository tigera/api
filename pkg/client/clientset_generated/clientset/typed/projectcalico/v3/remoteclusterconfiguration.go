// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"time"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	scheme "github.com/tigera/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RemoteClusterConfigurationsGetter has a method to return a RemoteClusterConfigurationInterface.
// A group's client should implement this interface.
type RemoteClusterConfigurationsGetter interface {
	RemoteClusterConfigurations() RemoteClusterConfigurationInterface
}

// RemoteClusterConfigurationInterface has methods to work with RemoteClusterConfiguration resources.
type RemoteClusterConfigurationInterface interface {
	Create(*v3.RemoteClusterConfiguration) (*v3.RemoteClusterConfiguration, error)
	Update(*v3.RemoteClusterConfiguration) (*v3.RemoteClusterConfiguration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v3.RemoteClusterConfiguration, error)
	List(opts v1.ListOptions) (*v3.RemoteClusterConfigurationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.RemoteClusterConfiguration, err error)
	RemoteClusterConfigurationExpansion
}

// remoteClusterConfigurations implements RemoteClusterConfigurationInterface
type remoteClusterConfigurations struct {
	client rest.Interface
}

// newRemoteClusterConfigurations returns a RemoteClusterConfigurations
func newRemoteClusterConfigurations(c *ProjectcalicoV3Client) *remoteClusterConfigurations {
	return &remoteClusterConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the remoteClusterConfiguration, and returns the corresponding remoteClusterConfiguration object, and an error if there is any.
func (c *remoteClusterConfigurations) Get(name string, options v1.GetOptions) (result *v3.RemoteClusterConfiguration, err error) {
	result = &v3.RemoteClusterConfiguration{}
	err = c.client.Get().
		Resource("remoteclusterconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RemoteClusterConfigurations that match those selectors.
func (c *remoteClusterConfigurations) List(opts v1.ListOptions) (result *v3.RemoteClusterConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.RemoteClusterConfigurationList{}
	err = c.client.Get().
		Resource("remoteclusterconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested remoteClusterConfigurations.
func (c *remoteClusterConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("remoteclusterconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a remoteClusterConfiguration and creates it.  Returns the server's representation of the remoteClusterConfiguration, and an error, if there is any.
func (c *remoteClusterConfigurations) Create(remoteClusterConfiguration *v3.RemoteClusterConfiguration) (result *v3.RemoteClusterConfiguration, err error) {
	result = &v3.RemoteClusterConfiguration{}
	err = c.client.Post().
		Resource("remoteclusterconfigurations").
		Body(remoteClusterConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a remoteClusterConfiguration and updates it. Returns the server's representation of the remoteClusterConfiguration, and an error, if there is any.
func (c *remoteClusterConfigurations) Update(remoteClusterConfiguration *v3.RemoteClusterConfiguration) (result *v3.RemoteClusterConfiguration, err error) {
	result = &v3.RemoteClusterConfiguration{}
	err = c.client.Put().
		Resource("remoteclusterconfigurations").
		Name(remoteClusterConfiguration.Name).
		Body(remoteClusterConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the remoteClusterConfiguration and deletes it. Returns an error if one occurs.
func (c *remoteClusterConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("remoteclusterconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *remoteClusterConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("remoteclusterconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched remoteClusterConfiguration.
func (c *remoteClusterConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.RemoteClusterConfiguration, err error) {
	result = &v3.RemoteClusterConfiguration{}
	err = c.client.Patch(pt).
		Resource("remoteclusterconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
