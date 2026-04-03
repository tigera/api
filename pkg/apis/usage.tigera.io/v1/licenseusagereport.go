// Copyright (c) 2026 Tigera, Inc. All rights reserved.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:resource:scope=Cluster
type LicenseUsageReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              LicenseUsageReportSpec `json:"spec"`
}

type LicenseUsageReportSpec struct {
	// The base64-encoded JSON data for this report. The data represents an interval of time when license usage was
	// monitored in the cluster, along with data that binds the report to its cluster context.
	ReportData string `json:"reportData"`

	HMAC string `json:"hmac"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LicenseUsageReportList contains a list of ManagedCluster resources.
type LicenseUsageReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []LicenseUsageReport `json:"items"`
}
