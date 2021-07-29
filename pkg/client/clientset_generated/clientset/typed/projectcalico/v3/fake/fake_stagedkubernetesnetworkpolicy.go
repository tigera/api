// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
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
}

var stagedkubernetesnetworkpoliciesResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "v3", Resource: "stagedkubernetesnetworkpolicies"}

var stagedkubernetesnetworkpoliciesKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: "StagedKubernetesNetworkPolicy"}

// Get takes name of the stagedKubernetesNetworkPolicy, and returns the corresponding stagedKubernetesNetworkPolicy object, and an error if there is any.
func (c *FakeStagedKubernetesNetworkPolicies) Get(name string, options v1.GetOptions) (result *v3.StagedKubernetesNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(stagedkubernetesnetworkpoliciesResource, name), &v3.StagedKubernetesNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.StagedKubernetesNetworkPolicy), err
}

// List takes label and field selectors, and returns the list of StagedKubernetesNetworkPolicies that match those selectors.
func (c *FakeStagedKubernetesNetworkPolicies) List(opts v1.ListOptions) (result *v3.StagedKubernetesNetworkPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(stagedkubernetesnetworkpoliciesResource, stagedkubernetesnetworkpoliciesKind, opts), &v3.StagedKubernetesNetworkPolicyList{})
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
func (c *FakeStagedKubernetesNetworkPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(stagedkubernetesnetworkpoliciesResource, opts))
}

// Create takes the representation of a stagedKubernetesNetworkPolicy and creates it.  Returns the server's representation of the stagedKubernetesNetworkPolicy, and an error, if there is any.
func (c *FakeStagedKubernetesNetworkPolicies) Create(stagedKubernetesNetworkPolicy *v3.StagedKubernetesNetworkPolicy) (result *v3.StagedKubernetesNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(stagedkubernetesnetworkpoliciesResource, stagedKubernetesNetworkPolicy), &v3.StagedKubernetesNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.StagedKubernetesNetworkPolicy), err
}

// Update takes the representation of a stagedKubernetesNetworkPolicy and updates it. Returns the server's representation of the stagedKubernetesNetworkPolicy, and an error, if there is any.
func (c *FakeStagedKubernetesNetworkPolicies) Update(stagedKubernetesNetworkPolicy *v3.StagedKubernetesNetworkPolicy) (result *v3.StagedKubernetesNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(stagedkubernetesnetworkpoliciesResource, stagedKubernetesNetworkPolicy), &v3.StagedKubernetesNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.StagedKubernetesNetworkPolicy), err
}

// Delete takes name of the stagedKubernetesNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeStagedKubernetesNetworkPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(stagedkubernetesnetworkpoliciesResource, name), &v3.StagedKubernetesNetworkPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStagedKubernetesNetworkPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(stagedkubernetesnetworkpoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v3.StagedKubernetesNetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched stagedKubernetesNetworkPolicy.
func (c *FakeStagedKubernetesNetworkPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.StagedKubernetesNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(stagedkubernetesnetworkpoliciesResource, name, pt, data, subresources...), &v3.StagedKubernetesNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.StagedKubernetesNetworkPolicy), err
}
