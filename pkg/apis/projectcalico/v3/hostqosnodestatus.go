// Copyright (c) 2026 Tigera, Inc. All rights reserved.

package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindHostQoSNodeStatus     = "HostQoSNodeStatus"
	KindHostQoSNodeStatusList = "HostQoSNodeStatusList"
)

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HostQoSNodeStatusList contains a list of HostQoSNodeStatus resources.
type HostQoSNodeStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []HostQoSNodeStatus `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster,shortName={hqns}
// +kubebuilder:printcolumn:name="Policy",type="string",JSONPath=".spec.policy"
// +kubebuilder:printcolumn:name="Node",type="string",JSONPath=".spec.node"
// +kubebuilder:printcolumn:name="HostType",type="string",JSONPath=".spec.hostType"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// HostQoSNodeStatus is one host's view of one HostQoSPolicy, written by that host's Felix.
type HostQoSNodeStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`

	Spec HostQoSNodeStatusSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`

	// +optional
	Status HostQoSNodeStatusStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// HostQoSNodeStatusSpec identifies the policy and host this status describes.
type HostQoSNodeStatusSpec struct {
	// Policy is the name of the HostQoSPolicy.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Policy string `json:"policy" validate:"required,name"`

	// Node is the cluster node name, or the hostname of a non-cluster host.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Node string `json:"node" validate:"required"`

	// HostType keeps a cluster node apart from a non-cluster host of the same name.
	// +kubebuilder:validation:Required
	HostType HostType `json:"hostType"`
}

// HostQoSNodeStatusStatus is the state of the policy's bindings on this host.
type HostQoSNodeStatusStatus struct {
	// State summarizes the bindings; PartiallyApplied when they mix Applied with other states.
	// +optional
	State *HostQoSPolicyState `json:"state,omitempty"`

	// Bindings has one entry per (HostEndpoint, interface) the policy resolves to on this host.
	// +optional
	// +listType=map
	// +listMapKey=hostEndpointName
	// +listMapKey=interfaceName
	// +kubebuilder:validation:MaxItems=256
	Bindings []HostQoSPolicyBindingStatus `json:"bindings,omitempty"`

	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`
}

// NewHostQoSNodeStatus creates a new (zeroed) HostQoSNodeStatus struct with the TypeMetadata
// initialised to the current version.
func NewHostQoSNodeStatus() *HostQoSNodeStatus {
	return &HostQoSNodeStatus{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindHostQoSNodeStatus,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewHostQoSNodeStatusList creates a new (zeroed) HostQoSNodeStatusList struct with the
// TypeMetadata initialised to the current version.
func NewHostQoSNodeStatusList() *HostQoSNodeStatusList {
	return &HostQoSNodeStatusList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindHostQoSNodeStatusList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
