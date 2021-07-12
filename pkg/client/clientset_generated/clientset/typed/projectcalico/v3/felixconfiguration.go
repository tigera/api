// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	scheme "github.com/tigera/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FelixConfigurationsGetter has a method to return a FelixConfigurationInterface.
// A group's client should implement this interface.
type FelixConfigurationsGetter interface {
	FelixConfigurations() FelixConfigurationInterface
}

// FelixConfigurationInterface has methods to work with FelixConfiguration resources.
type FelixConfigurationInterface interface {
	Create(ctx context.Context, felixConfiguration *v3.FelixConfiguration, opts v1.CreateOptions) (*v3.FelixConfiguration, error)
	Update(ctx context.Context, felixConfiguration *v3.FelixConfiguration, opts v1.UpdateOptions) (*v3.FelixConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.FelixConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.FelixConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.FelixConfiguration, err error)
	FelixConfigurationExpansion
}

// felixConfigurations implements FelixConfigurationInterface
type felixConfigurations struct {
	client rest.Interface
}

// newFelixConfigurations returns a FelixConfigurations
func newFelixConfigurations(c *ProjectcalicoV3Client) *felixConfigurations {
	return &felixConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the felixConfiguration, and returns the corresponding felixConfiguration object, and an error if there is any.
func (c *felixConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.FelixConfiguration, err error) {
	result = &v3.FelixConfiguration{}
	err = c.client.Get().
		Resource("felixconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FelixConfigurations that match those selectors.
func (c *felixConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v3.FelixConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.FelixConfigurationList{}
	err = c.client.Get().
		Resource("felixconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested felixConfigurations.
func (c *felixConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("felixconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a felixConfiguration and creates it.  Returns the server's representation of the felixConfiguration, and an error, if there is any.
func (c *felixConfigurations) Create(ctx context.Context, felixConfiguration *v3.FelixConfiguration, opts v1.CreateOptions) (result *v3.FelixConfiguration, err error) {
	result = &v3.FelixConfiguration{}
	err = c.client.Post().
		Resource("felixconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(felixConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a felixConfiguration and updates it. Returns the server's representation of the felixConfiguration, and an error, if there is any.
func (c *felixConfigurations) Update(ctx context.Context, felixConfiguration *v3.FelixConfiguration, opts v1.UpdateOptions) (result *v3.FelixConfiguration, err error) {
	result = &v3.FelixConfiguration{}
	err = c.client.Put().
		Resource("felixconfigurations").
		Name(felixConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(felixConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the felixConfiguration and deletes it. Returns an error if one occurs.
func (c *felixConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("felixconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *felixConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("felixconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched felixConfiguration.
func (c *felixConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.FelixConfiguration, err error) {
	result = &v3.FelixConfiguration{}
	err = c.client.Patch(pt).
		Resource("felixconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
