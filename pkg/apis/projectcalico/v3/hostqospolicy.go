// Copyright (c) 2026 Tigera, Inc. All rights reserved.

package v3

import (
	"github.com/tigera/api/pkg/lib/numorstring"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindHostQoSPolicy     = "HostQoSPolicy"
	KindHostQoSPolicyList = "HostQoSPolicyList"
)

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HostQoSPolicyList contains a list of HostQoSPolicy resources.
type HostQoSPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []HostQoSPolicy `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,shortName={hqp}
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="AppliedBindings",type="integer",JSONPath=".status.appliedBindings"
// +kubebuilder:printcolumn:name="ErroredBindings",type="integer",JSONPath=".status.erroredBindings"
// +kubebuilder:printcolumn:name="AppliedNodes",type="integer",JSONPath=".status.appliedNodes"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// HostQoSPolicy is cluster-scoped; it controls per-host-NIC traffic shaping.
type HostQoSPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`

	Spec HostQoSPolicySpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`

	// +optional
	Status HostQoSPolicyStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// CEL cannot sum reservations or read the safety margin, so Felix checks that per interface.

// +kubebuilder:validation:XValidation:rule="!has(self.linkCapacity) || self.resourcePools.all(p, !has(p.egress) || ((!has(p.egress.limit) || quantity(string(p.egress.limit)).compareTo(quantity(string(self.linkCapacity))) <= 0) && quantity(string(p.egress.reservation)).compareTo(quantity(string(self.linkCapacity))) <= 0))",message="egress reservation and limit must be <= linkCapacity",reason=FieldValueInvalid
// +kubebuilder:validation:XValidation:rule="!has(self.linkCapacity) || self.resourcePools.all(p, !has(p.ingress) || ((!has(p.ingress.limit) || quantity(string(p.ingress.limit)).compareTo(quantity(string(self.linkCapacity))) <= 0) && quantity(string(p.ingress.reservation)).compareTo(quantity(string(self.linkCapacity))) <= 0))",message="ingress reservation and limit must be <= linkCapacity",reason=FieldValueInvalid

// HostQoSPolicySpec is shaping intent: link capacity, resource pools, and a selector for the
// HostEndpoints it shapes.
type HostQoSPolicySpec struct {
	// Selector picks the HostEndpoints to shape; the alphabetically-first policy wins a contested one.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=1024
	Selector string `json:"selector" validate:"required,selector"`

	// LinkCapacity in bits/sec, such as 10G; omit to detect it.
	// +optional
	// +kubebuilder:validation:XValidation:rule="quantity(string(self)).isGreaterThan(quantity('0'))",message="linkCapacity must be greater than zero",reason=FieldValueInvalid
	// +kubebuilder:validation:XIntOrString
	// +kubebuilder:validation:Pattern=`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`
	// +kubebuilder:validation:MaxLength=32
	LinkCapacity *resource.Quantity `json:"linkCapacity,omitempty"`

	// ResourcePools are the traffic classes, uniquely named; the first matching pool wins.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=100
	// +listType=map
	// +listMapKey=name
	ResourcePools []NetworkResourcePool `json:"resourcePools" validate:"required,min=1,max=100,dive"`
}

// +kubebuilder:validation:XValidation:rule="has(self.egress) || has(self.ingress)",message="at least one of egress or ingress must be set",reason=FieldValueRequired

// NetworkResourcePool is one traffic class on the shaped interface.
type NetworkResourcePool struct {
	// Name is unique within the policy. Used for status reporting and Prometheus labels.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=128
	// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`
	Name string `json:"name" validate:"required,name"`

	// Egress shapes and classifies traffic leaving the host interface, and can mark DSCP.
	// +optional
	Egress *EgressPoolControl `json:"egress,omitempty" validate:"omitempty"`

	// Ingress shapes and classifies traffic arriving at the host interface.
	// +optional
	Ingress *IngressPoolControl `json:"ingress,omitempty" validate:"omitempty"`

	// Shares weights this pool's share of spare bandwidth; defaults to 1.
	// +optional
	// +kubebuilder:default=1
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1000
	Shares *int32 `json:"shares,omitempty" validate:"omitempty,gte=1,lte=1000"`
}

