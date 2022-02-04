// Copyright (c) 2019,2021 Tigera, Inc. All rights reserved.

package v3

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindGlobalAlert     = "GlobalAlert"
	KindGlobalAlertList = "GlobalAlertList"

	GlobalAlertDataSetAudit = "audit"
	GlobalAlertDataSetDNS   = "dns"
	GlobalAlertDataSetFlows = "flows"
	GlobalAlertDataSetL7    = "l7"
	GlobalAlertDataSetWAF   = "waf"

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
	Summary       string                    `json:"summary,omitempty" validate:"omitempty"`
	Description   string                    `json:"description" validate:"required"`
	Severity      int                       `json:"severity" validate:"required,min=1,max=100"`
	Period        *metav1.Duration          `json:"period,omitempty" validate:"omitempty"`
	Lookback      *metav1.Duration          `json:"lookback,omitempty" validate:"omitempty"`
	DataSet       string                    `json:"dataSet" validate:"required,oneof=flows dns audit l7 waf"`
	Query         string                    `json:"query,omitempty" validate:"omitempty"`
	AggregateBy   []string                  `json:"aggregateBy,omitempty" validate:"omitempty"`
	Field         string                    `json:"field,omitempty" validate:"omitempty"`
	Metric        string                    `json:"metric,omitempty" validate:"omitempty,oneof=avg max min sum count"`
	Condition     string                    `json:"condition,omitempty" validate:"omitempty,oneof=eq not_eq gt gte lt lte"`
	Threshold     float64                   `json:"threshold,omitempty" validate:"omitempty"`
	Substitutions []GlobalAlertSubstitution `json:"substitutions,omitempty" validate:"omitempty"`
}

type GlobalAlertStatus struct {
	LastUpdate      *metav1.Time     `json:"lastUpdate,omitempty"`
	Active          bool             `json:"active"`
	Healthy         bool             `json:"healthy"`
	LastExecuted    *metav1.Time     `json:"lastExecuted,omitempty"`
	LastEvent       *metav1.Time     `json:"lastEvent,omitempty"`
	ErrorConditions []ErrorCondition `json:"errorConditions,omitempty"`
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
