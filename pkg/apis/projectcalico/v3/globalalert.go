// Copyright (c) 2019,2021 Tigera, Inc. All rights reserved.

package v3

import (
	"encoding/json"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindGlobalAlert     = "GlobalAlert"
	KindGlobalAlertList = "GlobalAlertList"

	GlobalAlertDataSetAudit         = "audit"
	GlobalAlertDataSetDNS           = "dns"
	GlobalAlertDataSetFlows         = "flows"
	GlobalAlertDataSetL7            = "l7"
	GlobalAlertDataSetWAF           = "waf"
	GlobalAlertDataSetVulnerability = "vulnerability"

	GlobalAlertMetricAvg   = "avg"
	GlobalAlertMetricMax   = "max"
	GlobalAlertMetrixMin   = "min"
	GlobalAlertMetricSum   = "sum"
	GlobalAlertMetricCount = "count"

	GlobalAlertMinPeriod   = time.Minute
	GlobalAlertMinLookback = GlobalAlertMinPeriod
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GlobalAlert struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the GlobalAlert.
	Spec   GlobalAlertSpec   `json:"spec,omitempty"`
	Status GlobalAlertStatus `json:"status,omitempty"`
}

type GlobalAlertSpec struct {
	// Type will dictate how the fields of the GlobalAlert will be utilized.
	// Each Type will have different usages and defaults for the fields. [Default: RuleBased]
	Type GlobalAlertType `json:"type,omitempty" validate:"omitempty,globalAlertType"`
	// Template for the description field in generated events, description is used if this is omitted.
	Summary string `json:"summary,omitempty" validate:"omitempty"`
	// Human-readable description of the template.
	Description string `json:"description" validate:"required"`
	// Severity of the alert for display in Manager.
	Severity int `json:"severity" validate:"required,min=1,max=100"`
	// If Type is RuleBased, it is how often the query defined will run.
	// If Type is AnomalyDetection it is how often the detector will be run.
	Period *metav1.Duration `json:"period,omitempty" validate:"omitempty"`
	// How much data to gather at once.
	// If Type is RuleBased, it must exceed audit log flush interval, dnsLogsFlushInterval, or flowLogsFlushInterval as appropriate.
	Lookback *metav1.Duration `json:"lookback,omitempty" validate:"omitempty"`
	// DataSet determines which dataset type the Query will use.
	// Required and used only if Type is RuleBased.
	DataSet string `json:"dataSet,omitempty" validate:"omitempty,oneof=flows dns audit l7 waf vulnerability"`
	// Which data to include from the source data set. Written in a domain-specific query language. Only used if Type is RuleBased.
	Query string `json:"query,omitempty" validate:"omitempty"`
	// An optional list of fields to aggregate results.
	// Only used if Type is RuleBased.
	AggregateBy []string `json:"aggregateBy,omitempty" validate:"omitempty"`
	// Which field to aggregate results by if using a metric other than count.
	// Only used if Type is RuleBased.
	Field string `json:"field,omitempty" validate:"omitempty"`
	// A metric to apply to aggregated results. count is the number of log entries matching the aggregation pattern.
	// Others are applied only to numeric fields in the logs.
	// Only used if Type is RuleBased.
	Metric string `json:"metric,omitempty" validate:"omitempty,oneof=avg max min sum count"`
	// Compare the value of the metric to the threshold using this condition.
	// Only used if Type is RuleBased.
	Condition string `json:"condition,omitempty" validate:"omitempty,oneof=eq not_eq gt gte lt lte"`
	// A numeric value to compare the value of the metric against.
	// Only used if Type is RuleBased.
	Threshold float64 `json:"threshold,omitempty" validate:"omitempty"`
	// An optional list of values to replace variable names in query.
	// Only used if Type is RuleBased.
	Substitutions []GlobalAlertSubstitution `json:"substitutions,omitempty" validate:"omitempty"`
	// Parameters for configuring an AnomalyDetection run.
	// Only used if Type is AnomalyDetection.
	Detector *DetectorParams `json:"detector,omitempty" validate:"omitempty"`
}

type GlobalAlertType string

const (
	GlobalAlertTypeRuleBased        GlobalAlertType = "RuleBased"
	GlobalAlertTypeAnomalyDetection GlobalAlertType = "AnomalyDetection"
)

func (t *GlobalAlertType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		*t = GlobalAlertTypeRuleBased
	} else {
		*t = GlobalAlertType(s)
	}
	return nil
}

type GlobalAlertStatus struct {
	LastUpdate      *metav1.Time     `json:"lastUpdate,omitempty"`
	Active          bool             `json:"active"`
	Healthy         bool             `json:"healthy"`
	LastExecuted    *metav1.Time     `json:"lastExecuted,omitempty"`
	LastEvent       *metav1.Time     `json:"lastEvent,omitempty"`
	ErrorConditions []ErrorCondition `json:"errorConditions,omitempty"`
}

type DetectorParams struct {
	// Name specifies the AnomalyDetection Detector to run.
	Name string `json:"name"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalAlertList contains a list of GlobalAlert resources.
type GlobalAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GlobalAlert `json:"items"`
}

// GlobalAlertSubstitution substitutes for the variables in the set operators of a Query.
type GlobalAlertSubstitution struct {
	Name   string   `json:"name" validate:"required"`
	Values []string `json:"values,omitempty"`
}

// NewGlobalAlert creates a new (zeroed) GlobalAlert struct with the TypeMetadata
// initialized to the current version.
func NewGlobalAlert() *GlobalAlert {
	return &GlobalAlert{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindGlobalAlert,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewGlobalAlertList creates a new (zeroed) GlobalAlertList struct with the TypeMetadata
// initialized to the current version.
func NewGlobalAlertList() *GlobalAlertList {
	return &GlobalAlertList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindGlobalAlertList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
