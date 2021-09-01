// Copyright (c) 2021 Tigera, Inc. All rights reserved.

package v3

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	KindUISettingsGroup     = "UISettingsGroup"
	KindUISettingsGroupList = "UISettingsGroupList"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UISettingsGroup contains the settings that dictate how many UI settings may be created for a
// specific cluster/user combination. UI settings may only be persisted if there is a
// corresponding UISettingsGroup resource.
type UISettingsGroup struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the UISettingsGroup.
	Spec UISettingsGroupSpec `json:"spec,omitempty"`
}

// UISettingsGroupSpec contains the specification for a UISettingsGroup resource.
type UISettingsGroupSpec struct {
	// This description is displayed by the UI when asking where to store any UI-specific settings
	// such as views, layers, dashboards etc. This name should be a short description that relates
	// the settings to the set of clusters defined below, the set of users or groups that are able to
	// access to these settings (defined via RBAC) or the set of applications common to the set of
	// users or groups that can access these settings.
	// Examples might be:
	// - "cluster" when these settings apply to the whole cluster
	// - "global" when these settings apply to all clusters (in an Multi-Cluster environment)
	// - "security team" if these settings are accessible only to the security group and therefore
	//   applicable to the applications accessible by that team
	// - "storefront" if these settings are accessible to all users and groups that can access the
	//   storefront set of applications
	// - "user" if these settings are accessible to only a single user
	Description string `json:"description"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UISettingsGroupList contains a list of UISettingsGroup resources.
type UISettingsGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []UISettingsGroup `json:"items"`
}

// NewUISettingsGroup creates a new (zeroed) UISettingsGroup struct with the TypeMetadata
// initialised to the current version.
func NewUISettingsGroup() *UISettingsGroup {
	return &UISettingsGroup{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindUISettingsGroup,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewUISettingsGroupList creates a new (zeroed) UISettingsGroupList struct with the
// TypeMetadata initialised to the current version.
func NewUISettingsGroupList() *UISettingsGroupList {
	return &UISettingsGroupList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindUISettingsGroupList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
