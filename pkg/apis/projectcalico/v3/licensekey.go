// Copyright (c) 2018,2021 Tigera, Inc. All rights reserved.

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
	KindLicenseKey     = "LicenseKey"
	KindLicenseKeyList = "LicenseKeyList"
)

// +kubebuilder:validation:Enum=CloudCommunity;CloudStarter;CloudPro;Enterprise
type LicensePackageType string

const (
	CloudCommunity LicensePackageType = "CloudCommunity"
	CloudStarter   LicensePackageType = "CloudStarter"
	CloudPro       LicensePackageType = "CloudPro"
	Enterprise     LicensePackageType = "Enterprise"
)

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Expiry",type="string",JSONPath=".status.expiry"
// +kubebuilder:printcolumn:name="Grace-Period",type="string",JSONPath=".status.gracePeriod"
// +kubebuilder:printcolumn:name="Max-Nodes",type="integer",JSONPath=".status.maxnodes"
// +kubebuilder:printcolumn:name="Valid",type="string",JSONPath=".status.conditions[?(@.type=='Valid')].status"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// LicenseKey contains the Calico Enterprise license key for the cluster.
type LicenseKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              LicenseKeySpec `json:"spec"`

	// +optional
	Status LicenseKeyStatus `json:"status"`
}

// LicenseKeySpec contains the license key itself.
type LicenseKeySpec struct {
	// Token is the JWT containing the license claims
	Token string `json:"token" yaml:"token"`

	// Certificate is used to validate the token.
	Certificate string `json:"certificate,omitempty" yaml:"certificate" validate:"omitempty"`
}

// LicenseKeyStatus contains the license key information.
type LicenseKeyStatus struct {
	// Expiry is the expiry date of License
	// +nullable
	// +optional
	Expiry metav1.Time `json:"expiry" yaml:"expiry"`

	// GracePeriod is how long after expiry the license remains functional (e.g. "90d")
	// +optional
	GracePeriod string `json:"gracePeriod,omitempty" yaml:"gracePeriod" validate:"omitempty"`

	// Maximum Number of Allowed Nodes
	// +optional
	MaxNodes int `json:"maxnodes,omitempty" yaml:"maxnodes" validate:"omitempty"`

	// License package defines type of Calico license that is being enforced
	// +optional
	Package LicensePackageType `json:"package,omitempty" yaml:"package" validate:"omitempty"`

	// List of features that are available via the applied license
	// +optional
	// +listType=atomic
	Features []string `json:"features,omitempty" yaml:"features" validate:"omitempty"`

	// Conditions is a list of conditions related to the license key. This can be used to indicate if the license is valid, expired, etc.
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" yaml:"conditions" validate:"omitempty"`
}

const (
	LicenseKeyConditionValid = "Valid"
)

const (
	LicenseKeyReasonValidLicense   = "ValidLicense"
	LicenseKeyReasonExpiredLicense = "ExpiredLicense"
	LicenseKeyReasonInvalidLicense = "InvalidLicense"
)

// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LicenseKeyList contains a list of LicenseKey resources
// (even though there should only be one).
type LicenseKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []LicenseKey `json:"items"`
}

// New LicenseKey creates a new (zeroed) LicenseKey struct with the TypeMetadata
// initialized to the current version.
func NewLicenseKey() *LicenseKey {
	return &LicenseKey{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindLicenseKey,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewLicenseKeyList creates a new (zeroed) LicenseKeyList struct with the TypeMetadata
// initialized to the current version.
func NewLicenseKeyList() *LicenseKeyList {
	return &LicenseKeyList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindLicenseKeyList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
