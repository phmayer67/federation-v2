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

// FederatedSecretOverrideLister helps list FederatedSecretOverrides.
type FederatedSecretOverrideLister interface {
	// List lists all FederatedSecretOverrides in the indexer.
	List(selector labels.Selector) (ret []*federation.FederatedSecretOverride, err error)
	// FederatedSecretOverrides returns an object that can list and get FederatedSecretOverrides.
	FederatedSecretOverrides(namespace string) FederatedSecretOverrideNamespaceLister
	FederatedSecretOverrideListerExpansion
}

// federatedSecretOverrideLister implements the FederatedSecretOverrideLister interface.
type federatedSecretOverrideLister struct {
	indexer cache.Indexer
}

// NewFederatedSecretOverrideLister returns a new FederatedSecretOverrideLister.
func NewFederatedSecretOverrideLister(indexer cache.Indexer) FederatedSecretOverrideLister {
	return &federatedSecretOverrideLister{indexer: indexer}
}

// List lists all FederatedSecretOverrides in the indexer.
func (s *federatedSecretOverrideLister) List(selector labels.Selector) (ret []*federation.FederatedSecretOverride, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*federation.FederatedSecretOverride))
	})
	return ret, err
}

// FederatedSecretOverrides returns an object that can list and get FederatedSecretOverrides.
func (s *federatedSecretOverrideLister) FederatedSecretOverrides(namespace string) FederatedSecretOverrideNamespaceLister {
	return federatedSecretOverrideNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedSecretOverrideNamespaceLister helps list and get FederatedSecretOverrides.
type FederatedSecretOverrideNamespaceLister interface {
	// List lists all FederatedSecretOverrides in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*federation.FederatedSecretOverride, err error)
	// Get retrieves the FederatedSecretOverride from the indexer for a given namespace and name.
	Get(name string) (*federation.FederatedSecretOverride, error)
	FederatedSecretOverrideNamespaceListerExpansion
}

// federatedSecretOverrideNamespaceLister implements the FederatedSecretOverrideNamespaceLister
// interface.
type federatedSecretOverrideNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedSecretOverrides in the indexer for a given namespace.
func (s federatedSecretOverrideNamespaceLister) List(selector labels.Selector) (ret []*federation.FederatedSecretOverride, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*federation.FederatedSecretOverride))
	})
	return ret, err
}

// Get retrieves the FederatedSecretOverride from the indexer for a given namespace and name.
func (s federatedSecretOverrideNamespaceLister) Get(name string) (*federation.FederatedSecretOverride, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(federation.Resource("federatedsecretoverride"), name)
	}
	return obj.(*federation.FederatedSecretOverride), nil
}
