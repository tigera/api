// Copyright (c) 2024 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindBFDConfiguration     = "BFDConfiguration"
	KindBFDConfigurationList = "BFDConfigurationList"
)

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BFDConfigurationList is a list of BFDConfiguration resources.
type BFDConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []BFDConfiguration `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BFDConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec BFDConfigurationSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// BFDConfigurationSpec contains the specification for a BFDConfiguration resource.
type BFDConfigurationSpec struct {
	NodeSelector string `json:"nodeSelector,omitempty"`

	Interfaces []BFDInterface `json:"interfaces,omitempty"`
}

// BFDInterface contains per-interface parameters for BFD failure detection.
type BFDInterface struct {
	// MatchPattern is a pattern to match one or more interfaces.
	// Supports exact interface names, match on interface prefixes (e.g., “eth*”),
	// or “*” to select all interfaces on the selected node(s).
	// +required
	MatchPattern string `json:"matchPattern,omitempty"`

	// MinimumRecvInterval is the minimum interval between received BFD packets. Must be a whole number of milliseconds greater than 0.
	// +optional
	// +kubebuilder:default="10ms"
	MinimumRecvInterval *metav1.Duration `json:"minimumRecvInterval,omitempty"`

	// MinimumSendInterval is the minimum interval between transmitted BFD packets. Must be a whole number of milliseconds greater than 0.
	// +optional
	// +kubebuilder:default="100ms"
	MinimumSendInterval *metav1.Duration `json:"minimumSendInterval,omitempty"`

	// IdleSendInterval is the interval between transmitted BFD packets when the BFD peer is idle. Must be a whole number of milliseconds greater than 0.
	// +optional
	// +kubebuilder:default="1m"
	IdleSendInterval *metav1.Duration `json:"idleSendInterval,omitempty"`

	// Multiplier is the number of intervals that must pass without receiving a BFD packet before the peer is considered down.
	// +optional
	// +kubebuilder:default=5
	Multiplier int `json:"multiplier,omitempty"`
}

// NewBFDConfiguration creates a new (zeroed) BFDConfiguration struct with the TypeMetadata initialised to the current
// version.
func NewBFDConfiguration() *BFDConfiguration {
	return &BFDConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindBFDConfiguration,
			APIVersion: GroupVersionCurrent,
		},
	}
}
