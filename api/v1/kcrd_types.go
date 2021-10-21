/*
Copyright 2021.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KcrdSpec defines the desired state of Kcrd
type KcrdSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Name       string   `json:"name"`
	Finder     string   `json:"finder"`
	Domain     string   `json:"domain"`
	Image      string   `json:"image"`
	Port       int      `json:"port"`
	TargetPort int      `json:"target-port"`
	Paths      []string `json:"paths"`
}

// KcrdStatus defines the observed state of Kcrd
type KcrdStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	AllReady bool `json:"all_ready"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Kcrd is the Schema for the kcrds API
type Kcrd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KcrdSpec   `json:"spec,omitempty"`
	Status KcrdStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KcrdList contains a list of Kcrd
type KcrdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kcrd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Kcrd{}, &KcrdList{})
}
