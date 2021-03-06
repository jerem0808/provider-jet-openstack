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

type ZoneV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ZoneV2Parameters struct {

	// +kubebuilder:validation:Optional
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DisableStatusCheck *bool `json:"disableStatusCheck,omitempty" tf:"disable_status_check,omitempty"`

	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	Masters []*string `json:"masters,omitempty" tf:"masters,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// ZoneV2Spec defines the desired state of ZoneV2
type ZoneV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ZoneV2Parameters `json:"forProvider"`
}

// ZoneV2Status defines the observed state of ZoneV2.
type ZoneV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ZoneV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneV2 is the Schema for the ZoneV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type ZoneV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ZoneV2Spec   `json:"spec"`
	Status            ZoneV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneV2List contains a list of ZoneV2s
type ZoneV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZoneV2 `json:"items"`
}

// Repository type metadata.
var (
	ZoneV2_Kind             = "ZoneV2"
	ZoneV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ZoneV2_Kind}.String()
	ZoneV2_KindAPIVersion   = ZoneV2_Kind + "." + CRDGroupVersion.String()
	ZoneV2_GroupVersionKind = CRDGroupVersion.WithKind(ZoneV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ZoneV2{}, &ZoneV2List{})
}
