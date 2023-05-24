// Copyright (c) 2019 Tigera, Inc. All rights reserved.

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
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/apis/audit"
)

// ReportData contains the aggregated data available for rendering in report templates. The data available is dependent
// on the selector/report configuration.
//
// The data is stored directly in elastic search. To reduce nesting and simplify indexing, all summary values are
// contained at the top level.
type ReportData struct {
	ReportName     string         `json:"reportName"`
	ReportTypeName string         `json:"reportTypeName"`
	ReportSpec     ReportSpec     `json:"reportSpec"`
	ReportTypeSpec ReportTypeSpec `json:"reportTypeSpec"`
	StartTime      metav1.Time    `json:"startTime"`
	EndTime        metav1.Time    `json:"endTime"`
	GenerationTime metav1.Time    `json:"generationTime"`

	// The set of in-scope endpoints.
	Endpoints []EndpointsReportEndpoint `json:"endpoints,omitempty"`

	// Endpoint stats in a reporting period.
	EndpointsSummary EndpointsSummary `json:"endpointsSummary,omitempty"`

	// The set of namespaces containing in-scope endpoints.
	Namespaces []EndpointsReportNamespace `json:"namespaces,omitempty"`

	// Endpoint stats for given namespaces in a reporting period.
	NamespacesSummary EndpointsSummary `json:"namespacesSummary,omitempty"`

	// The set of services containing in-scope endpoints.
	Services []EndpointsReportService `json:"services,omitempty"`

	// Endpoint stats for services in a reporting period.
	ServicesSummary EndpointsSummary `json:"servicesSummary,omitempty"`

	// The time-ordered set of in-scope audit events that occurred within the reporting interval.
	AuditEvents []audit.Event `json:"auditEvents,omitempty"`

	// Audit log stats in a reporting period.
	AuditSummary AuditSummary `json:"auditSummary,omitempty"`

	// Flows for in-scope endpoints that have been recorded within the reporting period.
	Flows []EndpointsReportFlow `json:"flows,omitempty"`

	// CISBenchmark contains the per-node results of a cis benchmark scan.
	CISBenchmark []CISBenchmarkNode `json:"cisBenchmark,omitempty"`

	// CISBenchmarkSummary high level test results.
	CISBenchmarkSummary CISBenchmarkSummary `json:"cisBenchmarkSummary,omitempty"`
}

// This tracks different statistics for Endpoints, Summary and Services.
type EndpointsSummary struct {
	// For endpoints: the total number of in-scope endpoints.
	//    Namespaces: the total number of namespaces containing in-scope endpoints.
	//      Services: the total number of services containing in-scope endpoints.
	//
	// Source: Calculated from pod/wep, hep, namespace and service account labels.
	NumTotal int `json:"numTotal,omitempty"`

	// For endpoints: the total number of service accounts for in-scope endpoints.
	//    Namespaces: n/a.
	//      Services: n/a.
	NumServiceAccounts int `json:"numServiceAccounts,omitempty"`

	// For endpoints: the number of in-scope endpoints that were ingress protected during the reporting interval.
	//    Namespaces: the number of namespaces whose in-scope endpoints were ingress protected during
	//                the reporting interval.
	//      Services: the number of services whose in-scope endpoints were ingress protected during the reporting
	//                interval.
	//
	// See below for defn of ingress-protected.
	NumIngressProtected int `json:"numIngressProtected,omitempty"`

	// For endpoints: the number of in-scope endpoints that were egress protected during the reporting interval.
	//    Namespaces: the number of namespaces whose in-scope endpoints were egress protected during the reporting
	//                interval.
	//
	// See below for defn of egress-protected.
	NumEgressProtected int `json:"numEgressProtected,omitempty"`

	// For endpoints: the number of in-scope endpoints whose policy would allow ingress traffic from the Internet
	//                for *any* period within the reporting interval.
	//                (See below for how this is calculated for an endpoint.)
	//    Namespaces: the number of namespaces that contained in-scope endpoints that would allow ingress traffic
	//                from the Internet for *any* period within the reporting interval.
	//      Services: the number of services that contained in-scope endpoints that would allow ingress traffic
	//                from the Internet for *any* period within the reporting interval.
	NumIngressFromInternet int `json:"numIngressFromInternet,omitempty"`

	// For endpoints: the number of in-scope endpoints whose policy would allow egress traffic to the Internet
	//                for *any* period within the reporting interval.
	//                (See below for how this is calculated for an endpoint.)
	//    Namespaces: the number of namespaces that contained in-scope endpoints that would allow egress traffic
	//                to the Internet for *any* period within the reporting interval.
	NumEgressToInternet int `json:"numEgressToInternet,omitempty"`

	// For endpoints: the number of in-scope endpoints whose policy would allow ingress traffic from a
	//                different namespace for *any* period within the reporting interval.
	//                (See below for how this is calculated for an endpoint.)
	//    Namespaces: the number of namespaces that contained in-scope endpoints that would allow ingress
	//                traffic from another namespace for *any* period within the reporting interval.
	//      Services: the number of services that contained in-scope endpoints that would allow ingress
	//                traffic from another namespace for *any* period within the reporting interval.
	NumIngressFromOtherNamespace int `json:"numIngressFromOtherNamespace,omitempty"`

	// For endpoints: the number of in-scope endpoints whose policy would allow ingress traffic from
	//                a different namespace for *any* period within the reporting interval.
	//                (See below for how this is calculated for an endpoint.)
	//    Namespaces: the number of namespaces that contained in-scope endpoints that would allow egress
	//                traffic to another namespace for *any* period within the reporting interval.
	NumEgressToOtherNamespace int `json:"numEgressToOtherNamespace,omitempty"`

	// For endpoints: the number of in-scope endpoints that were envoy-enabled within the reporting interval.
	//    Namespaces: the number of namespaces whose in-scope endpoints were always Envoy-enabled
	//      Services: the number of services whose in-scope endpoints were always Envoy-enabled
	//
	// See below for defn of envoy-enabled
	NumEnvoyEnabled int `json:"numEnvoyEnabled,omitempty"`
}

