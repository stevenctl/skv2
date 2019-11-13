package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CanaryDeployment is the Schema for the canaryDeployment API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=canaryDeployment,scope=Namespaced
type CanaryDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CanaryDeploymentSpec   `json:"spec,omitempty"`
	Status CanaryDeploymentStatus `json:"status,omitempty"`
}

// CanaryDeploymentStatus defines an observed condition of CanaryDeployment
// +k8s:openapi-gen=true
type CanaryDeploymentStatus struct {
	// Phase represents the current state of the CanaryDeployment.
	Phase CanaryDeploymentPhase `json:"phase,omitempty"`

	// ObservedGeneration was the last metadata.generation of the CanaryDeployment
	// observed by the operator. If this does not match the metadata.generation of the CanaryDeployment,
	// it means the operator has not yet reconciled the current generation of the operator
	ObservedGeneration int64 `json:"observedGeneration",omitempty`

	// StatusInfo defines the observed state of the CanaryDeployment in the cluster
	CanaryDeploymentStatusInfo
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CanaryDeploymentList contains a list of CanaryDeployment
type CanaryDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CanaryDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CanaryDeployment{}, &CanaryDeploymentList{})
}