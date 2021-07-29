// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RemoteClusterConfigurationLister helps list RemoteClusterConfigurations.
type RemoteClusterConfigurationLister interface {
	// List lists all RemoteClusterConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v3.RemoteClusterConfiguration, err error)
	// Get retrieves the RemoteClusterConfiguration from the index for a given name.
	Get(name string) (*v3.RemoteClusterConfiguration, error)
	RemoteClusterConfigurationListerExpansion
}

// remoteClusterConfigurationLister implements the RemoteClusterConfigurationLister interface.
type remoteClusterConfigurationLister struct {
	indexer cache.Indexer
}

// NewRemoteClusterConfigurationLister returns a new RemoteClusterConfigurationLister.
func NewRemoteClusterConfigurationLister(indexer cache.Indexer) RemoteClusterConfigurationLister {
	return &remoteClusterConfigurationLister{indexer: indexer}
}

// List lists all RemoteClusterConfigurations in the indexer.
func (s *remoteClusterConfigurationLister) List(selector labels.Selector) (ret []*v3.RemoteClusterConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.RemoteClusterConfiguration))
	})
	return ret, err
}

// Get retrieves the RemoteClusterConfiguration from the index for a given name.
func (s *remoteClusterConfigurationLister) Get(name string) (*v3.RemoteClusterConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("remoteclusterconfiguration"), name)
	}
	return obj.(*v3.RemoteClusterConfiguration), nil
}
