// Copyright (c) 2022 Tigera, Inc. All rights reserved.

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

// IPReservationsGetter has a method to return a IPReservationInterface.
// A group's client should implement this interface.
type IPReservationsGetter interface {
	IPReservations() IPReservationInterface
}

// IPReservationInterface has methods to work with IPReservation resources.
type IPReservationInterface interface {
	Create(ctx context.Context, iPReservation *v3.IPReservation, opts v1.CreateOptions) (*v3.IPReservation, error)
	Update(ctx context.Context, iPReservation *v3.IPReservation, opts v1.UpdateOptions) (*v3.IPReservation, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.IPReservation, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.IPReservationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.IPReservation, err error)
	IPReservationExpansion
}

// iPReservations implements IPReservationInterface
type iPReservations struct {
	client rest.Interface
}

// newIPReservations returns a IPReservations
func newIPReservations(c *ProjectcalicoV3Client) *iPReservations {
	return &iPReservations{
		client: c.RESTClient(),
	}
}

// Get takes name of the iPReservation, and returns the corresponding iPReservation object, and an error if there is any.
func (c *iPReservations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.IPReservation, err error) {
	result = &v3.IPReservation{}
	err = c.client.Get().
		Resource("ipreservations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IPReservations that match those selectors.
func (c *iPReservations) List(ctx context.Context, opts v1.ListOptions) (result *v3.IPReservationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.IPReservationList{}
	err = c.client.Get().
		Resource("ipreservations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iPReservations.
func (c *iPReservations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ipreservations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a iPReservation and creates it.  Returns the server's representation of the iPReservation, and an error, if there is any.
func (c *iPReservations) Create(ctx context.Context, iPReservation *v3.IPReservation, opts v1.CreateOptions) (result *v3.IPReservation, err error) {
	result = &v3.IPReservation{}
	err = c.client.Post().
		Resource("ipreservations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPReservation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a iPReservation and updates it. Returns the server's representation of the iPReservation, and an error, if there is any.
func (c *iPReservations) Update(ctx context.Context, iPReservation *v3.IPReservation, opts v1.UpdateOptions) (result *v3.IPReservation, err error) {
	result = &v3.IPReservation{}
	err = c.client.Put().
		Resource("ipreservations").
		Name(iPReservation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPReservation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the iPReservation and deletes it. Returns an error if one occurs.
func (c *iPReservations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ipreservations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iPReservations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ipreservations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched iPReservation.
func (c *iPReservations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.IPReservation, err error) {
	result = &v3.IPReservation{}
	err = c.client.Patch(pt).
		Resource("ipreservations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
