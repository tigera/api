// Copyright (c) 2023 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EgressGatewayPolicyLister helps list EgressGatewayPolicies.
// All objects returned here must be treated as read-only.
type EgressGatewayPolicyLister interface {
	// List lists all EgressGatewayPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.EgressGatewayPolicy, err error)
	// Get retrieves the EgressGatewayPolicy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.EgressGatewayPolicy, error)
	EgressGatewayPolicyListerExpansion
}

// egressGatewayPolicyLister implements the EgressGatewayPolicyLister interface.
type egressGatewayPolicyLister struct {
	indexer cache.Indexer
}

// NewEgressGatewayPolicyLister returns a new EgressGatewayPolicyLister.
func NewEgressGatewayPolicyLister(indexer cache.Indexer) EgressGatewayPolicyLister {
	return &egressGatewayPolicyLister{indexer: indexer}
}

// List lists all EgressGatewayPolicies in the indexer.
func (s *egressGatewayPolicyLister) List(selector labels.Selector) (ret []*v3.EgressGatewayPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.EgressGatewayPolicy))
	})
	return ret, err
}

// Get retrieves the EgressGatewayPolicy from the index for a given name.
func (s *egressGatewayPolicyLister) Get(name string) (*v3.EgressGatewayPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("egressgatewaypolicy"), name)
	}
	return obj.(*v3.EgressGatewayPolicy), nil
}