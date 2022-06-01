// Copyright (c) 2022 Tigera, Inc. All rights reserved.

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

// FakeDeepPacketInspections implements DeepPacketInspectionInterface
type FakeDeepPacketInspections struct {
	Fake *FakeProjectcalicoV3
	ns   string
}

var deeppacketinspectionsResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "v3", Resource: "deeppacketinspections"}

var deeppacketinspectionsKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: "DeepPacketInspection"}

// Get takes name of the deepPacketInspection, and returns the corresponding deepPacketInspection object, and an error if there is any.
func (c *FakeDeepPacketInspections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.DeepPacketInspection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deeppacketinspectionsResource, c.ns, name), &v3.DeepPacketInspection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.DeepPacketInspection), err
}

// List takes label and field selectors, and returns the list of DeepPacketInspections that match those selectors.
func (c *FakeDeepPacketInspections) List(ctx context.Context, opts v1.ListOptions) (result *v3.DeepPacketInspectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deeppacketinspectionsResource, deeppacketinspectionsKind, c.ns, opts), &v3.DeepPacketInspectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.DeepPacketInspectionList{ListMeta: obj.(*v3.DeepPacketInspectionList).ListMeta}
	for _, item := range obj.(*v3.DeepPacketInspectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deepPacketInspections.
func (c *FakeDeepPacketInspections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deeppacketinspectionsResource, c.ns, opts))

}

// Create takes the representation of a deepPacketInspection and creates it.  Returns the server's representation of the deepPacketInspection, and an error, if there is any.
func (c *FakeDeepPacketInspections) Create(ctx context.Context, deepPacketInspection *v3.DeepPacketInspection, opts v1.CreateOptions) (result *v3.DeepPacketInspection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deeppacketinspectionsResource, c.ns, deepPacketInspection), &v3.DeepPacketInspection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.DeepPacketInspection), err
}

// Update takes the representation of a deepPacketInspection and updates it. Returns the server's representation of the deepPacketInspection, and an error, if there is any.
func (c *FakeDeepPacketInspections) Update(ctx context.Context, deepPacketInspection *v3.DeepPacketInspection, opts v1.UpdateOptions) (result *v3.DeepPacketInspection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deeppacketinspectionsResource, c.ns, deepPacketInspection), &v3.DeepPacketInspection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.DeepPacketInspection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeepPacketInspections) UpdateStatus(ctx context.Context, deepPacketInspection *v3.DeepPacketInspection, opts v1.UpdateOptions) (*v3.DeepPacketInspection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deeppacketinspectionsResource, "status", c.ns, deepPacketInspection), &v3.DeepPacketInspection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.DeepPacketInspection), err
}

// Delete takes name of the deepPacketInspection and deletes it. Returns an error if one occurs.
func (c *FakeDeepPacketInspections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(deeppacketinspectionsResource, c.ns, name), &v3.DeepPacketInspection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeepPacketInspections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deeppacketinspectionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.DeepPacketInspectionList{})
	return err
}

// Patch applies the patch and returns the patched deepPacketInspection.
func (c *FakeDeepPacketInspections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.DeepPacketInspection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deeppacketinspectionsResource, c.ns, name, pt, data, subresources...), &v3.DeepPacketInspection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.DeepPacketInspection), err
}