// NetworkPoolControl is the reservation, limit and traffic match for one direction of a pool.
type NetworkPoolControl struct {
	// Reservation is the guaranteed bandwidth (HTB rate) in decimal SI bits/sec; e.g. "3G".
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XIntOrString
	// +kubebuilder:validation:Pattern=`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`
	// +kubebuilder:validation:MaxLength=32
	Reservation resource.Quantity `json:"reservation"`

	// Limit caps bandwidth (HTB ceil) between reservation and the link capacity, its default.
	// +optional
	// +kubebuilder:validation:XIntOrString
	// +kubebuilder:validation:Pattern=`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`
	// +kubebuilder:validation:MaxLength=32
	Limit *resource.Quantity `json:"limit,omitempty"`

	// Match classifies traffic into this pool; omitted or empty, the pool catches everything.
	// +optional
	Match *TrafficMatch `json:"match,omitempty" validate:"omitempty"`
}

// +kubebuilder:validation:XValidation:rule="quantity(string(self.reservation)).isGreaterThan(quantity('0'))",message="reservation must be greater than zero",reason=FieldValueInvalid
// +kubebuilder:validation:XValidation:rule="!has(self.limit) || quantity(string(self.limit)).compareTo(quantity(string(self.reservation))) >= 0",message="limit must be >= reservation",reason=FieldValueInvalid

// IngressPoolControl is a pool's ingress controls.
type IngressPoolControl struct {
	NetworkPoolControl `json:",inline"`
}

// +kubebuilder:validation:XValidation:rule="quantity(string(self.reservation)).isGreaterThan(quantity('0'))",message="reservation must be greater than zero",reason=FieldValueInvalid
// +kubebuilder:validation:XValidation:rule="!has(self.limit) || quantity(string(self.limit)).compareTo(quantity(string(self.reservation))) >= 0",message="limit must be >= reservation",reason=FieldValueInvalid
// +kubebuilder:validation:XValidation:rule="!has(self.dscp) || type(self.dscp) != int || (self.dscp >= 0 && self.dscp <= 63)",message="dscp must be between 0 and 63",reason=FieldValueInvalid

// EgressPoolControl is a pool's egress controls, plus DSCP marking.
type EgressPoolControl struct {
	NetworkPoolControl `json:",inline"`

	// DSCP is marked on packets in this pool as they leave the host interface.
	// +optional
	// +kubebuilder:validation:XIntOrString
	// +kubebuilder:validation:MaxLength=4
	// +kubebuilder:validation:Pattern=`^(0*([0-9]|[1-5][0-9]|6[0-3])|[dD][fF]|[eE][fF]|[aA][fF][1-4][1-3]|[cC][sS][0-6])$`
	DSCP *numorstring.DSCP `json:"dscp,omitempty" validate:"omitempty"`
}

// +kubebuilder:validation:XValidation:rule="((!has(self.source) || !has(self.source.ports) || size(self.source.ports) == 0) && (!has(self.destination) || !has(self.destination.ports) || size(self.destination.ports) == 0)) || has(self.protocol)",message="protocol must be set when source.ports or destination.ports is specified",reason=FieldValueRequired
// +kubebuilder:validation:XValidation:rule="!has(self.protocol) || (type(self.protocol) == int ? self.protocol in [6, 17, 132] : self.protocol in ['TCP', 'UDP', 'SCTP'])",message="protocol must be one of TCP, UDP, SCTP, or their numbers 6, 17, 132",reason=FieldValueInvalid
// +kubebuilder:validation:XValidation:rule="!has(self.dscp) || type(self.dscp) != int || (self.dscp >= 0 && self.dscp <= 63)",message="dscp must be between 0 and 63",reason=FieldValueInvalid

