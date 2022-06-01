// Copyright (c) 2022 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package projectcalico

import (
	internalinterfaces "github.com/tigera/api/pkg/client/informers_generated/externalversions/internalinterfaces"
	v3 "github.com/tigera/api/pkg/client/informers_generated/externalversions/projectcalico/v3"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V3 provides access to shared informers for resources in V3.
	V3() v3.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V3 returns a new v3.Interface.
func (g *group) V3() v3.Interface {
	return v3.New(g.factory, g.namespace, g.tweakListOptions)
}
