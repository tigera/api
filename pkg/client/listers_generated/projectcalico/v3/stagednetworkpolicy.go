// Copyright (c) 2023 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StagedNetworkPolicyLister helps list StagedNetworkPolicies.
// All objects returned here must be treated as read-only.
type StagedNetworkPolicyLister interface {
	// List lists all StagedNetworkPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.StagedNetworkPolicy, err error)
	// StagedNetworkPolicies returns an object that can list and get StagedNetworkPolicies.
	StagedNetworkPolicies(namespace string) StagedNetworkPolicyNamespaceLister
	StagedNetworkPolicyListerExpansion
}

// stagedNetworkPolicyLister implements the StagedNetworkPolicyLister interface.
type stagedNetworkPolicyLister struct {
	indexer cache.Indexer
}

// NewStagedNetworkPolicyLister returns a new StagedNetworkPolicyLister.
func NewStagedNetworkPolicyLister(indexer cache.Indexer) StagedNetworkPolicyLister {
	return &stagedNetworkPolicyLister{indexer: indexer}
}

// List lists all StagedNetworkPolicies in the indexer.
func (s *stagedNetworkPolicyLister) List(selector labels.Selector) (ret []*v3.StagedNetworkPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.StagedNetworkPolicy))
	})
	return ret, err
}

// StagedNetworkPolicies returns an object that can list and get StagedNetworkPolicies.
func (s *stagedNetworkPolicyLister) StagedNetworkPolicies(namespace string) StagedNetworkPolicyNamespaceLister {
	return stagedNetworkPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StagedNetworkPolicyNamespaceLister helps list and get StagedNetworkPolicies.
// All objects returned here must be treated as read-only.
type StagedNetworkPolicyNamespaceLister interface {
	// List lists all StagedNetworkPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.StagedNetworkPolicy, err error)
	// Get retrieves the StagedNetworkPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.StagedNetworkPolicy, error)
	StagedNetworkPolicyNamespaceListerExpansion
}

// stagedNetworkPolicyNamespaceLister implements the StagedNetworkPolicyNamespaceLister
// interface.
type stagedNetworkPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StagedNetworkPolicies in the indexer for a given namespace.
func (s stagedNetworkPolicyNamespaceLister) List(selector labels.Selector) (ret []*v3.StagedNetworkPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.StagedNetworkPolicy))
	})
	return ret, err
}

// Get retrieves the StagedNetworkPolicy from the indexer for a given namespace and name.
func (s stagedNetworkPolicyNamespaceLister) Get(name string) (*v3.StagedNetworkPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("stagednetworkpolicy"), name)
	}
	return obj.(*v3.StagedNetworkPolicy), nil
}
