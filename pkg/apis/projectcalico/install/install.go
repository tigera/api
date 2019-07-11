// Copyright (c) 2019 Tigera, Inc. All rights reserved.

package install

import (
	"github.com/tigera/api/pkg/apis/projectcalico"
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
)

// Install registers the API group and adds types to a scheme
func Install(groupFactoryRegistry announced.APIGroupFactoryRegistry, registry *registered.APIRegistrationManager, scheme *runtime.Scheme) {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  projectcalico.GroupName,
			RootScopedKinds:            sets.NewString("GlobalReportType"),
			VersionPreferenceOrder:     []string{v3.SchemeGroupVersion.Version},
			AddInternalObjectsToScheme: projectcalico.AddToScheme,
		},
		announced.VersionToSchemeFunc{
			v3.SchemeGroupVersion.Version: v3.AddToScheme,
		},
	).Announce(groupFactoryRegistry).RegisterAndEnable(registry, scheme); err != nil {
		panic(err)
	}
}
