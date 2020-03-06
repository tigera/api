// Copyright (c) 2020 Tigera, Inc. All rights reserved.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GlobalAlertTemplate struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the GlobalAlert.
	Spec GlobalAlertTemplateSpec `json:"spec,omitempty"`
}

type GlobalAlertTemplateSpec struct {
	Summary     string           `json:"summary,omitempty" validate:"omitempty"`
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

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalAlertList contains a list of GlobalAlert resources.
type GlobalAlertTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GlobalAlertTemplate `json:"items"`
}