// TrafficMatch matches headers at the qdisc: egress after SNAT, ingress before DNAT.
type TrafficMatch struct {
	// DSCP class such as AF31, or 0-63; it survives NAT, so prefer it across NAT.
	// +optional
	// +kubebuilder:validation:XIntOrString
	// +kubebuilder:validation:MaxLength=4
	// +kubebuilder:validation:Pattern=`^(0*([0-9]|[1-5][0-9]|6[0-3])|[dD][fF]|[eE][fF]|[aA][fF][1-4][1-3]|[cC][sS][0-6])$`
	DSCP *numorstring.DSCP `json:"dscp,omitempty" validate:"omitempty"`

	// Protocol is TCP, UDP or SCTP, or their numbers. Required with source or destination ports.
	// +optional
	// +kubebuilder:validation:XIntOrString
	// +kubebuilder:validation:MaxLength=8
	Protocol *numorstring.Protocol `json:"protocol,omitempty" validate:"omitempty"`

	// Source entity match.
	//
	// On ingress this is the source on the wire.
	//
	// On egress it is the source after SNAT: right for host traffic, but the node IP for pods.
	//
	// So a pod-CIDR source.nets does not match egress pod traffic.
	// +optional
	Source *EntityMatch `json:"source,omitempty" validate:"omitempty"`

	// Destination entity match.
	//
	// On egress this is the destination the traffic is headed to.
	//
	// On ingress it is the destination before DNAT: the node IP for NodePort traffic.
	//
	// So a pod-CIDR destination.nets does not match ingress service traffic.
	// +optional
	Destination *EntityMatch `json:"destination,omitempty" validate:"omitempty"`
}

// EntityMatch is a source or destination match for QoS classification.
type EntityMatch struct {
	// Nets is a list of CIDRs, which may mix IPv4 and IPv6.
	// +optional
	// +listType=atomic
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:items:MaxLength=64
	// +kubebuilder:validation:XValidation:rule="self.all(n, isIP(n) || (isCIDR(n) && cidr(n) == cidr(n).masked()))",message="nets entries must be IP addresses or strictly masked CIDRs",reason=FieldValueInvalid
	Nets []string `json:"nets,omitempty" validate:"omitempty,max=64,dive,net"`

	// Ports lists port numbers or ranges such as 8080:8090; named ports are rejected.
	// +optional
	// +listType=atomic
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:items:XIntOrString
	// +kubebuilder:validation:items:MaxLength=11
	// +kubebuilder:validation:items:Pattern=`^0*([1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])(:0*([1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5]))?$`
	// +kubebuilder:validation:XValidation:rule="self.all(p, type(p) != int || (p >= 1 && p <= 65535))",message="ports must be between 1 and 65535",reason=FieldValueInvalid
	// +kubebuilder:validation:items:XValidation:rule="type(self) != string || !self.contains(':') || int(self.split(':')[0]) <= int(self.split(':')[1])",message="port ranges must be written low:high",reason=FieldValueInvalid
	Ports []numorstring.Port `json:"ports,omitempty" validate:"omitempty,max=64,dive"`
}

// HostQoSPolicyStatus rolls up every binding of the policy. Per-node detail lives in its
// HostQoSNodeStatus resources.
type HostQoSPolicyStatus struct {
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// AppliedBindings counts the shaped (HostEndpoint, interface) bindings.
	// +optional
	AppliedBindings *int32 `json:"appliedBindings,omitempty"`
	// ErroredBindings counts the bindings in the Error state.
	// +optional
	ErroredBindings *int32 `json:"erroredBindings,omitempty"`
	// CollidedBindings counts the bindings in the HEPCollision or PolicyCollision state.
	// +optional
	CollidedBindings *int32 `json:"collidedBindings,omitempty"`

	// AppliedNodes counts the nodes where every binding is Applied.
	// +optional
	AppliedNodes *int32 `json:"appliedNodes,omitempty"`
	// ErroredNodes counts the nodes with at least one binding in the Error state.
	// +optional
	ErroredNodes *int32 `json:"erroredNodes,omitempty"`
}

