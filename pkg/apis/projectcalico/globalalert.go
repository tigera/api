// Copyright (c) 2020 Tigera, Inc. All rights reserved.

package projectcalico

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
	Description string           `json:"description" validate:"required"`
	Severity    int              `json:"severity" validate:"required,min=1,max=100"`
	Period      *metav1.Duration `json:"period,omitempty" validate:"omitempty"`
	Lookback    *metav1.Duration `json:"lookback,omitempty" validate:"omitempty"`
	DataSet     string           `json:"dataSet" validate:"required,oneof=flows dns audit"`
	Query       string           `json:"query,omitempty" validate:"omitempty"`
	AggregateBy []string         `json:"aggregateBy,omitempty" validate:"omitempty"`
	Field       string           `json:"field,omitempty" validate:"omitempty"`
	Metric      string           `json:"metric,omitempty" validate:"omitempty,oneof=avg max min sum count"`
	Condition   string           `json:"condition,omitempty" validate:"omitempty,oneof=eq not_eq gt gte lt lte"`
	Threshold   float64          `json:"threshold,omitempty" validate:"omitempty"`
}

type GlobalAlertStatus struct {
	LastUpdate      *metav1.Time     `json:"lastUpdate,omitempty"`
	Active          bool             `json:"active"`
	Healthy         bool             `json:"healthy"`
	ExecutionState  string           `json:"executionState,omitempty"`
	LastExecuted    *metav1.Time     `json:"lastExecuted,omitempty"`
	LastEvent       *metav1.Time     `json:"lastEvent,omitempty"`
	ErrorConditions []ErrorCondition `json:"errorConditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalAlertList contains a list of GlobalAlert resources.
type GlobalAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GlobalAlert `json:"items"`
}

type ErrorCondition struct {
	Type    string `json:"type" validate:"required"`
	Message string `json:"message" validate:"required"`
}