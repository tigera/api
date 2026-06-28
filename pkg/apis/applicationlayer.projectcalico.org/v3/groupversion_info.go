// Copyright (c) 2026 Tigera, Inc. All rights reserved.

// Package v3 contains API Schema definitions for the applicationlayer.projectcalico.org v3 API group
// +kubebuilder:object:generate=true
// +groupName=applicationlayer.projectcalico.org
package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupName is the group name used in this package.
const GroupName = "applicationlayer.projectcalico.org"

// GroupVersion is group version used to register these objects.
var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v3"}

// SchemeGroupVersion is an alias kept for callers that follow the
// projectcalico.org/v3 register.go naming convention.
var SchemeGroupVersion = GroupVersion

var (
	// SchemeBuilder collects the scheme registration functions.
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder

	// AllKnownTypes lists every type registered with the scheme.
	AllKnownTypes = []runtime.Object{
		&GlobalWAFPolicy{}, &GlobalWAFPolicyList{},
		&WAFPolicy{}, &WAFPolicyList{},
		&WAFPlugin{}, &WAFPluginList{},
		&GlobalWAFPlugin{}, &GlobalWAFPluginList{},
		&GlobalWAFValidationPolicy{}, &GlobalWAFValidationPolicyList{},
		&WAFValidationPolicy{}, &WAFValidationPolicyList{},
	}
)

func init() {
	localSchemeBuilder.Register(addKnownTypes)
}

// AddToScheme adds the types in this group-version to the given scheme.
func AddToScheme(scheme *runtime.Scheme) error {
	return localSchemeBuilder.AddToScheme(scheme)
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion, AllKnownTypes...)
	metav1.AddToGroupVersion(scheme, GroupVersion)
	return nil
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return GroupVersion.WithResource(resource).GroupResource()
}
