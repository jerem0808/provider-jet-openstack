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

type RoleV3Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RoleV3Parameters struct {

	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// RoleV3Spec defines the desired state of RoleV3
type RoleV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleV3Parameters `json:"forProvider"`
}

// RoleV3Status defines the observed state of RoleV3.
type RoleV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleV3 is the Schema for the RoleV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type RoleV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleV3Spec   `json:"spec"`
	Status            RoleV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleV3List contains a list of RoleV3s
type RoleV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleV3 `json:"items"`
}

// Repository type metadata.
var (
	RoleV3_Kind             = "RoleV3"
	RoleV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleV3_Kind}.String()
	RoleV3_KindAPIVersion   = RoleV3_Kind + "." + CRDGroupVersion.String()
	RoleV3_GroupVersionKind = CRDGroupVersion.WithKind(RoleV3_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleV3{}, &RoleV3List{})
}