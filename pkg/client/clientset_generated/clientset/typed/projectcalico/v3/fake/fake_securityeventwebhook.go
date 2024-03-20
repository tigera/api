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

// FakeSecurityEventWebhooks implements SecurityEventWebhookInterface
type FakeSecurityEventWebhooks struct {
	Fake *FakeProjectcalicoV3
}

var securityeventwebhooksResource = v3.SchemeGroupVersion.WithResource("securityeventwebhooks")

var securityeventwebhooksKind = v3.SchemeGroupVersion.WithKind("SecurityEventWebhook")

// Get takes name of the securityEventWebhook, and returns the corresponding securityEventWebhook object, and an error if there is any.
func (c *FakeSecurityEventWebhooks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.SecurityEventWebhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(securityeventwebhooksResource, name), &v3.SecurityEventWebhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.SecurityEventWebhook), err
}

// List takes label and field selectors, and returns the list of SecurityEventWebhooks that match those selectors.
func (c *FakeSecurityEventWebhooks) List(ctx context.Context, opts v1.ListOptions) (result *v3.SecurityEventWebhookList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(securityeventwebhooksResource, securityeventwebhooksKind, opts), &v3.SecurityEventWebhookList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.SecurityEventWebhookList{ListMeta: obj.(*v3.SecurityEventWebhookList).ListMeta}
	for _, item := range obj.(*v3.SecurityEventWebhookList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested securityEventWebhooks.
func (c *FakeSecurityEventWebhooks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(securityeventwebhooksResource, opts))
}

// Create takes the representation of a securityEventWebhook and creates it.  Returns the server's representation of the securityEventWebhook, and an error, if there is any.
func (c *FakeSecurityEventWebhooks) Create(ctx context.Context, securityEventWebhook *v3.SecurityEventWebhook, opts v1.CreateOptions) (result *v3.SecurityEventWebhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(securityeventwebhooksResource, securityEventWebhook), &v3.SecurityEventWebhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.SecurityEventWebhook), err
}

// Update takes the representation of a securityEventWebhook and updates it. Returns the server's representation of the securityEventWebhook, and an error, if there is any.
func (c *FakeSecurityEventWebhooks) Update(ctx context.Context, securityEventWebhook *v3.SecurityEventWebhook, opts v1.UpdateOptions) (result *v3.SecurityEventWebhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(securityeventwebhooksResource, securityEventWebhook), &v3.SecurityEventWebhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.SecurityEventWebhook), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecurityEventWebhooks) UpdateStatus(ctx context.Context, securityEventWebhook *v3.SecurityEventWebhook, opts v1.UpdateOptions) (*v3.SecurityEventWebhook, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(securityeventwebhooksResource, "status", securityEventWebhook), &v3.SecurityEventWebhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.SecurityEventWebhook), err
}

// Delete takes name of the securityEventWebhook and deletes it. Returns an error if one occurs.
func (c *FakeSecurityEventWebhooks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(securityeventwebhooksResource, name, opts), &v3.SecurityEventWebhook{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecurityEventWebhooks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(securityeventwebhooksResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.SecurityEventWebhookList{})
	return err
}

// Patch applies the patch and returns the patched securityEventWebhook.
func (c *FakeSecurityEventWebhooks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.SecurityEventWebhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(securityeventwebhooksResource, name, pt, data, subresources...), &v3.SecurityEventWebhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.SecurityEventWebhook), err
}
