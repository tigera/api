// Copyright (c) 2023 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BGPFilterLister helps list BGPFilters.
// All objects returned here must be treated as read-only.
type BGPFilterLister interface {
	// List lists all BGPFilters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.BGPFilter, err error)
	// Get retrieves the BGPFilter from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.BGPFilter, error)
	BGPFilterListerExpansion
}

// bGPFilterLister implements the BGPFilterLister interface.
type bGPFilterLister struct {
	indexer cache.Indexer
}

// NewBGPFilterLister returns a new BGPFilterLister.
func NewBGPFilterLister(indexer cache.Indexer) BGPFilterLister {
	return &bGPFilterLister{indexer: indexer}
}

// List lists all BGPFilters in the indexer.
func (s *bGPFilterLister) List(selector labels.Selector) (ret []*v3.BGPFilter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.BGPFilter))
	})
	return ret, err
}

// Get retrieves the BGPFilter from the index for a given name.
func (s *bGPFilterLister) Get(name string) (*v3.BGPFilter, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("bgpfilter"), name)
	}
	return obj.(*v3.BGPFilter), nil
}
