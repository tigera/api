// Copyright (c) 2024 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEgressGatewayPolicies implements EgressGatewayPolicyInterface
type FakeEgressGatewayPolicies struct {
	Fake *FakeProjectcalicoV3
}

var egressgatewaypoliciesResource = v3.SchemeGroupVersion.WithResource("egressgatewaypolicies")

var egressgatewaypoliciesKind = v3.SchemeGroupVersion.WithKind("EgressGatewayPolicy")

// Get takes name of the egressGatewayPolicy, and returns the corresponding egressGatewayPolicy object, and an error if there is any.
func (c *FakeEgressGatewayPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.EgressGatewayPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(egressgatewaypoliciesResource, name), &v3.EgressGatewayPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.EgressGatewayPolicy), err
}

// List takes label and field selectors, and returns the list of EgressGatewayPolicies that match those selectors.
func (c *FakeEgressGatewayPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v3.EgressGatewayPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(egressgatewaypoliciesResource, egressgatewaypoliciesKind, opts), &v3.EgressGatewayPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.EgressGatewayPolicyList{ListMeta: obj.(*v3.EgressGatewayPolicyList).ListMeta}
	for _, item := range obj.(*v3.EgressGatewayPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested egressGatewayPolicies.
func (c *FakeEgressGatewayPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(egressgatewaypoliciesResource, opts))
}

// Create takes the representation of a egressGatewayPolicy and creates it.  Returns the server's representation of the egressGatewayPolicy, and an error, if there is any.
func (c *FakeEgressGatewayPolicies) Create(ctx context.Context, egressGatewayPolicy *v3.EgressGatewayPolicy, opts v1.CreateOptions) (result *v3.EgressGatewayPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(egressgatewaypoliciesResource, egressGatewayPolicy), &v3.EgressGatewayPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.EgressGatewayPolicy), err
}

// Update takes the representation of a egressGatewayPolicy and updates it. Returns the server's representation of the egressGatewayPolicy, and an error, if there is any.
func (c *FakeEgressGatewayPolicies) Update(ctx context.Context, egressGatewayPolicy *v3.EgressGatewayPolicy, opts v1.UpdateOptions) (result *v3.EgressGatewayPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(egressgatewaypoliciesResource, egressGatewayPolicy), &v3.EgressGatewayPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.EgressGatewayPolicy), err
}

// Delete takes name of the egressGatewayPolicy and deletes it. Returns an error if one occurs.
func (c *FakeEgressGatewayPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(egressgatewaypoliciesResource, name, opts), &v3.EgressGatewayPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEgressGatewayPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(egressgatewaypoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.EgressGatewayPolicyList{})
	return err
}

// Patch applies the patch and returns the patched egressGatewayPolicy.
func (c *FakeEgressGatewayPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.EgressGatewayPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(egressgatewaypoliciesResource, name, pt, data, subresources...), &v3.EgressGatewayPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.EgressGatewayPolicy), err
}
