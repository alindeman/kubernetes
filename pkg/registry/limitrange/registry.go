/*
Copyright 2014 Google Inc. All rights reserved.

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

package limitrange

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/registry/generic"
	etcdgeneric "github.com/GoogleCloudPlatform/kubernetes/pkg/registry/generic/etcd"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/tools"
)

// registry implements custom changes to generic.Etcd.
type registry struct {
	*etcdgeneric.Etcd
}

// NewEtcdRegistry returns a registry which will store LimitRange in the given helper
func NewEtcdRegistry(h tools.EtcdHelper) generic.Registry {
	prefix := "/limitranges"
	return registry{
		Etcd: &etcdgeneric.Etcd{
			NewFunc:      func() runtime.Object { return &api.LimitRange{} },
			NewListFunc:  func() runtime.Object { return &api.LimitRangeList{} },
			EndpointName: "limitranges",
			KeyRootFunc: func(ctx api.Context) string {
				return etcdgeneric.NamespaceKeyRootFunc(ctx, prefix)
			},
			KeyFunc: func(ctx api.Context, id string) (string, error) {
				return etcdgeneric.NamespaceKeyFunc(ctx, prefix, id)
			},
			Helper: h,
		},
	}
}
