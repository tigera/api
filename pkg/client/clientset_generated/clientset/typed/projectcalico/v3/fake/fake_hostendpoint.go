// Copyright (c) 2023 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHostEndpoints implements HostEndpointInterface
type FakeHostEndpoints struct {
	Fake *FakeProjectcalicoV3
}

var hostendpointsResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "v3", Resource: "hostendpoints"}

var hostendpointsKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: "HostEndpoint"}

// Get takes name of the hostEndpoint, and returns the corresponding hostEndpoint object, and an error if there is any.
func (c *FakeHostEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.HostEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(hostendpointsResource, name), &v3.HostEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.HostEndpoint), err
}

// List takes label and field selectors, and returns the list of HostEndpoints that match those selectors.
func (c *FakeHostEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v3.HostEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(hostendpointsResource, hostendpointsKind, opts), &v3.HostEndpointList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.HostEndpointList{ListMeta: obj.(*v3.HostEndpointList).ListMeta}
	for _, item := range obj.(*v3.HostEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hostEndpoints.
func (c *FakeHostEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(hostendpointsResource, opts))
}

// Create takes the representation of a hostEndpoint and creates it.  Returns the server's representation of the hostEndpoint, and an error, if there is any.
func (c *FakeHostEndpoints) Create(ctx context.Context, hostEndpoint *v3.HostEndpoint, opts v1.CreateOptions) (result *v3.HostEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(hostendpointsResource, hostEndpoint), &v3.HostEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.HostEndpoint), err
}

// Update takes the representation of a hostEndpoint and updates it. Returns the server's representation of the hostEndpoint, and an error, if there is any.
func (c *FakeHostEndpoints) Update(ctx context.Context, hostEndpoint *v3.HostEndpoint, opts v1.UpdateOptions) (result *v3.HostEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(hostendpointsResource, hostEndpoint), &v3.HostEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.HostEndpoint), err
}

// Delete takes name of the hostEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeHostEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(hostendpointsResource, name, opts), &v3.HostEndpoint{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHostEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(hostendpointsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.HostEndpointList{})
	return err
}

// Patch applies the patch and returns the patched hostEndpoint.
func (c *FakeHostEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.HostEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(hostendpointsResource, name, pt, data, subresources...), &v3.HostEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.HostEndpoint), err
}
