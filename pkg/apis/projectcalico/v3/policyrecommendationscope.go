// Copyright (c) 2022 Tigera, Inc. All rights reserved.
package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindPolicyRecommendationScope     = "PolicyRecommendationScope"
	KindPolicyRecommendationScopeList = "PolicyRecommendationScopeList"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:subresource:status
type PolicyRecommendationScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   PolicyRecommendationScopeSpec   `json:"spec,omitempty"`
	Status PolicyRecommendationScopeStatus `json:"status,omitempty"`
}

type PolicyRecommendationScopeSpec struct {
	// How frequently to run the recommendation engine to create and refine recommended policies.
	// [Default: 150s]
	// +optional
	Interval *metav1.Duration `json:"interval,omitempty"`

	// How far back to look in flow logs when first creating a recommended policy.
	// [Default: 24h]
	// +optional
	InitialLookback *metav1.Duration `json:"initialLookback,omitempty"`

	// StabilizationPeriod is the amount of time a recommended policy should remain unchanged to be
	// deemed stable and ready to be enforced.
	// [Default: 10m]
	// +optional
	StabilizationPeriod *metav1.Duration `json:"stabilizationPeriod,omitempty"`

	// The maximum number of rules that are permitted in the ingress or egress set. For egress rules,
	// any egress domain rules will be simplified by contracting all domains into a single egress
	// domain NetworkSet. If the number of rules exceeds this limit, the recommendation engine will
	// treat this as an error condition.
	// [Default: 20]
	// +optional
	MaxRules *int `json:"maxRules,omitempty"`

	// The number of staged policies that are actively learning at any one time, after which the
	// policy recommendation engine will stop adding new recommendations.
	// [Default: 20]
	// +optional
	PoliciesLearningCutOff *int `json:"policiesLearningCutOff,omitempty"`

	// The namespace spec contains the namespace relative recommendation vars.
	NamespaceSpec PolicyRecommendationScopeNamespaceSpec `json:"namespaceSpec,omitempty"`
}

type PolicyRecommendationScopeStatus struct {
	Conditions []PolicyRecommendationScopeStatusCondition `json:"conditions,omitempty"`
}

type PolicyRecommendationScopeStatusType string
type PolicyRecommendationScopeStatusValue string

// Condition contains various status information
type PolicyRecommendationScopeStatusCondition struct {
	Message string                               `json:"message,omitempty"`
	Reason  string                               `json:"reason,omitempty"`
	Status  PolicyRecommendationScopeStatusValue `json:"status"`
	Type    PolicyRecommendationScopeStatusType  `json:"type"`
}

// PolicyRecommendationScopeNamespaceSpec contains namespace information that defines the namespace based
// recommended policy.
type PolicyRecommendationScopeNamespaceSpec struct {
	// Pass intra-namespace traffic.
	// [Default: false]
	// +optional
	IntraNamespacePassThroughTraffic bool `json:"intraNamespacePassThroughTraffic,omitempty"`
	// Recommendation status. One of Enabled, Disabled.
	RecStatus PolicyRecommendationNamespaceStatus `json:"recStatus,omitempty" validate:"omitempty,policyrecstatus"`
	// The namespace selector is an expression used to pick out the namespaces that the policy
	// recommendation engine should create policies for. The syntax is the same as the
	// NetworkPolicy.projectcalico.org resource selectors.
	Selector string `json:"selector" validate:"selector"`
	// The name of the policy recommendation tier for namespace-isolated policies.
	// [Default: "namespace-isolation"]
	// +optional
	TierName string `json:"tierName,omitempty" validate:"omitempty,name"`
}

type PolicyRecommendationNamespaceStatus string

const (
	PolicyRecommendationScopeEnabled  PolicyRecommendationNamespaceStatus = "Enabled"
	PolicyRecommendationScopeDisabled PolicyRecommendationNamespaceStatus = "Disabled"
)

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PolicyRecommendationList contains a list of Monitor
type PolicyRecommendationScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyRecommendationScope `json:"items"`
}

// NewPolicyRecommendationScope creates a new (zeroed) PolicyRecommendationScope struct.
// TypeMetadata initialized to the current version.
func NewPolicyRecommendationScope() *PolicyRecommendationScope {
	return &PolicyRecommendationScope{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindPolicyRecommendationScope,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewPolicyRecommendationScopeList creates a new (zeroed) PolicyRecommendationScopeList struct with the
// TypeMetadata initialized to the current version.
func NewPolicyRecommendationScopeList() *PolicyRecommendationScopeList {
	return &PolicyRecommendationScopeList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindPolicyRecommendationScopeList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
