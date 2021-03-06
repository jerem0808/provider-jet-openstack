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

type ImageAccessV2Observation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ImageAccessV2Parameters struct {

	// +kubebuilder:validation:Required
	ImageID *string `json:"imageId" tf:"image_id,omitempty"`

	// +kubebuilder:validation:Required
	MemberID *string `json:"memberId" tf:"member_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// ImageAccessV2Spec defines the desired state of ImageAccessV2
type ImageAccessV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageAccessV2Parameters `json:"forProvider"`
}

// ImageAccessV2Status defines the observed state of ImageAccessV2.
type ImageAccessV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageAccessV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ImageAccessV2 is the Schema for the ImageAccessV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type ImageAccessV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageAccessV2Spec   `json:"spec"`
	Status            ImageAccessV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageAccessV2List contains a list of ImageAccessV2s
type ImageAccessV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageAccessV2 `json:"items"`
}

// Repository type metadata.
var (
	ImageAccessV2_Kind             = "ImageAccessV2"
	ImageAccessV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageAccessV2_Kind}.String()
	ImageAccessV2_KindAPIVersion   = ImageAccessV2_Kind + "." + CRDGroupVersion.String()
	ImageAccessV2_GroupVersionKind = CRDGroupVersion.WithKind(ImageAccessV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageAccessV2{}, &ImageAccessV2List{})
}
