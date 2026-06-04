// Copyright (c) 2026 Tigera, Inc. All rights reserved.

// +k8s:deepcopy-gen=package,register
// +k8s:openapi-gen=true
// +k8s:openapi-model-package=com.github.tigera.api.pkg.apis.applicationlayer.projectcalico.v3

// Package v3 is the v3 version of the applicationlayer.projectcalico.org API.
//
// The group name `applicationlayer.projectcalico.org` (rather than the
// existing `projectcalico.org/v3`) is a deliberate choice — see
// tigera/designs#25 for the full rationale. Short version:
// `projectcalico.org/v3` is served by the aggregated apiserver and adding a
// type there requires editing its hardcoded `AllKnownTypes` list, whereas
// `applicationlayer.projectcalico.org` isn't registered as an API group
// anywhere in calico-private today (only as a label/annotation prefix on
// existing resources). That makes it a vanilla CRD group with no
// apiserver/scheme work, which is the apiserver-independent direction the
// codebase is moving in.
//
// +groupName=applicationlayer.projectcalico.org

package v3 // import "github.com/tigera/api/pkg/apis/applicationlayer.projectcalico.org/v3"
