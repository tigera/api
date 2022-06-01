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

// StagedNetworkPoliciesGetter has a method to return a StagedNetworkPolicyInterface.
// A group's client should implement this interface.
type StagedNetworkPoliciesGetter interface {
	StagedNetworkPolicies(namespace string) StagedNetworkPolicyInterface
}

// StagedNetworkPolicyInterface has methods to work with StagedNetworkPolicy resources.
type StagedNetworkPolicyInterface interface {
	Create(ctx context.Context, stagedNetworkPolicy *v3.StagedNetworkPolicy, opts v1.CreateOptions) (*v3.StagedNetworkPolicy, error)
	Update(ctx context.Context, stagedNetworkPolicy *v3.StagedNetworkPolicy, opts v1.UpdateOptions) (*v3.StagedNetworkPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.StagedNetworkPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.StagedNetworkPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.StagedNetworkPolicy, err error)
	StagedNetworkPolicyExpansion
}

// stagedNetworkPolicies implements StagedNetworkPolicyInterface
type stagedNetworkPolicies struct {
	client rest.Interface
	ns     string
}

// newStagedNetworkPolicies returns a StagedNetworkPolicies
func newStagedNetworkPolicies(c *ProjectcalicoV3Client, namespace string) *stagedNetworkPolicies {
	return &stagedNetworkPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the stagedNetworkPolicy, and returns the corresponding stagedNetworkPolicy object, and an error if there is any.
func (c *stagedNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.StagedNetworkPolicy, err error) {
	result = &v3.StagedNetworkPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("stagednetworkpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StagedNetworkPolicies that match those selectors.
func (c *stagedNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v3.StagedNetworkPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.StagedNetworkPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("stagednetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested stagedNetworkPolicies.
func (c *stagedNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("stagednetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a stagedNetworkPolicy and creates it.  Returns the server's representation of the stagedNetworkPolicy, and an error, if there is any.
func (c *stagedNetworkPolicies) Create(ctx context.Context, stagedNetworkPolicy *v3.StagedNetworkPolicy, opts v1.CreateOptions) (result *v3.StagedNetworkPolicy, err error) {
	result = &v3.StagedNetworkPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("stagednetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(stagedNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a stagedNetworkPolicy and updates it. Returns the server's representation of the stagedNetworkPolicy, and an error, if there is any.
func (c *stagedNetworkPolicies) Update(ctx context.Context, stagedNetworkPolicy *v3.StagedNetworkPolicy, opts v1.UpdateOptions) (result *v3.StagedNetworkPolicy, err error) {
	result = &v3.StagedNetworkPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("stagednetworkpolicies").
		Name(stagedNetworkPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(stagedNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the stagedNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *stagedNetworkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("stagednetworkpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *stagedNetworkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("stagednetworkpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched stagedNetworkPolicy.
func (c *stagedNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.StagedNetworkPolicy, err error) {
	result = &v3.StagedNetworkPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("stagednetworkpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
