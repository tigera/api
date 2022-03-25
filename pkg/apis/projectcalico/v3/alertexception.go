// Copyright (c) 2022 Tigera, Inc. All rights reserved.

package v3

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindAlertException     = "AlertException"
	KindAlertExceptionList = "AlertExceptionList"

	AlertExceptionMinPeriod = time.Minute
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

	// Period controls how long an alert exception will be active. It is optional and
	// omitting Period will make the alert exception active forever.
	// +optional
	Period *metav1.Duration `json:"period,omitempty" validate:"omitempty"`
}

// AlertExceptionStatus contains the status of an alert exception.
type AlertExceptionStatus struct {
	LastExecuted *metav1.Time `json:"lastExecuted,omitempty"`
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
