// Copyright (c) 2019,2021 Tigera, Inc. All rights reserved.

package v3

// The contents of this file create the model for how to do authorization header exchanges with the tigera-apiserver for
// the purpose of authentication and obtaining user info. However, no storage is required for achieving this and no
// libcalico client code will be created for the purpose of doing so.
// The tigera-apiserver will expose a create method just like k8s has for the TokenReviews api. A call to this endpoint
// will only reach the api-server if a valid authorization header is added to the request, otherwise the k8s api-server
// will respond directly with a 40x. If the request header is valid, the tigera-apiserver obtains the user information
// automatically from the k8s-apiserver and then return it in the AuthenticationReviewStatus.
// Since the response is entirely based on the authorization header, the generated client is not very suitable for
// interacting with this api. It would mean a new client config has to be created for each incoming request. By creating
// a separate client dedicated to authn only, simpler and easier-to-maintain code can be created.

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindAuthenticationReview     = "AuthenticationReview"
	KindAuthenticationReviewList = "AuthenticationReviewList"
)

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuthenticationReviewList is a list of AuthenticationReview objects.
type AuthenticationReviewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []AuthenticationReview `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AuthenticationReview struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Status AuthenticationReviewStatus `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"`
}

type AuthenticationReviewStatus struct {
	Name   string              `json:"name,omitempty" validate:"omitempty"`
	UID    string              `json:"uid,omitempty" validate:"omitempty"`
	Groups []string            `json:"groups,omitempty" validate:"omitempty"`
	Extra  map[string][]string `json:"extra,omitempty" validate:"omitempty"`
}

// New AuthenticationReview creates a new (zeroed) AuthenticationReview struct with the TypeMetadata
// initialized to the current version.
func NewAuthenticationReview() *AuthenticationReview {
	return &AuthenticationReview{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindAuthenticationReview,
			APIVersion: GroupVersionCurrent,
		},
	}
}
