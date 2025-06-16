// Copyright (c) 2023 Tigera, Inc. All rights reserved.

package v3

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SecurityEventWebhookConsumer string
type SecurityEventWebhookState string

const (
	KindSecurityEventWebhook     = "SecurityEventWebhook"
	KindSecurityEventWebhookList = "SecurityEventWebhookList"

	SecurityEventWebhookConsumerSlack        SecurityEventWebhookConsumer = "Slack"
	SecurityEventWebhookConsumerJira         SecurityEventWebhookConsumer = "Jira"
	SecurityEventWebhookConsumerGeneric      SecurityEventWebhookConsumer = "Generic"
	SecurityEventWebhookConsumerAlertManager SecurityEventWebhookConsumer = "AlertManager"

	SecurityEventWebhookStateEnabled  SecurityEventWebhookState = "Enabled"
	SecurityEventWebhookStateDisabled SecurityEventWebhookState = "Disabled"
	SecurityEventWebhookStateDebug    SecurityEventWebhookState = "Debug"
	SecurityEventWebhookStateTest     SecurityEventWebhookState = "Test"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SecurityEventWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Status            []metav1.Condition       `json:"status,omitempty" validate:"omitempty"`
	Spec              SecurityEventWebhookSpec `json:"spec" validate:"required"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SecurityEventWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []SecurityEventWebhook `json:"items"`
}

type SecurityEventWebhookSpec struct {
	// indicates the SecurityEventWebhook intended consumer, one of: Slack, Jira, Generic, AlertManager
	Consumer SecurityEventWebhookConsumer `json:"consumer" validate:"required,oneof=Slack Jira Generic AlertManager"`
	// defines the webhook desired state, one of: Enabled, Disabled, Test or Debug
	State SecurityEventWebhookState `json:"state" validate:"required,oneof=Enabled Disabled Test Debug"`
	// defines the SecurityEventWebhook query to be executed against fields of SecurityEvents
	Query string `json:"query" validate:"required"`
	// contains the SecurityEventWebhook's configuration associated with the intended Consumer
	Config []SecurityEventWebhookConfigVar `json:"config" validate:"required"`
}

type SecurityEventWebhookConfigVar struct {
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// +optional
	Value string `json:"value,omitempty"`
	// +optional
	ValueFrom *SecurityEventWebhookConfigVarSource `json:"valueFrom,omitempty"`
}

type SecurityEventWebhookConfigVarSource struct {
	// +optional
	ConfigMapKeyRef *corev1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty" protobuf:"bytes,3,opt,name=configMapKeyRef"`
	// +optional
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty" protobuf:"bytes,4,opt,name=secretKeyRef"`
}

// NewSecurityEventWebhook creates a new SecurityEventWebhook struct
// with the TypeMetadata initialized to the current API version.
func NewSecurityEventWebhook() *SecurityEventWebhook {
	return &SecurityEventWebhook{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindSecurityEventWebhook,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewSecurityEventWebhookList creates a new SecurityEventWebhookList struct
// with the TypeMetadata initialized to the current API version.
func NewSecurityEventWebhookList() *SecurityEventWebhookList {
	return &SecurityEventWebhookList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindSecurityEventWebhookList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
