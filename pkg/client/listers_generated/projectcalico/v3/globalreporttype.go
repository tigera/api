// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GlobalReportTypeLister helps list GlobalReportTypes.
// All objects returned here must be treated as read-only.
type GlobalReportTypeLister interface {
	// List lists all GlobalReportTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.GlobalReportType, err error)
	// GlobalReportTypes returns an object that can list and get GlobalReportTypes.
	GlobalReportTypes(namespace string) GlobalReportTypeNamespaceLister
	GlobalReportTypeListerExpansion
}

// globalReportTypeLister implements the GlobalReportTypeLister interface.
type globalReportTypeLister struct {
	indexer cache.Indexer
}

// NewGlobalReportTypeLister returns a new GlobalReportTypeLister.
func NewGlobalReportTypeLister(indexer cache.Indexer) GlobalReportTypeLister {
	return &globalReportTypeLister{indexer: indexer}
}

// List lists all GlobalReportTypes in the indexer.
func (s *globalReportTypeLister) List(selector labels.Selector) (ret []*v3.GlobalReportType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.GlobalReportType))
	})
	return ret, err
}

// GlobalReportTypes returns an object that can list and get GlobalReportTypes.
func (s *globalReportTypeLister) GlobalReportTypes(namespace string) GlobalReportTypeNamespaceLister {
	return globalReportTypeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GlobalReportTypeNamespaceLister helps list and get GlobalReportTypes.
// All objects returned here must be treated as read-only.
type GlobalReportTypeNamespaceLister interface {
	// List lists all GlobalReportTypes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.GlobalReportType, err error)
	// Get retrieves the GlobalReportType from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.GlobalReportType, error)
	GlobalReportTypeNamespaceListerExpansion
}

// globalReportTypeNamespaceLister implements the GlobalReportTypeNamespaceLister
// interface.
type globalReportTypeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GlobalReportTypes in the indexer for a given namespace.
func (s globalReportTypeNamespaceLister) List(selector labels.Selector) (ret []*v3.GlobalReportType, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.GlobalReportType))
	})
	return ret, err
}

// Get retrieves the GlobalReportType from the indexer for a given namespace and name.
func (s globalReportTypeNamespaceLister) Get(name string) (*v3.GlobalReportType, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("globalreporttype"), name)
	}
	return obj.(*v3.GlobalReportType), nil
}
