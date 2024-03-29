// Copyright (c) 2024 Tigera, Inc. All rights reserved.

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

// PolicyRecommendationScopesGetter has a method to return a PolicyRecommendationScopeInterface.
// A group's client should implement this interface.
type PolicyRecommendationScopesGetter interface {
	PolicyRecommendationScopes() PolicyRecommendationScopeInterface
}

// PolicyRecommendationScopeInterface has methods to work with PolicyRecommendationScope resources.
type PolicyRecommendationScopeInterface interface {
	Create(ctx context.Context, policyRecommendationScope *v3.PolicyRecommendationScope, opts v1.CreateOptions) (*v3.PolicyRecommendationScope, error)
	Update(ctx context.Context, policyRecommendationScope *v3.PolicyRecommendationScope, opts v1.UpdateOptions) (*v3.PolicyRecommendationScope, error)
	UpdateStatus(ctx context.Context, policyRecommendationScope *v3.PolicyRecommendationScope, opts v1.UpdateOptions) (*v3.PolicyRecommendationScope, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.PolicyRecommendationScope, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.PolicyRecommendationScopeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.PolicyRecommendationScope, err error)
	PolicyRecommendationScopeExpansion
}

// policyRecommendationScopes implements PolicyRecommendationScopeInterface
type policyRecommendationScopes struct {
	client rest.Interface
}

// newPolicyRecommendationScopes returns a PolicyRecommendationScopes
func newPolicyRecommendationScopes(c *ProjectcalicoV3Client) *policyRecommendationScopes {
	return &policyRecommendationScopes{
		client: c.RESTClient(),
	}
}

// Get takes name of the policyRecommendationScope, and returns the corresponding policyRecommendationScope object, and an error if there is any.
func (c *policyRecommendationScopes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.PolicyRecommendationScope, err error) {
	result = &v3.PolicyRecommendationScope{}
	err = c.client.Get().
		Resource("policyrecommendationscopes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PolicyRecommendationScopes that match those selectors.
func (c *policyRecommendationScopes) List(ctx context.Context, opts v1.ListOptions) (result *v3.PolicyRecommendationScopeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.PolicyRecommendationScopeList{}
	err = c.client.Get().
		Resource("policyrecommendationscopes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested policyRecommendationScopes.
func (c *policyRecommendationScopes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("policyrecommendationscopes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a policyRecommendationScope and creates it.  Returns the server's representation of the policyRecommendationScope, and an error, if there is any.
func (c *policyRecommendationScopes) Create(ctx context.Context, policyRecommendationScope *v3.PolicyRecommendationScope, opts v1.CreateOptions) (result *v3.PolicyRecommendationScope, err error) {
	result = &v3.PolicyRecommendationScope{}
	err = c.client.Post().
		Resource("policyrecommendationscopes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(policyRecommendationScope).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a policyRecommendationScope and updates it. Returns the server's representation of the policyRecommendationScope, and an error, if there is any.
func (c *policyRecommendationScopes) Update(ctx context.Context, policyRecommendationScope *v3.PolicyRecommendationScope, opts v1.UpdateOptions) (result *v3.PolicyRecommendationScope, err error) {
	result = &v3.PolicyRecommendationScope{}
	err = c.client.Put().
		Resource("policyrecommendationscopes").
		Name(policyRecommendationScope.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(policyRecommendationScope).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *policyRecommendationScopes) UpdateStatus(ctx context.Context, policyRecommendationScope *v3.PolicyRecommendationScope, opts v1.UpdateOptions) (result *v3.PolicyRecommendationScope, err error) {
	result = &v3.PolicyRecommendationScope{}
	err = c.client.Put().
		Resource("policyrecommendationscopes").
		Name(policyRecommendationScope.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(policyRecommendationScope).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the policyRecommendationScope and deletes it. Returns an error if one occurs.
func (c *policyRecommendationScopes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("policyrecommendationscopes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *policyRecommendationScopes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("policyrecommendationscopes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched policyRecommendationScope.
func (c *policyRecommendationScopes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.PolicyRecommendationScope, err error) {
	result = &v3.PolicyRecommendationScope{}
	err = c.client.Patch(pt).
		Resource("policyrecommendationscopes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
