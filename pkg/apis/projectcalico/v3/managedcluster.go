// Copyright (c) 2020-2021 Tigera, Inc. All rights reserved.

package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindManagedCluster     = "ManagedCluster"
	KindManagedClusterList = "ManagedClusterList"
)

type ManagedClusterStatusType string
type ManagedClusterStatusValue string

const (
	// Status for Type ManagedClusterConnected will be Unknown when ManagedCluster is created,
	// True when ManagedCluster is connected to ManagementCluster via tunnel,
	// False when the tunnel drops
	ManagedClusterStatusTypeConnected ManagedClusterStatusType  = "ManagedClusterConnected"
	ManagedClusterStatusValueUnknown  ManagedClusterStatusValue = "Unknown"
	ManagedClusterStatusValueTrue     ManagedClusterStatusValue = "True"
	ManagedClusterStatusValueFalse    ManagedClusterStatusValue = "False"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManagedCluster represents a cluster that is being managed by the multi-cluster
// management plane. This object configures how Tigera multi-cluster management
// components communicate with the corresponding cluster.
type ManagedCluster struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the ManagedCluster.
	Spec ManagedClusterSpec `json:"spec,omitempty"`
	// Status of the ManagedCluster
	Status ManagedClusterStatus `json:"status,omitempty"`
}

// ManagedClusterSpec contains the specification of a ManagedCluster resource.
type ManagedClusterSpec struct {
	// Field to store dynamically generated manifest for installing component into
	// the actual application cluster corresponding to this Managed Cluster
	InstallationManifest string `json:"installationManifest,omitempty"`
	// The namespace of the managed cluster's operator. This value is used in
	// the generation of the InstallationManifest.
	OperatorNamespace string `json:"operatorNamespace,omitempty"`
}

type ManagedClusterStatus struct {
	Conditions []ManagedClusterStatusCondition `json:"conditions,omitempty"`
}

// Condition contains various status information
type ManagedClusterStatusCondition struct {
	Message string                    `json:"message,omitempty"`
	Reason  string                    `json:"reason,omitempty"`
	Status  ManagedClusterStatusValue `json:"status"`
	Type    ManagedClusterStatusType  `json:"type"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManagedClusterList contains a list of ManagedCluster resources.
type ManagedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ManagedCluster `json:"items"`
}

// NewManagedCluster creates a new (zeroed) ManagedCluster struct with the TypeMetadata initialised to the current
// version.
func NewManagedCluster() *ManagedCluster {
	return &ManagedCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindManagedCluster,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewManagedClusterList creates a new (zeroed) ManagedClusterList struct with the TypeMetadata initialised to the current
// version.
func NewManagedClusterList() *ManagedClusterList {
	return &ManagedClusterList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindManagedClusterList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
