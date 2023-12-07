package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MyAppSpec defines the desired state of MyApp
type MyAppSpec struct {
	FrontendImage string `json:"frontendImage"`
	BackendImage  string `json:"backendImage"`
	AppImage      string `json:"appImage"`
}

// MyAppStatus defines the observed state of MyApp
type MyAppStatus struct {
	FrontendDeploymentStatus string `json:"frontendDeploymentStatus,omitempty"`
	BackendDeploymentStatus  string `json:"backendDeploymentStatus,omitempty"`
	AppDeploymentStatus      string `json:"appDeploymentStatus,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type MyApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyAppSpec   `json:"spec,omitempty"`
	Status MyAppStatus `json:"status,omitempty"`
}
