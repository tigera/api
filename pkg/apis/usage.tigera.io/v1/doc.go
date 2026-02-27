// Copyright (c) 2026 Tigera, Inc. All rights reserved.

// +k8s:deepcopy-gen=package,register
// +groupName=usage.tigera.io
package v1

import "k8s.io/apimachinery/pkg/runtime/schema"

var UsageGroupVersion = schema.GroupVersion{Group: "usage.tigera.io", Version: "v1"}
