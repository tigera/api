// Copyright (c) 2026 Tigera, Inc. All rights reserved.

package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindNetwork     = "Network"
	KindNetworkList = "NetworkList"
)

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkList is a list of Network resources.
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Network `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// Network defines a logical network within a Calico cluster.  Each Network has a type
// (VRF, ...) that determines how pods on that network are isolated and how
// their traffic is routed.
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   NetworkSpec   `json:"spec" protobuf:"bytes,2,opt,name=spec"`
	Status NetworkStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// NetworkSpec contains the specification for a Network resource.  Exactly one of the
// network-type fields (vrf, ...) must be set.
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type NetworkSpec struct {
	// VRF network configuration.
	// Pods interfaces on a VRF network are isolated in a Linux VRF and can only access their own VRF.
	// +optional
	VRF *VRFNetworkSpec `json:"vrf,omitempty"`
}

// VRFNetworkSpec configures a VRF-based network that isolates pods in a Linux VRF.
type VRFNetworkSpec struct {
	// Routing controls cluster-wide routing behaviour for this VRF network.
	// +optional
	// +kubebuilder:default={inClusterMode: "Local"}
	Routing VRFRouting `json:"routing,omitempty"`

	// HostConfig defines per-node configuration for this VRF network.  At least one entry
	// must be specified.  When multiple entries are present (e.g. one per rack), they must
	// have disjoint nodeSelectors although this is not enforced.  For a given node, the
	// first matching entry is applied and all others are ignored.
	//
	// The list is stored and served as a JSON array; Kubernetes preserves the
	// order in which entries were submitted (including across Server-Side
	// Apply merges keyed on nodeSelector), so the result of `kubectl get -oyaml`
	// is authoritative for "first match wins".  Single-actor edits keep the
	// order users see in the YAML they edit.
	// +required
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=100
	// +listType=map
	// +listMapKey=nodeSelector
	HostConfig []VRFHostConfig `json:"hostConfig" validate:"required,dive"`
}

// InClusterRoutingMode controls how routes to pods on other nodes are programmed inside a
// VRF routing table.
// +kubebuilder:validation:Enum=Local
type InClusterRoutingMode string

const (
	// InClusterRoutingLocal programs only routes to local pods.  Distributing routes to pods on
	// other nodes must be handled by BGP.
	InClusterRoutingLocal InClusterRoutingMode = "Local"
)

// VRFRouting holds cluster-wide routing settings for a VRF network.
type VRFRouting struct {
	// InClusterMode controls how Felix programs routes to pods on remote nodes inside
	// the VRF routing table.
	//
	// - Local: Felix programs routes to VRF pods local to this node;
	//   routing to pods on other nodes must be handled by BGP.
	//
	// +optional
	// +kubebuilder:default=Local
	InClusterMode InClusterRoutingMode `json:"inClusterMode,omitempty"`
}

// InterfaceMatch identifies a network interface.  Exactly one match
// criterion must be set.
//
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type InterfaceMatch struct {
	// Name matches a network interface by its exact device name
	// (e.g. "bond0", "eth1", "ens192").
	// +optional
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=15
	Name string `json:"name,omitempty" validate:"omitempty,interface"`
}

// VRFHostConfig provides node-specific VRF settings which may vary across different hosts
// in the cluster.
type VRFHostConfig struct {
	// NodeSelector is a Calico selector expression that determines which nodes this
	// configuration applies to.  If omitted, the entry applies to all nodes.
	// When multiple HostConfig entries are present, the first entry whose selector
	// matches a given node wins.
	// +optional
	// +kubebuilder:default=""
	NodeSelector string `json:"nodeSelector,omitempty" validate:"omitempty,selector"`

	// RouteTableIndex is the Linux kernel routing table number to use for this VRF.
	// Must be unique on these nodes, must not overlap with the RouteTableRanges in
	// FelixConfiguration, and must not collide with tables used by other software on
	// the node.  Tables 253 (default), 254 (main), and 255 (local) are reserved by
	// the kernel.
	// +required
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2147483647
	RouteTableIndex uint32 `json:"routeTableIndex" validate:"required"`

	// Interfaces on the node to attach to the VRF.  The IP address, and local routes
	// for that IP address will move into the VRF routing table.  At least one interface
	// should be specified to allow pods in the VRF to communicate outside the node.
	// +optional
	// +listType=atomic
	HostInterfaces []InterfaceMatch `json:"hostInterfaces"`

	// StaticRoutes are additional routes programmed into the VRF routing table, beyond
	// the pod routes that Felix manages automatically and routes derived from VRF
	// interface addresses.
	// +optional
	// +listType=atomic
	StaticRoutes []VRFStaticRoute `json:"staticRoutes,omitempty" validate:"omitempty,dive"`
}

// VRFStaticRoute defines a static route to program in a VRF routing table.
type VRFStaticRoute struct {
	// Destination is the CIDR prefix for this route.  Use "0.0.0.0/0" or "::/0" for a
	// default route.
	// +required
	Destination string `json:"destination" validate:"required,cidr"`

	// Action determines how traffic matching this route is handled.  Exactly one action
	// field must be set.
	// +required
	Action StaticRouteAction `json:"action" validate:"required"`
}

// StaticRouteAction defines the forwarding behaviour for a static route.  Exactly one
// field must be set.
//
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type StaticRouteAction struct {
	// NextHop forwards matching traffic to the specified gateway IP.  The address must be
	// reachable on the subnet of one of the VRF interfaces on the node.
	// +optional
	NextHop *string `json:"nextHop,omitempty" validate:"omitempty,ip"`
}

// NetworkStatus reports the observed state of the Network resource.
type NetworkStatus struct {
	// Conditions is a list of conditions that apply to this network.
	// +listType=atomic
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// NewNetwork creates a new (zeroed) Network struct with the TypeMetadata initialised to the current
// version.
func NewNetwork() *Network {
	return &Network{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindNetwork,
			APIVersion: GroupVersionCurrent,
		},
	}
}
