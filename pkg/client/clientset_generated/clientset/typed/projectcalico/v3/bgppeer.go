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

// BGPPeersGetter has a method to return a BGPPeerInterface.
// A group's client should implement this interface.
type BGPPeersGetter interface {
	BGPPeers() BGPPeerInterface
}

// BGPPeerInterface has methods to work with BGPPeer resources.
type BGPPeerInterface interface {
	Create(ctx context.Context, bGPPeer *v3.BGPPeer, opts v1.CreateOptions) (*v3.BGPPeer, error)
	Update(ctx context.Context, bGPPeer *v3.BGPPeer, opts v1.UpdateOptions) (*v3.BGPPeer, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.BGPPeer, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.BGPPeerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.BGPPeer, err error)
	BGPPeerExpansion
}

// bGPPeers implements BGPPeerInterface
type bGPPeers struct {
	client rest.Interface
}

// newBGPPeers returns a BGPPeers
func newBGPPeers(c *ProjectcalicoV3Client) *bGPPeers {
	return &bGPPeers{
		client: c.RESTClient(),
	}
}

// Get takes name of the bGPPeer, and returns the corresponding bGPPeer object, and an error if there is any.
func (c *bGPPeers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.BGPPeer, err error) {
	result = &v3.BGPPeer{}
	err = c.client.Get().
		Resource("bgppeers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BGPPeers that match those selectors.
func (c *bGPPeers) List(ctx context.Context, opts v1.ListOptions) (result *v3.BGPPeerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.BGPPeerList{}
	err = c.client.Get().
		Resource("bgppeers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bGPPeers.
func (c *bGPPeers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("bgppeers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bGPPeer and creates it.  Returns the server's representation of the bGPPeer, and an error, if there is any.
func (c *bGPPeers) Create(ctx context.Context, bGPPeer *v3.BGPPeer, opts v1.CreateOptions) (result *v3.BGPPeer, err error) {
	result = &v3.BGPPeer{}
	err = c.client.Post().
		Resource("bgppeers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bGPPeer).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bGPPeer and updates it. Returns the server's representation of the bGPPeer, and an error, if there is any.
func (c *bGPPeers) Update(ctx context.Context, bGPPeer *v3.BGPPeer, opts v1.UpdateOptions) (result *v3.BGPPeer, err error) {
	result = &v3.BGPPeer{}
	err = c.client.Put().
		Resource("bgppeers").
		Name(bGPPeer.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bGPPeer).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bGPPeer and deletes it. Returns an error if one occurs.
func (c *bGPPeers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("bgppeers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bGPPeers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("bgppeers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bGPPeer.
func (c *bGPPeers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.BGPPeer, err error) {
	result = &v3.BGPPeer{}
	err = c.client.Patch(pt).
		Resource("bgppeers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