type AuditSummary struct {
	// The total number of in-scope audit logs.
	NumTotal int `json:"numTotal,omitempty"`

	// The number of in-scope audit log create events.
	NumCreate int `json:"numCreate,omitempty"`

	// The number of in-scope audit log patch or replace events.
	NumModify int `json:"numModify,omitempty"`

	// The number of in-scope audit log delete events.
	NumDelete int `json:"numDelete,omitempty"`
}

type EndpointsReportEndpoint struct {
	Endpoint ResourceID `json:"endpoint,omitempty"`

	// Whether ingress traffic to this endpoint was always protected during the reporting interval.
	//
	// Ingress protection is defined as denying ingress traffic unless explicitly whitelisted. This is translated as
	// the endpoint having some explicit ingress policy applied to it.
	//
	// Source: Calculated from the set of ingress policies that apply to each endpoint.
	//
	// Set to:
	// - false if there are no ingress policies applied to the endpoint at any point during the reporting interval.
	// - true otherwise.
	//
	// Note: Policy is not inspected for protection bypass: for example match-all-and-allow rules which would effectively
	//       short-circuit the default tier-drop behavior, in this case the match-all-and-allow would be considered to be
	//       an explicit whitelist of all traffic. We could include simplistic all-match rules and check that they
	//       don't result in an allow. To check for more circuitous match-all allows is much trickier (e.g. you have one
	//       rule that allows for src!=1.2.3.0/24 and another rule that allows for src==1.2.3.0/24, which combined
	//       is essentially an allow-all).
	IngressProtected bool `json:"ingressProtected,omitempty"`

	// Whether egress traffic to this endpoint was always protected during the reporting interval.
	//
	// Egress protection is defined as denying egress traffic unless explicitly whitelisted. This is translated as
	// the endpoint having some explicit egress policy applied to it.
	//
	// Source: Calculated from the set of egress policies that apply to each endpoint.
	//
	// Set to:
	// - false if there are no egress policies applied to the endpoint at any point during the reporting interval.
	// - true otherwise.
	//
	// Note: Policy is not inspected for protection bypass: for example match-all-and-allow rules which would effectively
	//       short-circuit the default tier-drop behavior, in this case the match-all-and-allow would be considered to be
	//       an explicit whitelist of all traffic. We could include simplistic all-match rules and check that they
	//       don't result in an allow. To check for more circuitous match-all allows is much trickier (e.g. you have one
	//       rule that allows for src!=1.2.3.0/24 and another rule that allows for src==1.2.3.0/24, which combined
	//       is essentially an allow-all). Similarly, policy that only contains pass rules would still count as being
	//       protected.
	EgressProtected bool `json:"egressProtected,omitempty"`

	// Whether the matching policy has any ingress allow rules from a public IP address (as defined by the complement of
	// the private addresses; private addresses default to those defined in RFC 1918, but may also be configured separately).
	//
	// Source: Calculated from the policies applied to the endpoint. The ingress allow rules in each policy are checked
	//         to determine if any CIDR specified in the rule, either directly or through a matching network set, is an
	//         internet address. Endpoint addresses are not included - therefore ingress from a pod that has a public
	//         IP address will not be considered as “from internet”.
	//
	// Note: This is a simplification since it does not examine the policies to determine if it's actually possible to
	//       hit one of these allow rules (e.g. a previous rule may be a match-all-deny).
	IngressFromInternet bool `json:"ingressFromInternet,omitempty"`

	// Whether the matching policy has any egress allow rules to a public IP address (as defined by the complement of
	// the private addresses; private addresses default to those defined in RFC 1918, but may also be configured separately).
	//
	// Source: Calculated from the policies applied to the endpoint. The egress allow rules in each policy are checked
	//         to determine if any CIDR specified in the rule, either directly or through a matching network set, is an
	//         internet address. Endpoint addresses are not included - therefore egress to a pod that has a public
	//         IP address will not be considered as “to internet”.
	//
	// Note 1: This is a simplification since it does not examine the policies to determine if it's actually possible to
	//         hit one of these allow rules (e.g. a previous rule may be a match-all-deny).
	EgressToInternet bool `json:"egressToInternet,omitempty"`

	// Whether the matching policy has any ingress allow rules from another namespace.
	//
	// Source: Calculated from the policies applied to the endpoint.
	//
	// Set to true if:
	// - this is a pod (i.e. namespaced) with an applied GlobalNetworkPolicy with an ingress allow rule with no CIDR match.
	// - this is a pod with an applied NetworkPolicy with an ingress allow rule with a non-empty NamespaceSelector.
	//
	// Note: This is a simplification since it does not examine the policies to determine if it's actually possible to
	//       hit one of these allow rules (e.g. a previous rule may be a match-all-deny, or endpoint selector may not
	//       match any endpoints within the namespace).
	IngressFromOtherNamespace bool `json:"ingressFromOtherNamespace,omitempty"`

	// Whether the matching policy has any egress allow rules to another namespace.
	//
	// Source: Calculated from the policies applied to the endpoint.
	//
	// Set to true if:
	// - this is a pod endpoint (i.e. namespaced) matches a GlobalNetworkPolicy with an egress allow rule with no CIDR match.
	// - this is a pod endpoint which matches a NetworkPolicy with an egress allow rule with a non-empty NamespaceSelector.
	//
	// Note: This is a simplification since it does not examine the policies to determine if it's actually possible to
	//       hit one of these allow rules (e.g. a previous rule may be a match-all-deny, or endpoint selector may not
	//       match any endpoints within the namespace).
	EgressToOtherNamespace bool `json:"egressToOtherNamespace,omitempty"`

	// Whether this pod is envoy-enabled. This is simply an indicator of whether an Envoy container is running within the pod.
	// Provided Istio is configured appropriately, this can provide a simplistic determination of whether the pod is mTLS
	// enabled.
	//
	// Source: Pod spec.
	//
	// Set to:
	// - true if envoy is running within the pod
	// - false if envoy is not running within the pod
	EnvoyEnabled bool `json:"envoyEnabled,omitempty"`

	// The set of policies that apply to an endpoint may change within the reporting interval, this is the superset of all
	// policies that applied to the endpoint during that interval.
	AppliedPolicies []ResourceID `json:"appliedPolicies,omitempty"`

	// The list of services that exposed this endpoint at any moment during the reporting interval.
	//
	// Source: Determined from the Kubernetes endpoints resource associated with the service.
	Services []ResourceID `json:"services,omitempty"`

	// The ServiceAccount configured on this endpoint.
	ServiceAccount string `json:"serviceAccount,omitempty"`

	// The flow log aggregation name. This is used to locate flow logs associated with this endpoint when flow log
	// aggregation is turned on.
	FlowLogAggregationName string `json:"flowLogAggregationName,omitempty"`
}

