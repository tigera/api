// Copyright (c) 2020-2021 Tigera, Inc. All rights reserved.

package v3

// The contents of this file create the model for performing authorization determination based on authorization header
// exchanges with the tigera-apiserver. No storage is required for achieving this and no libcalico client code will be
// created for the purpose of doing so.
// The tigera-apiserver will expose a create method just like k8s has for the TokenReviews api. A call to this endpoint
// will only reach the api-server if a valid authorization header is added to the request, otherwise the k8s api-server
// will respond directly with a 40x. If the request header is valid, the tigera-apiserver obtains the user information
// automatically from the k8s-apiserver and then performs an RBAC calculation based on the request data in the Spec.
// Since the response is based on the authorization header, the generated client may not very suitable for
// interacting with this api, depending on whether the client is handling requests for multiple users or not.

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindAuthorizationReview     = "AuthorizationReview"
	KindAuthorizationReviewList = "AuthorizationReviewList"
)

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuthorizationReviewList is a list of AuthorizationReview objects.
type AuthorizationReviewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AuthorizationReview `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AuthorizationReview struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AuthorizationReviewSpec   `json:"spec,omitempty"`
	Status AuthorizationReviewStatus `json:"status,omitempty"`
}

type AuthorizationReviewSpec struct {
	// The set of resource attributes that are being checked. Each resource attribute is expanded into individual
	// kind/resource and verbs.
	ResourceAttributes []AuthorizationReviewResourceAttributes `json:"resourceAttributes,omitempty" validate:"omitempty"`

	// User is the user you're testing for.
	// If you specify "User" but not "Groups", then is it interpreted as "What if User were not a member of any groups"
	// +optional
	User string `json:"user,omitempty" validate:"omitempty"`
	// Groups is the groups you're testing for.
	// +optional
	Groups []string `json:"groups,omitempty" validate:"omitempty"`
	// Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer
	// it needs a reflection here.
	// +optional
	Extra map[string][]string `json:"extra,omitempty" validate:"omitempty"`
	// UID information about the requesting user.
	// +optional
	UID string `json:"uid,omitempty" validate:"omitempty"`
}

type AuthorizationReviewResourceAttributes struct {
	// The API Group to check.
	APIGroup string `json:"apiGroup,omitempty" validate:"omitempty"`
	// The set of resources to check within the same API Group.
	Resources []string `json:"resources,omitempty" validate:"omitempty"`
	// The set of verbs to check. This is expanded for each resource and within the same API Group.
	Verbs []string `json:"verbs,omitempty" validate:"omitempty"`
}

type AuthorizationReviewStatus struct {
	// The set of authorized resource actions. A given API Group and resource combination will appear at most once in
	// this slice.
	AuthorizedResourceVerbs []AuthorizedResourceVerbs `json:"authorizedResourceVerbs,omitempty" validate:"omitempty"`
}

type AuthorizedResourceVerbs struct {
	// The API group.
	APIGroup string `json:"apiGroup,omitempty" validate:"omitempty"`
	// The resource.
	Resource string `json:"resource,omitempty" validate:"omitempty"`
	// The set of authorized actions for this resource. For a specific verb, this contains the set of resources for
	// which the user is authorized to perform that action. This is calculated to avoid duplication such that a single
	// resource instance can only be associated with a single entry in this slice. This allows a consumer of this API
	// to issue a minimal set of queries (e.g. watches) that cover, uniquely, the authorized set of resources.
	Verbs []AuthorizedResourceVerb `json:"verbs,omitempty" validate:"omitempty,dive"`
}

type AuthorizedResourceVerb struct {
	// The verb.
	Verb string `json:"verb"`
	// The group of resource instances that are authorized for this verb.
	ResourceGroups []AuthorizedResourceGroup `json:"resourceGroups"`
}

type AuthorizedResourceGroup struct {
	// The tier.  This is only valid for tiered policies, and tiers.
	Tier string `json:"tier,omitempty" validate:"omitempty"`

	// The namespace. If this is empty then the user is authorized cluster-wide (i.e. across all namespaces). This will
	// always be empty for cluster-scoped resources when the user is authorized.
	Namespace string `json:"namespace" validate:"omitempty"`

	// The UISettingsGroup name. This is only valid for uisettingsgroup/data sub resources.
	UISettingsGroup string `json:"uiSettingsGroup" validate:"omitempty"`

	// ManagedCluster is the name of the ManagedCluster. This is only valid for managedclusters.
	ManagedCluster string `json:"managedCluster" validate:"omitempty"`
}

// New AuthorizationReview creates a new (zeroed) AuthorizationReview struct with the TypeMetadata
// initialized to the current version.
func NewAuthorizationReview() *AuthorizationReview {
	return &AuthorizationReview{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindAuthorizationReview,
			APIVersion: GroupVersionCurrent,
		},
	}
}
