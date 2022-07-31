package v1aplha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{
	Group:   "sachinmaurya.tech",
	Version: "v1alpha1",
}

var (
	SchemeBuilder runtime.SchemeBuilder
)

func init() {
	SchemeBuilder.Register(addKnowTypes)
}

func addKnowTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, &HADeployment{}, &HADeploymentList{})

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
