/*


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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StarshipSpec defines the desired state of Starship
type StarshipSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name of Starship. Edit Starship_types.go to remove/update
	Name string `json:"name,omitempty"`
}

// StarshipStatus defines the observed state of Starship
type StarshipStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Name       string `json:"name,omitempty"`
	Model      string `json:"model,omitempty"`
	Crew       string `json:"crew,omitempty"`
	Passengers string `json:"passengers,omitempty"`
	Costs      string `json:"costs,omitempty"`
	Capacity   string `json:"capacity,omitempty"`
}

// +kubebuilder:object:root=true

// Starship is the Schema for the starships API
type Starship struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StarshipSpec   `json:"spec,omitempty"`
	Status StarshipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StarshipList contains a list of Starship
type StarshipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Starship `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Starship{}, &StarshipList{})
}
