/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1alpha1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/meta/v1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1beta1]

package v1alpha1

import (
	"k8s.io/kubernetes/pkg/api/errors"
	rbac "k8s.io/kubernetes/pkg/apis/rbac"
	v1alpha1 "k8s.io/kubernetes/pkg/apis/rbac/v1alpha1"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/labels"
)

// RoleBindingLister helps list RoleBindings.
type RoleBindingLister interface {
	// List lists all RoleBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RoleBinding, err error)
	// RoleBindings returns an object that can list and get RoleBindings.
	RoleBindings(namespace string) RoleBindingNamespaceLister
	RoleBindingListerExpansion
}

// roleBindingLister implements the RoleBindingLister interface.
type roleBindingLister struct {
	indexer cache.Indexer
}

// NewRoleBindingLister returns a new RoleBindingLister.
func NewRoleBindingLister(indexer cache.Indexer) RoleBindingLister {
	return &roleBindingLister{indexer: indexer}
}

// List lists all RoleBindings in the indexer.
func (s *roleBindingLister) List(selector labels.Selector) (ret []*v1alpha1.RoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RoleBinding))
	})
	return ret, err
}

// RoleBindings returns an object that can list and get RoleBindings.
func (s *roleBindingLister) RoleBindings(namespace string) RoleBindingNamespaceLister {
	return roleBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RoleBindingNamespaceLister helps list and get RoleBindings.
type RoleBindingNamespaceLister interface {
	// List lists all RoleBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RoleBinding, err error)
	// Get retrieves the RoleBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RoleBinding, error)
	RoleBindingNamespaceListerExpansion
}

// roleBindingNamespaceLister implements the RoleBindingNamespaceLister
// interface.
type roleBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RoleBindings in the indexer for a given namespace.
func (s roleBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RoleBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RoleBinding))
	})
	return ret, err
}

// Get retrieves the RoleBinding from the indexer for a given namespace and name.
func (s roleBindingNamespaceLister) Get(name string) (*v1alpha1.RoleBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(rbac.Resource("rolebinding"), name)
	}
	return obj.(*v1alpha1.RoleBinding), nil
}
