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

// This file was automatically generated by informer-gen

package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	apps_v1beta1 "k8s.io/kubernetes/pkg/apis/apps/v1beta1"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/internalinterfaces"
	v1beta1 "k8s.io/kubernetes/pkg/client/listers/apps/v1beta1"
	time "time"
)

// StatefulSetInformer provides access to a shared informer and lister for
// StatefulSets.
type StatefulSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.StatefulSetLister
}

type statefulSetInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newStatefulSetInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.AppsV1beta1().StatefulSets(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.AppsV1beta1().StatefulSets(v1.NamespaceAll).Watch(options)
			},
		},
		&apps_v1beta1.StatefulSet{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *statefulSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apps_v1beta1.StatefulSet{}, newStatefulSetInformer)
}

func (f *statefulSetInformer) Lister() v1beta1.StatefulSetLister {
	return v1beta1.NewStatefulSetLister(f.Informer().GetIndexer())
}
