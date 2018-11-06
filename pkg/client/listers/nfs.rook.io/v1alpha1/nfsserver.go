/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/rook/rook/pkg/apis/nfs.rook.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NFSServerLister helps list NFSServers.
type NFSServerLister interface {
	// List lists all NFSServers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NFSServer, err error)
	// NFSServers returns an object that can list and get NFSServers.
	NFSServers(namespace string) NFSServerNamespaceLister
	NFSServerListerExpansion
}

// nFSServerLister implements the NFSServerLister interface.
type nFSServerLister struct {
	indexer cache.Indexer
}

// NewNFSServerLister returns a new NFSServerLister.
func NewNFSServerLister(indexer cache.Indexer) NFSServerLister {
	return &nFSServerLister{indexer: indexer}
}

// List lists all NFSServers in the indexer.
func (s *nFSServerLister) List(selector labels.Selector) (ret []*v1alpha1.NFSServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NFSServer))
	})
	return ret, err
}

// NFSServers returns an object that can list and get NFSServers.
func (s *nFSServerLister) NFSServers(namespace string) NFSServerNamespaceLister {
	return nFSServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NFSServerNamespaceLister helps list and get NFSServers.
type NFSServerNamespaceLister interface {
	// List lists all NFSServers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NFSServer, err error)
	// Get retrieves the NFSServer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NFSServer, error)
	NFSServerNamespaceListerExpansion
}

// nFSServerNamespaceLister implements the NFSServerNamespaceLister
// interface.
type nFSServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NFSServers in the indexer for a given namespace.
func (s nFSServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NFSServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NFSServer))
	})
	return ret, err
}

// Get retrieves the NFSServer from the indexer for a given namespace and name.
func (s nFSServerNamespaceLister) Get(name string) (*v1alpha1.NFSServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("nfsserver"), name)
	}
	return obj.(*v1alpha1.NFSServer), nil
}