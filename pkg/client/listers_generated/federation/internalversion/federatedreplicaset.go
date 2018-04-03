/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by lister-gen

package internalversion

import (
	federation "github.com/marun/federation-v2/pkg/apis/federation"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedReplicaSetLister helps list FederatedReplicaSets.
type FederatedReplicaSetLister interface {
	// List lists all FederatedReplicaSets in the indexer.
	List(selector labels.Selector) (ret []*federation.FederatedReplicaSet, err error)
	// FederatedReplicaSets returns an object that can list and get FederatedReplicaSets.
	FederatedReplicaSets(namespace string) FederatedReplicaSetNamespaceLister
	FederatedReplicaSetListerExpansion
}

// federatedReplicaSetLister implements the FederatedReplicaSetLister interface.
type federatedReplicaSetLister struct {
	indexer cache.Indexer
}

// NewFederatedReplicaSetLister returns a new FederatedReplicaSetLister.
func NewFederatedReplicaSetLister(indexer cache.Indexer) FederatedReplicaSetLister {
	return &federatedReplicaSetLister{indexer: indexer}
}

// List lists all FederatedReplicaSets in the indexer.
func (s *federatedReplicaSetLister) List(selector labels.Selector) (ret []*federation.FederatedReplicaSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*federation.FederatedReplicaSet))
	})
	return ret, err
}

// FederatedReplicaSets returns an object that can list and get FederatedReplicaSets.
func (s *federatedReplicaSetLister) FederatedReplicaSets(namespace string) FederatedReplicaSetNamespaceLister {
	return federatedReplicaSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedReplicaSetNamespaceLister helps list and get FederatedReplicaSets.
type FederatedReplicaSetNamespaceLister interface {
	// List lists all FederatedReplicaSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*federation.FederatedReplicaSet, err error)
	// Get retrieves the FederatedReplicaSet from the indexer for a given namespace and name.
	Get(name string) (*federation.FederatedReplicaSet, error)
	FederatedReplicaSetNamespaceListerExpansion
}

// federatedReplicaSetNamespaceLister implements the FederatedReplicaSetNamespaceLister
// interface.
type federatedReplicaSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedReplicaSets in the indexer for a given namespace.
func (s federatedReplicaSetNamespaceLister) List(selector labels.Selector) (ret []*federation.FederatedReplicaSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*federation.FederatedReplicaSet))
	})
	return ret, err
}

// Get retrieves the FederatedReplicaSet from the indexer for a given namespace and name.
func (s federatedReplicaSetNamespaceLister) Get(name string) (*federation.FederatedReplicaSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(federation.Resource("federatedreplicaset"), name)
	}
	return obj.(*federation.FederatedReplicaSet), nil
}
