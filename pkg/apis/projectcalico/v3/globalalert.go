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

	GlobalAlertDataSetAudit = "audit"
	GlobalAlertDataSetDNS   = "dns"
	GlobalAlertDataSetFlows = "flows"
	GlobalAlertDataSetL7    = "l7"

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
	// Type will dictate how the fields of the GlobalAlert will be utilized. Each Type
	// will have different usages and defaults for the fields. [Default: UserDefined]
	Type        GlobalAlertType `json:"type,omitempty" validate:"omitempty"`
	Summary     string          `json:"summary,omitempty" validate:"omitempty"`
	Description string          `json:"description" validate:"required"`
	// Job sepcifies the AnomalyDetectionJob to run if GlobalAlert is of Type AnomalyDetection.
	// Required if Type is of AnomalyDetection.
	Job      string           `json:"job,omitempty" validate:"omitempty"`
	Severity int              `json:"severity" validate:"required,min=1,max=100"`
	Period   *metav1.Duration `json:"period,omitempty" validate:"omitempty"`
	Lookback *metav1.Duration `json:"lookback,omitempty" validate:"omitempty"`
	// DataSet detemines which dataset type the Query will use.  Required if Type is ofUserDefined.
	DataSet     string   `json:"dataSet" validate:"omitempty,oneof=flows dns audit"`
	Query       string   `json:"query,omitempty" validate:"omitempty"`
	AggregateBy []string `json:"aggregateBy,omitempty" validate:"omitempty"`
	Field       string   `json:"field,omitempty" validate:"omitempty"`
	Metric      string   `json:"metric,omitempty" validate:"omitempty,oneof=avg max min sum count"`
	Condition   string   `json:"condition,omitempty" validate:"omitempty,oneof=eq not_eq gt gte lt lte"`
	Threshold   float64  `json:"threshold,omitempty" validate:"omitempty"`
}

type GlobalAlertType string

const (
	GlobalAlertTypeUserDefined      GlobalAlertType = "UserDefined"
	GlobalAlertTypeAnomalyDetection GlobalAlertType = "AnomalyDetection"
)

func (t *Type) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		*t = Type("UserDefined")
	} else {
		*t = Type(s)
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

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalAlertList contains a list of GlobalAlert resources.
type GlobalAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GlobalAlert `json:"items"`
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
