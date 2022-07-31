package v1aplha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HADeployment struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec HADeploymentSpec
}

type HADeploymentSpec struct {
	image  string
	labels map[string]interface{}
}

type HADeploymentList struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Items []HADeployment
}
