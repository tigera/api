package projectcalico

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
)

const (
	GroupName = "projectcalico.org"
)

var (
	schemeBuilder = runtime.NewSchemeBuilder(v3.Install)
	// Install is a function which adds every version of this group to a scheme
	Install = schemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return schema.GroupResource{Group: GroupName, Resource: resource}
}

func Kind(kind string) schema.GroupKind {
	return schema.GroupKind{Group: GroupName, Kind: kind}
}
