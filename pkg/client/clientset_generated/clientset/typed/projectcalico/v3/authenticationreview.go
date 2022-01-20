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

// AuthenticationReviewsGetter has a method to return a AuthenticationReviewInterface.
// A group's client should implement this interface.
type AuthenticationReviewsGetter interface {
	AuthenticationReviews() AuthenticationReviewInterface
}

// AuthenticationReviewInterface has methods to work with AuthenticationReview resources.
type AuthenticationReviewInterface interface {
	Create(ctx context.Context, authenticationReview *v3.AuthenticationReview, opts v1.CreateOptions) (*v3.AuthenticationReview, error)
	Update(ctx context.Context, authenticationReview *v3.AuthenticationReview, opts v1.UpdateOptions) (*v3.AuthenticationReview, error)
	UpdateStatus(ctx context.Context, authenticationReview *v3.AuthenticationReview, opts v1.UpdateOptions) (*v3.AuthenticationReview, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.AuthenticationReview, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.AuthenticationReviewList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.AuthenticationReview, err error)
	AuthenticationReviewExpansion
}

// authenticationReviews implements AuthenticationReviewInterface
type authenticationReviews struct {
	client rest.Interface
}

// newAuthenticationReviews returns a AuthenticationReviews
func newAuthenticationReviews(c *ProjectcalicoV3Client) *authenticationReviews {
	return &authenticationReviews{
		client: c.RESTClient(),
	}
}

// Get takes name of the authenticationReview, and returns the corresponding authenticationReview object, and an error if there is any.
func (c *authenticationReviews) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.AuthenticationReview, err error) {
	result = &v3.AuthenticationReview{}
	err = c.client.Get().
		Resource("authenticationreviews").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AuthenticationReviews that match those selectors.
func (c *authenticationReviews) List(ctx context.Context, opts v1.ListOptions) (result *v3.AuthenticationReviewList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.AuthenticationReviewList{}
	err = c.client.Get().
		Resource("authenticationreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested authenticationReviews.
func (c *authenticationReviews) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("authenticationreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a authenticationReview and creates it.  Returns the server's representation of the authenticationReview, and an error, if there is any.
func (c *authenticationReviews) Create(ctx context.Context, authenticationReview *v3.AuthenticationReview, opts v1.CreateOptions) (result *v3.AuthenticationReview, err error) {
	result = &v3.AuthenticationReview{}
	err = c.client.Post().
		Resource("authenticationreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authenticationReview).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a authenticationReview and updates it. Returns the server's representation of the authenticationReview, and an error, if there is any.
func (c *authenticationReviews) Update(ctx context.Context, authenticationReview *v3.AuthenticationReview, opts v1.UpdateOptions) (result *v3.AuthenticationReview, err error) {
	result = &v3.AuthenticationReview{}
	err = c.client.Put().
		Resource("authenticationreviews").
		Name(authenticationReview.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authenticationReview).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *authenticationReviews) UpdateStatus(ctx context.Context, authenticationReview *v3.AuthenticationReview, opts v1.UpdateOptions) (result *v3.AuthenticationReview, err error) {
	result = &v3.AuthenticationReview{}
	err = c.client.Put().
		Resource("authenticationreviews").
		Name(authenticationReview.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authenticationReview).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the authenticationReview and deletes it. Returns an error if one occurs.
func (c *authenticationReviews) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("authenticationreviews").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *authenticationReviews) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("authenticationreviews").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched authenticationReview.
func (c *authenticationReviews) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.AuthenticationReview, err error) {
	result = &v3.AuthenticationReview{}
	err = c.client.Patch(pt).
		Resource("authenticationreviews").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
