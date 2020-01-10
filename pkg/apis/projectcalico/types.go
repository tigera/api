// Copyright (c) 2019 Tigera, Inc. All rights reserved.

package projectcalico

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalReportType contains the configuration for a non-namespaced report type.
type GlobalReportType struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the GlobalReport.
	Spec ReportTypeSpec `json:"spec,omitempty"`
}

// ReportTypeSpec contains the various templates, and configuration used to render a specific type of report.
type ReportTypeSpec struct {
	// The summary template, explicitly used by the UI to render a summary version of the report. This should render
	// to json containing a sets of widgets that the UI can use to render the summary. The rendered data is returned
	// on the list query of the reports.
	UISummaryTemplate ReportTemplate `json:"uiSummaryTemplate,omitempty" validate:"required"`

	// The set of templates used to render the report for downloads.
	DownloadTemplates []ReportTemplate `json:"downloadTemplates,omitempty" validate:"dive"`

	// Whether to include endpoint data in the report. The actual endpoints included may be filtered by the Report,
	// but will otherwise contain the full set of endpoints.
	IncludeEndpointData bool `json:"includeEndpointData,omitempty"`

	// Whether to include endpoint-to-endpoint flow log data in the report.
	IncludeEndpointFlowLogData bool `json:"includeEndpointFlowLogData,omitempty"`

	// What audit log data should be included in the report. If not specified, the report will contain no audit log
	// data. The selection may be further filtered by the Report.
	AuditEventsSelection *AuditEventsSelection `json:"auditEventsSelection,omitempty" validate:"omitempty"`

	// Whether to include the full cis benchmark test results in the report.
	IncludeCISBenchmarkData bool `json:"includeCISBenchmarkData,omitempty"`
}

// ReportTemplate defines a template used to render a report into downloadable or UI compatible format.
type ReportTemplate struct {
	// The name of this template. This should be unique across all template names within a ReportType. This will be used
	// by the UI as the suffix of the downloadable file name.
	Name string `json:"name,omitempty" validate:"name,required"`

	// A user-facing description of the template.
	Description string `json:"description,omitempty"`

	// The base-64 encoded go template used to render the report data.
	Template string `json:"template,omitempty" validate:"required"`
}

// AuditEventsSelection defines which set of resources should be audited.
type AuditEventsSelection struct {
	// Resources lists the resources that will be included in the audit logs in the ReportData.  Blank fields in the
	// listed ResourceID structs are treated as wildcards.
	Resources []AuditResource `json:"resources,omitempty" validate:"omitempty"`
}

// ResourceID is used to identify a resource instance in the report data.
type ResourceID struct {
	metav1.TypeMeta `json:",inline"`
	Name            string    `json:"name,omitempty" validate:"omitempty"`
	Namespace       string    `json:"namespace,omitempty" validate:"omitempty"`
	UUID            types.UID `json:"uuid,omitempty" validate:"omitempty"`
}

// AuditResource is used to filter Audit events in the Report configuration.
//
// An empty field value indicates a wildcard. For example, if Resource is set to "networkpolicies" and all other
// fields are blank then this filter would include all NetworkPolicy resources across all namespaces, and would include
// both Calico and Kubernetes resource types.
type AuditResource struct {
	// The resource type. The format is the lowercase plural as used in audit event selection and RBAC configuration.
	Resource string `json:"resource,omitempty" validate:"omitempty"`

	// APIGroup is the name of the API group that contains the referred object (e.g. projectcalico.org).
	APIGroup string `json:"apiGroup,omitempty" validate:"omitempty"`

	// APIVersion is the version of the API group that contains the referred object (e.g. v3).
	APIVersion string `json:"apiVersion,omitempty" validate:"omitempty"`

	// The resource name.
	Name string `json:"name,omitempty" validate:"omitempty"`

	// The resource namespace.
	Namespace string `json:"namespace,omitempty" validate:"omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalReportTypeList contains a list of GlobalReportType resources.
type GlobalReportTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GlobalReportType `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LicenseKey contains the Tigera CNX license key for the cluster.
type LicenseKey struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.  This resource is a singleton, always named "default".
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the LicenseKey.
	Spec LicenseKeySpec `json:"spec,omitempty"`
}

// LicenseKeySpec contains the license key itself.
type LicenseKeySpec struct {
	// Token is the JWT containing the license claims
	Token string `json:"token" yaml:"token"`
	// Certificate is used to validate the token.
	Certificate string `json:"certificate,omitempty" yaml:"certificate" validate:"omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LicenseKeyList contains a list of LicenseKey resources
// (even though there should only be one).
type LicenseKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []LicenseKey `json:"items"`
}

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
}

type ManagedClusterStatus struct {
	Conditions []ManagedClusterStatusCondition `json:"conditions"`
}

// Condition contains various status information
type ManagedClusterStatusCondition struct {
	Message string                    `json:"message,omitempty"`
	Reason  string                    `json:"reason,omitempty"`
	Status  ManagedClusterStatusValue `json:"status"`
	Type    ManagedClusterStatusType  `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManagedClusterList contains a list of ManagedCluster resources.
type ManagedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ManagedCluster `json:"items"`
}