type EndpointsReportFlow struct {
	// The source of the flow log.
	Source FlowEndpoint `json:"source"`

	// The destination of the flow log.
	Destination FlowEndpoint `json:"destination"`
}

type FlowEndpoint struct {
	// The endpoint type, indicating whether this is a Pod, HostEndpoint, NetworkSet, or internet.
	Kind string `json:"kind"`

	// The name of the endpoint. Note that this name may actually be a name prefix if flow logs have
	// been aggregated.
	Name string `json:"name"`

	// Whether the name is an aggregation prefix rather than the actual name.
	NameIsAggregationPrefix bool `json:"nameIsAggregationPrefix,omitempty"`

	// The namespace of the endpoint.
	Namespace string `json:"namespace,omitempty"`
}

type EndpointsReportNamespace struct {
	Namespace ResourceID `json:"namespace,omitempty"`

	// Whether ingress traffic was protected for all endpoints within this namespace within the reporting interval.
	// This is a summary of information contained in the endpoints data.
	IngressProtected bool `json:"ingressProtected,omitempty"`

	// Whether egress traffic was protected for all endpoints within this namespace within the reporting interval.
	// This is a summary of information contained in the endpoints data.
	EgressProtected bool `json:"egressProtected,omitempty"`

	// Whether ingress traffic was allowed from the internet for any endpoint within this namespace within the reporting
	// interval.
	IngressFromInternet bool `json:"ingressFromInternet,omitempty"`

	// Whether ingress traffic was allowed from the internet for any endpoint within this namespace within the reporting
	// interval.
	EgressToInternet bool `json:"egressToInternet,omitempty"`

	// Whether ingress traffic was allowed from another namespace for any endpoint within this namespace within the
	// reporting interval.
	IngressFromOtherNamespace bool `json:"ingressFromOtherNamespace,omitempty"`

	// Whether ingress traffic was allowed from another namespace for any endpoint within this namespace within the
	// reporting interval.
	EgressToOtherNamespace bool `json:"egressToOtherNamespace,omitempty"`

	// Whether envoy was enabled for all endpoints within this namespace within the reporting interval.
	// This is a summary of information contained in the endpoints data.
	EnvoyEnabled bool `json:"envoyEnabled,omitempty"`
}

