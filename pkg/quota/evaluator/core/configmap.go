/*
Copyright 2016 The Kubernetes Authors.

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

package core

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/v1"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
	"k8s.io/kubernetes/pkg/quota"
	"k8s.io/kubernetes/pkg/quota/generic"
)

// NewConfigMapEvaluator returns an evaluator that can evaluate configMaps
func NewConfigMapEvaluator(kubeClient clientset.Interface) quota.Evaluator {
	return &generic.ObjectCountEvaluator{
		AllowCreateOnUpdate: false,
		InternalGroupKind:   api.Kind("ConfigMap"),
		ResourceName:        api.ResourceConfigMaps,
		ListFuncByNamespace: func(namespace string, options v1.ListOptions) ([]runtime.Object, error) {
			itemList, err := kubeClient.Core().ConfigMaps(namespace).List(options)
			if err != nil {
				return nil, err
			}
			results := make([]runtime.Object, 0, len(itemList.Items))
			for i := range itemList.Items {
				results = append(results, &itemList.Items[i])
			}
			return results, nil
		},
	}
}
