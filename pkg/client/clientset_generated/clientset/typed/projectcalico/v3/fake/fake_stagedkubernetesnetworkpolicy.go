// Copyright (c) 2021 Tigera, Inc. All rights reserved.

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

// FakeStagedKubernetesNetworkPolicies implements StagedKubernetesNetworkPolicyInterface
type FakeStagedKubernetesNetworkPolicies struct {
	Fake *FakeProjectcalicoV3
	ns   string
}

var stagedkubernetesnetworkpoliciesResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "v3", Resource: "stagedkubernetesnetworkpolicies"}

var stagedkubernetesnetworkpoliciesKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: "StagedKubernetesNetworkPolicy"}

// Get takes name of the stagedKubernetesNetworkPolicy, and returns the corresponding stagedKubernetesNetworkPolicy object, and an error if there is any.
func (c *FakeStagedKubernetesNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.StagedKubernetesNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(stagedkubernetesnetworkpoliciesResource, c.ns, name), &v3.StagedKubernetesNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.StagedKubernetesNetworkPolicy), err
}

// List takes label and field selectors, and returns the list of StagedKubernetesNetworkPolicies that match those selectors.
func (c *FakeStagedKubernetesNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v3.StagedKubernetesNetworkPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(stagedkubernetesnetworkpoliciesResource, stagedkubernetesnetworkpoliciesKind, c.ns, opts), &v3.StagedKubernetesNetworkPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.StagedKubernetesNetworkPolicyList{ListMeta: obj.(*v3.StagedKubernetesNetworkPolicyList).ListMeta}
	for _, item := range obj.(*v3.StagedKubernetesNetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested stagedKubernetesNetworkPolicies.
func (c *FakeStagedKubernetesNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(stagedkubernetesnetworkpoliciesResource, c.ns, opts))

}

// Create takes the representation of a stagedKubernetesNetworkPolicy and creates it.  Returns the server's representation of the stagedKubernetesNetworkPolicy, and an error, if there is any.
func (c *FakeStagedKubernetesNetworkPolicies) Create(ctx context.Context, stagedKubernetesNetworkPolicy *v3.StagedKubernetesNetworkPolicy, opts v1.CreateOptions) (result *v3.StagedKubernetesNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(stagedkubernetesnetworkpoliciesResource, c.ns, stagedKubernetesNetworkPolicy), &v3.StagedKubernetesNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.StagedKubernetesNetworkPolicy), err
}

// Update takes the representation of a stagedKubernetesNetworkPolicy and updates it. Returns the server's representation of the stagedKubernetesNetworkPolicy, and an error, if there is any.
func (c *FakeStagedKubernetesNetworkPolicies) Update(ctx context.Context, stagedKubernetesNetworkPolicy *v3.StagedKubernetesNetworkPolicy, opts v1.UpdateOptions) (result *v3.StagedKubernetesNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(stagedkubernetesnetworkpoliciesResource, c.ns, stagedKubernetesNetworkPolicy), &v3.StagedKubernetesNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.StagedKubernetesNetworkPolicy), err
}

// Delete takes name of the stagedKubernetesNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeStagedKubernetesNetworkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(stagedkubernetesnetworkpoliciesResource, c.ns, name), &v3.StagedKubernetesNetworkPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStagedKubernetesNetworkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(stagedkubernetesnetworkpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.StagedKubernetesNetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched stagedKubernetesNetworkPolicy.
func (c *FakeStagedKubernetesNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.StagedKubernetesNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(stagedkubernetesnetworkpoliciesResource, c.ns, name, pt, data, subresources...), &v3.StagedKubernetesNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.StagedKubernetesNetworkPolicy), err
}