type EndpointsReportService struct {
	Service ResourceID `json:"service,omitempty"`

	// Whether ingress traffic was protected for all endpoints within this namespace within the reporting interval.
	// This is a summary of information contained in the endpoints data.
	IngressProtected bool `json:"ingressProtected,omitempty"`

	// Whether ingress traffic was allowed from the internet for any endpoint exposed by this service within the reporting
	// interval.
	IngressFromInternet bool `json:"ingressFromInternet,omitempty"`

	// Whether ingress traffic was allowed from another namespace for any endpoint exposed by this service within the
	// reporting interval.
	IngressFromOtherNamespace bool `json:"ingressFromOtherNamespace,omitempty"`

	// Whether envoy was enabled for all endpoints that were exposed by this service within the reporting interval.
	// This is a summary of information contained in the endpoints data.
	EnvoyEnabled bool `json:"envoyEnabled,omitempty"`
}

// Prints FlowEndpoint contents. This is a slightly less verbose version of the resource names but should have
// sufficient context to be useful.
func (f FlowEndpoint) String() string {
	switch f.Kind {
	case KindK8sPod:
		// We add in the v1 version to be inline with the ResourceID printed format.
		return fmt.Sprintf("%s.v1(%s/%s)", f.Kind, f.Namespace, f.Name)
	case KindHostEndpoint, KindGlobalNetworkSet:
		return fmt.Sprintf("%s(%s)", f.Kind, f.Name)
	case KindFlowPublic, KindFlowPrivate:
		return f.Kind
	}
	return fmt.Sprintf("%s(%s/%s)", f.Kind, f.Namespace, f.Name)
}

// CISBenchmarkSummary describes a CIS benchmarking result across an entire cluster.
type CISBenchmarkSummary struct {
	Type      string `json:"type"`
	HighCount int    `json:"highCount"`
	MedCount  int    `json:"medCount"`
	LowCount  int    `json:"lowCount"`
}

// CISBenchmarkNode describes a CIS benchmarking result on a single node.
type CISBenchmarkNode struct {
	// NodeName is the name of the node the this set of benchmark results is from.
	NodeName string `json:"nodeName"`

	// KubernetesVersion is the version of the kubelet running on this node.
	KubernetesVersion string `json:"kubernetesVersion"`

	// BenchmarksVersion is the version of the benchmarks that ran on this node.
	BenchmarksVersion string `json:"benchmarksVersion"`

	// Summary is a set of summary stats for this set of node-specific benchmarks.
	Summary CISBenchmarkNodeSummary `json:"summary"`

	// Results is the detailed set of results for this set of node-specific benchmarks.
	Results []CISBenchmarkSectionResult `json:"results"`
}

// CISBenchmarkNodeSummary keeps count of tests passed, failed, and marked as info on a single node.
// Also has a status field to describe whether it is in HIGH, MED, or LOW status (based on [high|med]Threshold).
type CISBenchmarkNodeSummary struct {
	Status    string `json:"status"`
	TotalPass int    `json:"totalPass"`
	TotalFail int    `json:"totalFail"`
	TotalInfo int    `json:"totalInfo"`
	Total     int    `json:"total"`
}

// CISBenchmarkSectionResult describes the result of running the CIS benchmark on a single component.
type CISBenchmarkSectionResult struct {
	Status  string               `json:"status"`
	Section string               `json:"section"`
	Desc    string               `json:"desc"`
	Pass    int                  `json:"pass"`
	Fail    int                  `json:"fail"`
	Info    int                  `json:"info"`
	Results []CISBenchmarkResult `json:"results"`
}

// CISBenchmarkResult describes the result of a single CIS benchmark check.
type CISBenchmarkResult struct {
	TestNumber string `json:"testNumber"`
	TestDesc   string `json:"testDesc"`
	TestInfo   string `json:"testInfo"`
	Status     string `json:"status"`
	Scored     bool   `json:"scored"`
}

// CISBenchmarkResultCount keeps track of how many nodes had a certain test result.
type CISBenchmarkResultCount struct {
	CISBenchmarkResult
	Count int `json:"count"`
}
