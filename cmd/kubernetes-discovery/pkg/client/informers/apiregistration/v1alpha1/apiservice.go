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

// This file was automatically generated by informer-gen with arguments: --input-dirs=[k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/apis/apiregistration,k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/apis/apiregistration/v1alpha1] --internal-clientset-package=k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/client/clientset_generated/internalclientset --listers-package=k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/client/listers --output-package=k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/client/informers --versioned-clientset-package=k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/client/clientset_generated/clientset

package v1alpha1

import (
	apiregistration_v1alpha1 "k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/apis/apiregistration/v1alpha1"
	clientset "k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/client/clientset_generated/clientset"
	internalinterfaces "k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/client/informers/internalinterfaces"
	v1alpha1 "k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/client/listers/apiregistration/v1alpha1"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	cache "k8s.io/kubernetes/pkg/client/cache"
	runtime "k8s.io/kubernetes/pkg/runtime"
	watch "k8s.io/kubernetes/pkg/watch"
	time "time"
)

// APIServiceInformer provides access to a shared informer and lister for
// APIServices.
type APIServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.APIServiceLister
}

type aPIServiceInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newAPIServiceInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.ApiregistrationV1alpha1().APIServices().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.ApiregistrationV1alpha1().APIServices().Watch(options)
			},
		},
		&apiregistration_v1alpha1.APIService{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *aPIServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.VersionedInformerFor(&apiregistration_v1alpha1.APIService{}, newAPIServiceInformer)
}

func (f *aPIServiceInformer) Lister() v1alpha1.APIServiceLister {
	return v1alpha1.NewAPIServiceLister(f.Informer().GetIndexer())
}
