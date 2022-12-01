package insights

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	insightsv1alpha1 "github.com/openshift/api/insights/v1alpha1"
)

const (
	GroupName = "insights.openshift.io"
)

var (
	schemeBuilder = runtime.NewSchemeBuilder(insightsv1alpha1.Install)
	// Install is a function which adds every version of this group to a scheme
	Install = schemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return schema.GroupResource{Group: GroupName, Resource: resource}
}

func Kind(kind string) schema.GroupKind {
	return schema.GroupKind{Group: GroupName, Kind: kind}
}
