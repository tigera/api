// Copyright (c) 2019,2021 Tigera, Inc. All rights reserved.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindGlobalReport     = "GlobalReport"
	KindGlobalReportList = "GlobalReportList"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalReport contains the configuration for a non-namespaced Report.
type GlobalReport struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the GlobalReport.
	Spec   ReportSpec   `json:"spec,omitempty"`
	Status ReportStatus `json:"status,omitempty"`
}

// ReportSpec contains the values of the GlobalReport.
type ReportSpec struct {
	// The name of the report type.
	ReportType string `json:"reportType" validate:"name,required"`

	// Endpoints is used to specify which endpoints are in-scope and stored in the generated report data.
	// Only used if endpoints data and/or audit logs are gathered in the report. If omitted, treated as everything
	// in-scope.
	Endpoints *EndpointsSelection `json:"endpoints,omitempty" validate:"omitempty,selector"`

	// The report schedule specified in cron format. This specifies both the start and end times of each report,
	// where the end time of one report becomes the start time of the next report.
	// Separate jobs are created to generate a report, and the job generates the report data from archived audit
	// and traffic data. To ensure this data is actually archived, the jobs to generate each report starts at a
	// configurable time *after* the end time of the report that is being generated. The default job start delay is
	// 30m, but is configurable through the compliance-controller environments.
	// The cron format has minute accuracy, but only up to two values may be configured for the minute column which
	// means you may only have at most two reports for each hour period.
	Schedule string `json:"schedule,omitempty" validate:"omitempty"`

	// The node selector used to specify which nodes the report job may be scheduled on.
	JobNodeSelector map[string]string `json:"jobNodeSelector,omitempty" validate:"omitempty"`

	// This flag tells the controller to suspend subsequent jobs for generating reports, it does not apply to already
	// started jobs. If jobs are resumed then the controller will start creating jobs for any reports that were missed
	// while the job was suspended.
	Suspend *bool `json:"suspend,omitempty" validate:"omitempty"`

	// This field contain all the parameters for configuring a CIS benchmark report.
	CIS *CISBenchmarkParams `json:"cis,omitempty" validate:"omitempty"`
}

// CISBenchmarkParams contains the parameters for configuring a CIS benchmark report.
type CISBenchmarkParams struct {
	// Specifies if the report should also show results for scored/not-scored tests.
	IncludeUnscoredTests bool `json:"includeUnscoredTests,omitempty"`

	// Configure the number of top failed tests to show up on the report.
	NumFailedTests *int `json:"numFailedTests,omitempty" validate:"gt=0"`

	// Benchmark results filters. The first matching set of filters is applied to each set of benchmark results.
	// If there are no matching filters, the full set of benchmark results will be included in the report.
	ResultsFilters []CISBenchmarkFilter `json:"resultsFilters,omitempty"`

	// Interpretted as a percentage to indicate at what levels of passing tests a node should be considered
	// HIGH, MED, and LOW.
	// - If >= HighThreshold flag as high
	// - Otherwise, if > MedThreshold flag as med
	// - Otherwise flag as low.
	HighThreshold *int `json:"highThreshold,omitempty" validate:"gte=0,lte=100,gtfield=MedThreshold"`
	MedThreshold  *int `json:"medThreshold,omitempty" validate:"gte=0,lte=100"`
}

// CISBenchmarkFilter provides filters for a set of benchmarks that match particular selection criteria.
type CISBenchmarkFilter struct {
	// BenchmarkSelection specifies which benchmarks this filter applies to. If not specified, applies to all.
	BenchmarkSelection *CISBenchmarkSelection `json:"benchmarkSelection,omitempty" validate:"omitempty"`

	// Exclude is an array of test indices to exclude from the report.
	Exclude []string `json:"exclude,omitempty"`

	// Include is an array of test indices to show in the report.
	// Is additive if IncludeUnscoredTests is true.
	// Takes precedence over Exclude.
	Include []string `json:"include,omitempty"`
}

// CISBenchmarkSelection selects a particular set of benchmarks.
type CISBenchmarkSelection struct {
	// KubernetesVersion is used select nodes that are running a specific version of kubelet. The full version need not
	// be fully specified down to the patch level, in which case the significant parts of the version are matched.
	// e.g. "1.0" will match versions "1.0.1" and "1.0.2"
	// If not specified, matches all versions.
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
}

// ReportStatus contains the status of the automated report generation.
type ReportStatus struct {
	// The configured report jobs that have completed successfully.
	LastSuccessfulReportJobs []CompletedReportJob `json:"lastSuccessfulReportJobs,omitempty"`

	// The configured report jobs that have failed.
	LastFailedReportJobs []CompletedReportJob `json:"lastFailedReportJobs,omitempty"`

	// The set of active report jobs.
	ActiveReportJobs []ReportJob `json:"activeReportJobs,omitempty"`

	// The last scheduled report job.
	LastScheduledReportJob *ReportJob `json:"lastScheduledReportJob,omitempty"`
}

// ReportJob contains
type ReportJob struct {
	// The start time of the report.
	Start metav1.Time `json:"start"`

	// The end time of the report.
	End metav1.Time `json:"end"`

	// A reference to the report creation job if known.
	Job *corev1.ObjectReference `json:"job"`
}

// CompletedReportJob augments the ReportJob with completion details.
type CompletedReportJob struct {
	ReportJob `json:",inline"`

	// The time the report job completed.
	JobCompletionTime *metav1.Time `json:"jobCompletionTime,omitempty"`
}

// EndpointsSelection is a set of selectors used to select the endpoints that are considered to be in-scope for the
// report. An empty selector is equivalent to all(). All three selectors are ANDed together.
type EndpointsSelection struct {
	// Selector, selects endpoints by endpoint labels. If omitted, all endpoints are included in the report
	// data.
	Selector string `json:"selector,omitempty" validate:"omitempty,selector"`

	// Namespace match restricts endpoint selection to those in the selected namespaces.
	Namespaces *NamesAndLabelsMatch `json:"namespaces,omitempty" validate:"omitempty"`

	// ServiceAccount match restricts endpoint selection to those in the selected service accounts.
	ServiceAccounts *NamesAndLabelsMatch `json:"serviceAccounts,omitempty" validate:"omitempty"`
}

// NamesAndLabelsMatch is used to specify resource matches using both label and name selection.
type NamesAndLabelsMatch struct {
	// Names is an optional field that specifies a set of resources by name.
	Names []string `json:"names,omitempty" validate:"omitempty"`

	// Selector is an optional field that selects a set of resources by label.
	// If both Names and Selector are specified then they are AND'ed.
	Selector string `json:"selector,omitempty" validate:"omitempty,selector"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalReportList contains a list of GlobalReport resources.
type GlobalReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GlobalReport `json:"items"`
}

// NewGlobalReport creates a new (zeroed) GlobalReport struct with the TypeMetadata
// initialized to the current version.
func NewGlobalReport() *GlobalReport {
	return &GlobalReport{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindGlobalReport,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewGlobalReportList creates a new (zeroed) GlobalReportList struct with the TypeMetadata
// initialized to the current version.
func NewGlobalReportList() *GlobalReportList {
	return &GlobalReportList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindGlobalReportList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
