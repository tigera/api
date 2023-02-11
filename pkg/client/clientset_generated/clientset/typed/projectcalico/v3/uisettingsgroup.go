// Copyright (c) 2023 Tigera, Inc. All rights reserved.

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

// UISettingsGroupsGetter has a method to return a UISettingsGroupInterface.
// A group's client should implement this interface.
type UISettingsGroupsGetter interface {
	UISettingsGroups() UISettingsGroupInterface
}

// UISettingsGroupInterface has methods to work with UISettingsGroup resources.
type UISettingsGroupInterface interface {
	Create(ctx context.Context, uISettingsGroup *v3.UISettingsGroup, opts v1.CreateOptions) (*v3.UISettingsGroup, error)
	Update(ctx context.Context, uISettingsGroup *v3.UISettingsGroup, opts v1.UpdateOptions) (*v3.UISettingsGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.UISettingsGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.UISettingsGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.UISettingsGroup, err error)
	UISettingsGroupExpansion
}

// uISettingsGroups implements UISettingsGroupInterface
type uISettingsGroups struct {
	client rest.Interface
}

// newUISettingsGroups returns a UISettingsGroups
func newUISettingsGroups(c *ProjectcalicoV3Client) *uISettingsGroups {
	return &uISettingsGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the uISettingsGroup, and returns the corresponding uISettingsGroup object, and an error if there is any.
func (c *uISettingsGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.UISettingsGroup, err error) {
	result = &v3.UISettingsGroup{}
	err = c.client.Get().
		Resource("uisettingsgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UISettingsGroups that match those selectors.
func (c *uISettingsGroups) List(ctx context.Context, opts v1.ListOptions) (result *v3.UISettingsGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.UISettingsGroupList{}
	err = c.client.Get().
		Resource("uisettingsgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested uISettingsGroups.
func (c *uISettingsGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("uisettingsgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a uISettingsGroup and creates it.  Returns the server's representation of the uISettingsGroup, and an error, if there is any.
func (c *uISettingsGroups) Create(ctx context.Context, uISettingsGroup *v3.UISettingsGroup, opts v1.CreateOptions) (result *v3.UISettingsGroup, err error) {
	result = &v3.UISettingsGroup{}
	err = c.client.Post().
		Resource("uisettingsgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(uISettingsGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a uISettingsGroup and updates it. Returns the server's representation of the uISettingsGroup, and an error, if there is any.
func (c *uISettingsGroups) Update(ctx context.Context, uISettingsGroup *v3.UISettingsGroup, opts v1.UpdateOptions) (result *v3.UISettingsGroup, err error) {
	result = &v3.UISettingsGroup{}
	err = c.client.Put().
		Resource("uisettingsgroups").
		Name(uISettingsGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(uISettingsGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the uISettingsGroup and deletes it. Returns an error if one occurs.
func (c *uISettingsGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("uisettingsgroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *uISettingsGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("uisettingsgroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched uISettingsGroup.
func (c *uISettingsGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.UISettingsGroup, err error) {
	result = &v3.UISettingsGroup{}
	err = c.client.Patch(pt).
		Resource("uisettingsgroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
