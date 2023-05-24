// Copyright (c) 2022 Tigera, Inc. All rights reserved.

package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindAlertException     = "AlertException"
	KindAlertExceptionList = "AlertExceptionList"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlertException defines exceptions for alert events.
type AlertException struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlertExceptionSpec   `json:"spec,omitempty"`
	Status AlertExceptionStatus `json:"status,omitempty"`
}

// AlertExceptionSpec contains the specification for an alert exception resource.
type AlertExceptionSpec struct {
	// The description is displayed by the UI.
	Description string `json:"description" validate:"required"`

	// Selector defines a query string for alert events to be excluded from UI search results.
	Selector string `json:"selector" validate:"required"`

	// StartTime defines the start time from which this alert exception will take effect.
	// If the value is in the past, matched alerts will be filtered immediately.
	// If the value is changed to a future time, alert exceptions will restart at that time.
	// +kubebuilder:validation:Format="date-time"
	StartTime metav1.Time `json:"startTime" validate:"required"`

	// EndTime defines the end time at which this alert exception will expire.
	// If omitted the alert exception filtering will continue indefinitely.
	// +optional
	//+kubebuilder:validation:Format="date-time"
	EndTime *metav1.Time `json:"endTime,omitempty" validate:"omitempty"`
}

// AlertExceptionStatus contains the status of an alert exception.
type AlertExceptionStatus struct {
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlertExceptionList contains a list of AlertException resources.
type AlertExceptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AlertException `json:"items"`
}

func NewAlertException() *AlertException {
	return &AlertException{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindAlertException,
			APIVersion: GroupVersionCurrent,
		},
	}
}

func NewAlertExceptionList() *AlertExceptionList {
	return &AlertExceptionList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindAlertExceptionList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
