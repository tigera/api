// Copyright (c) 2023 Tigera, Inc. All rights reserved.

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
	KindEgressGatewayPolicy     = "EgressGatewayPolicy"
	KindEgressGatewayPolicyList = "EgressGatewayPolicyList"
)

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EgressGatewayPolicyList is a list of EgressGatewayPolicy resources.
type EgressGatewayPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []EgressGatewayPolicy `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EgressGatewayPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec EgressGatewayPolicySpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// EgressGatewayPolicySpec contains the egress policy rules for each destination network
type EgressGatewayPolicySpec struct {
	// The ordered set of Egress Gateway Policies to define how traffic exit a cluster
	Rules []EgressGatewayRule `json:"rules,omitempty" validate:"required"`
}

// EgressGatewayRule defines an Egress Gateway to reach a destination network
type EgressGatewayRule struct {
	// The destination network that can be reached via egress gateway.
	// If no destination is set, the default route, 0.0.0.0/0, is used instead.
	// +optional
	Destination *EgressGatewayPolicyDestinationSpec `json:"destination,omitempty" validate:"omitempty"`

	// The description of the EgressGatewayPolicy rule.
	// +optional
	Description string `json:"description,omitempty" validate:"omitempty,uiDescription"`

	// Gateway specifies the egress gateway that should be used for the specified destination.
	// If no gateway is set then the destination is routed normally rather than via an egress gateway.
	// +optional
	Gateway *EgressSpec `json:"gateway,omitempty" validate:"omitempty"`

	// GatewayPreference specifies which egress gateways to use. If set to PreferNodeLocal, egress gateways in the same node as
	// the client will be used if available. Otherwise all the active egress gateways will be used.
	// +kubebuilder:default=None
	// +optional
	GatewayPreference *GatewayPreferenceType `json:"gatewayPreference,omitempty" validate:"omitempty,oneof=None,PreferNodeLocal"`
}

// DestinationSpec define a destination network that can be reached via an egress gateway
type EgressGatewayPolicyDestinationSpec struct {
	// The destination network CIDR.
	CIDR string `json:"cidr,omitempty" validate:"omitempty,net"`
}

// New EgressGatewayPolicy creates a new (zeroed) EgressGatewayPolicy struct with the TypeMetadata
// initialized to the current version.
func NewEgressGatewayPolicy() *EgressGatewayPolicy {
	return &EgressGatewayPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindEgressGatewayPolicy,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// +kubebuilder:validation:Enum=None;PreferNodeLocal
type GatewayPreferenceType string

const (
	GatewayPreferenceNone      GatewayPreferenceType = "None"
	GatewayPreferenceNodeLocal GatewayPreferenceType = "PreferNodeLocal"
)
