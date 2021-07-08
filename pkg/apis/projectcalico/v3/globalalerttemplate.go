// Copyright (c) 2019,2021 Tigera, Inc. All rights reserved.

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

const (
	KindGlobalAlertTemplate     = "GlobalAlertTemplate"
	KindGlobalAlertTemplateList = "GlobalAlertTemplateList"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GlobalAlertTemplate struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the GlobalAlert.
	Spec GlobalAlertSpec `json:"spec,omitempty"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalAlertList contains a list of GlobalAlert resources.
type GlobalAlertTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GlobalAlertTemplate `json:"items"`
}

// NewGlobalAlert creates a new (zeroed) GlobalAlert struct with the TypeMetadata
// initialized to the current version.
func NewGlobalAlertTemplate() *GlobalAlertTemplate {
	return &GlobalAlertTemplate{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindGlobalAlertTemplate,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewGlobalAlertTemplateList creates a new (zeroed) GlobalAlertTemplateList struct with the TypeMetadata
// initialized to the current version.
func NewGlobalAlertTemplateList() *GlobalAlertTemplateList {
	return &GlobalAlertTemplateList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindGlobalAlertTemplateList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
