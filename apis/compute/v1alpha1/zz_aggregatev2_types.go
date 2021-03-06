/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AggregateV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AggregateV2Parameters struct {

	// +kubebuilder:validation:Optional
	Hosts []*string `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// AggregateV2Spec defines the desired state of AggregateV2
type AggregateV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AggregateV2Parameters `json:"forProvider"`
}

// AggregateV2Status defines the observed state of AggregateV2.
type AggregateV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AggregateV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AggregateV2 is the Schema for the AggregateV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type AggregateV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AggregateV2Spec   `json:"spec"`
	Status            AggregateV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AggregateV2List contains a list of AggregateV2s
type AggregateV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AggregateV2 `json:"items"`
}

// Repository type metadata.
var (
	AggregateV2_Kind             = "AggregateV2"
	AggregateV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AggregateV2_Kind}.String()
	AggregateV2_KindAPIVersion   = AggregateV2_Kind + "." + CRDGroupVersion.String()
	AggregateV2_GroupVersionKind = CRDGroupVersion.WithKind(AggregateV2_Kind)
)

func init() {
	SchemeBuilder.Register(&AggregateV2{}, &AggregateV2List{})
}