// HostQoSPolicyBindingStatus is the state of one (HostEndpoint, interface) binding.
// A HostEndpoint with interfaceName "*" has one binding per interface.
type HostQoSPolicyBindingStatus struct {
	// HostEndpointName is the HostEndpoint the policy selected.
	// +kubebuilder:validation:MaxLength=253
	HostEndpointName string `json:"hostEndpointName"`

	// InterfaceName is the interface this binding shapes, or attempted to.
	// +kubebuilder:validation:MaxLength=64
	InterfaceName string `json:"interfaceName"`

	// State is Applied, Error, PolicyCollision or HEPCollision; never PartiallyApplied.
	State HostQoSPolicyState `json:"state"`

	// Message explains a non-Applied state, as "<HostQoSPolicyCause>: <detail>".
	// +optional
	// +kubebuilder:validation:MaxLength=1024
	Message *string `json:"message,omitempty"`

	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`
}

// HostQoSPolicyState is the state of a binding, or the summary state of a node.
// +kubebuilder:validation:Enum=Applied;PartiallyApplied;Error;PolicyCollision;HEPCollision
type HostQoSPolicyState string

const (
	// HostQoSPolicyStateApplied: the binding is shaped, or every binding on the node is.
	HostQoSPolicyStateApplied HostQoSPolicyState = "Applied"
	// HostQoSPolicyStateError: the binding was refused or failed; its message carries the cause.
	HostQoSPolicyStateError HostQoSPolicyState = "Error"
	// HostQoSPolicyStatePolicyCollision: an alphabetically-earlier policy selects the HostEndpoint.
	HostQoSPolicyStatePolicyCollision HostQoSPolicyState = "PolicyCollision"
	// HostQoSPolicyStateHEPCollision: an alphabetically-earlier HostEndpoint owns the interface.
	HostQoSPolicyStateHEPCollision HostQoSPolicyState = "HEPCollision"
	// HostQoSPolicyStatePartiallyApplied is a node summary only: Applied mixed with other states.
	HostQoSPolicyStatePartiallyApplied HostQoSPolicyState = "PartiallyApplied"
)

// HostType distinguishes a Kubernetes cluster node from a non-cluster host.
// +kubebuilder:validation:Enum=ClusterNode;NonClusterHost
type HostType string

const (
	HostTypeClusterNode    HostType = "ClusterNode"
	HostTypeNonClusterHost HostType = "NonClusterHost"
)

// Condition Type and Reason constants for HostQoSPolicy.status.conditions[].
const (
	HostQoSPolicyConditionReady = "Ready"

	HostQoSPolicyReasonApplied          = "Applied"          // status=True; every binding Applied
	HostQoSPolicyReasonPartiallyApplied = "PartiallyApplied" // status=True; some bindings Applied, some not
	HostQoSPolicyReasonPolicyCollision  = "PolicyCollision"  // status=False; every binding lost a policy collision
	HostQoSPolicyReasonHEPCollision     = "HEPCollision"     // status=False; every binding lost a HEP collision
	HostQoSPolicyReasonError            = "Error"            // status=False; no binding Applied and at least one errored
	HostQoSPolicyReasonNoBindings       = "NoBindings"       // status=Unknown; no HostEndpoint matches the selector
)

// Stable prefixes of HostQoSPolicyBindingStatus.Message, which reads "<Cause>: <detail>".
// Match on the prefix; the detail is free-form.
const (
	HostQoSPolicyCausePolicyCollision            = "PolicyCollision"
	HostQoSPolicyCauseHEPCollision               = "HEPCollision"
	HostQoSPolicyCauseInterfaceExcluded          = "InterfaceExcluded"
	HostQoSPolicyCauseVersionSkew                = "VersionSkew"
	HostQoSPolicyCauseInterfaceDown              = "InterfaceDown"
	HostQoSPolicyCauseInterfaceBondMember        = "InterfaceBondMember" // bind the bond, not its member
	HostQoSPolicyCauseForeignQdisc               = "ForeignQdisc"
	HostQoSPolicyCauseClassIDCollision           = "ClassIDCollision"
	HostQoSPolicyCauseIFBUnavailable             = "IFBUnavailable"
	HostQoSPolicyCauseDataplaneFailure           = "DataplaneFailure"
	HostQoSPolicyCauseCapacityUnknown            = "CapacityUnknown" // linkCapacity omitted, speed undetectable
	HostQoSPolicyCauseReservationsExceedCapacity = "ReservationsExceedCapacity"
)

// NewHostQoSPolicy creates a new (zeroed) HostQoSPolicy struct with the TypeMetadata
// initialised to the current version.
func NewHostQoSPolicy() *HostQoSPolicy {
	return &HostQoSPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindHostQoSPolicy,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewHostQoSPolicyList creates a new (zeroed) HostQoSPolicyList struct with the TypeMetadata
// initialised to the current version.
func NewHostQoSPolicyList() *HostQoSPolicyList {
	return &HostQoSPolicyList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindHostQoSPolicyList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
