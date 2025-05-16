// Copyright (c) 2020-2021 Tigera, Inc. All rights reserved.

package v3

import (
	"github.com/tigera/api/pkg/lib/numorstring"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindPacketCapture     = "PacketCapture"
	KindPacketCaptureList = "PacketCaptureList"
)

// PacketCaptureState represents the state of the PacketCapture
type PacketCaptureState string

const (
	// PacketCaptureStateCapturing represents the active state of a PacketCapture of capturing traffic
	PacketCaptureStateCapturing PacketCaptureState = "Capturing"
	// PacketCaptureStateFinished represents the inactive state of a PacketCapture of not capturing traffic
	PacketCaptureStateFinished = "Finished"
	// PacketCaptureStateScheduled represents the inactive state of a PacketCapture of being
	// scheduled, but not capturing traffic
	PacketCaptureStateScheduled = "Scheduled"
	// PacketCaptureStateError represents the error state of a PacketCapture
	PacketCaptureStateError = "Error"
	// PacketCaptureStateWaitingForTraffic represents the active state of a PacketCapture of capturing from a live
	// interface, but waiting for traffic on that interface
	PacketCaptureStateWaitingForTraffic = "WaitingForTraffic"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PacketCapture contains the configuration for any packet capture.
type PacketCapture struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the PacketCapture.
	Spec PacketCaptureSpec `json:"spec,omitempty"`
	// Status of the PacketCapture
	Status PacketCaptureStatus `json:"status,omitempty"`
}

// PacketCaptureSpec contains the values of the packet capture.
type PacketCaptureSpec struct {
	// The selector is an expression used to pick out the endpoints that the policy should
	// be applied to.  The selector will only match endpoints in the same namespace as the
	// PacketCapture resource.
	//
	// Selector expressions follow this syntax:
	//
	// 	label == "string_literal"  ->  comparison, e.g. my_label == "foo bar"
	// 	label != "string_literal"   ->  not equal; also matches if label is not present
	// 	label in { "a", "b", "c", ... }  ->  true if the value of label X is one of "a", "b", "c"
	// 	label not in { "a", "b", "c", ... }  ->  true if the value of label X is not one of "a", "b", "c"
	// 	has(label_name)  -> True if that label is present
	// 	! expr -> negation of expr
	// 	expr && expr  -> Short-circuit and
	// 	expr || expr  -> Short-circuit or
	// 	( expr ) -> parens for grouping
	// 	all() -> matches all endpoints.
	// 	an empty selector will default to all
	//
	// Label names are allowed to contain alphanumerics, -, _ and /. String literals are more permissive
	// but they do not support escape characters.
	//
	// Examples (with made-up labels):
	//
	// 	type == "webserver" && deployment == "prod"
	// 	type in {"frontend", "backend"}
	// 	deployment != "dev"
	// 	! has(label_name)
	// +kubebuilder:default:="all()"
	Selector string `json:"selector,omitempty" validate:"selector"`

	// The ordered set of filters applied to traffic captured from an interface.  Each rule contains a set of
	// packet match criteria.
	Filters []PacketCaptureRule `json:"filters,omitempty" validate:"omitempty,dive"`

	// Defines the start time from which this PacketCapture will capture packets.
	// If omitted or the value is in the past, the capture will start immediately.
	// If the value is changed to a future time, capture will stop immediately and restart at that time
	// +optional
	// +kubebuilder:validation:Format="date-time"
	StartTime *metav1.Time `json:"startTime,omitempty" validate:"omitempty"`

	// Defines the end time at which this PacketCapture will stop capturing packets.
	// If omitted the capture will continue indefinitely.
	// If the value is changed to the past, capture will stop immediately.
	// +optional
	//+kubebuilder:validation:Format="date-time"
	EndTime *metav1.Time `json:"endTime,omitempty" validate:"omitempty"`
}

// A PacketCaptureRule encapsulates a set of match criteria for traffic captured from an interface.
type PacketCaptureRule struct {
	// Protocol is an optional field that defines a filter for all traffic for
	// a specific IP protocol.
	//
	// Must be one of these string values: "TCP", "UDP", "ICMP", "ICMPv6", "SCTP", "UDPLite"
	// or an integer in the range 1-255.
	Protocol *numorstring.Protocol `json:"protocol,omitempty" validate:"omitempty"`

	// Ports is an optional field that defines a filter for all traffic that has a
	// source or destination port that matches one of these ranges/values. This value is a
	// list of integers or strings that represent ranges of ports.
	Ports []numorstring.Port `json:"ports,omitempty" validate:"omitempty,dive"`
}

// PacketCaptureStatus describes the files that have been captured, for a given PacketCapture, on each node
// that generates packet capture files
type PacketCaptureStatus struct {
	Files []PacketCaptureFile `json:"files,omitempty"`
}

// PacketCaptureFile describes files generated by a PacketCapture. It describes the location of the packet capture files
// that is identified via a node, its directory and the file names generated.
type PacketCaptureFile struct {
	// Node identifies with a physical node from the cluster via its hostname
	Node string `json:"node,omitempty" validate:"omitempty"`

	// Directory represents the path inside the calico-node container for the the generated files
	Directory string `json:"directory,omitempty" validate:"omitempty"`

	// FileNames represents the name of the generated file for a PacketCapture ordered alphanumerically.
	// The active packet capture file will be identified using the following schema:
	// "{workload endpoint name}_{host network interface}.pcap" .
	// Rotated capture files name will contain an index matching the rotation timestamp.
	FileNames []string `json:"fileNames,omitempty" validate:"omitempty,dive"`

	// Determines whether a PacketCapture is capturing traffic from any interface
	// attached to the current node

	// +kubebuilder:validation:Enum=Capturing;Finished;Scheduled;Error;WaitingForTraffic
	State *PacketCaptureState `json:"state,omitempty" validate:"omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PacketCaptureList contains a list of PacketCapture resources.
type PacketCaptureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []PacketCapture `json:"items"`
}

// NewPacketCapture creates a new (zeroed) PacketCapture struct with the TypeMetadata initialised to the current
// version.
func NewPacketCapture() *PacketCapture {
	return &PacketCapture{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindPacketCapture,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewPacketCaptureList creates a new (zeroed) PacketCaptureList struct with the TypeMetadata initialised to the current
// version.
func NewPacketCaptureList() *PacketCaptureList {
	return &PacketCaptureList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindPacketCaptureList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
