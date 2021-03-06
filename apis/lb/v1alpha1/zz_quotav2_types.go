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

type QuotaV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type QuotaV2Parameters struct {

	// +kubebuilder:validation:Optional
	HealthMonitor *float64 `json:"healthMonitor,omitempty" tf:"health_monitor,omitempty"`

	// +kubebuilder:validation:Optional
	L7Policy *float64 `json:"l7Policy,omitempty" tf:"l7_policy,omitempty"`

	// +kubebuilder:validation:Optional
	L7Rule *float64 `json:"l7Rule,omitempty" tf:"l7_rule,omitempty"`

	// +kubebuilder:validation:Optional
	Listener *float64 `json:"listener,omitempty" tf:"listener,omitempty"`

	// +kubebuilder:validation:Optional
	Loadbalancer *float64 `json:"loadbalancer,omitempty" tf:"loadbalancer,omitempty"`

	// +kubebuilder:validation:Optional
	Member *float64 `json:"member,omitempty" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Pool *float64 `json:"pool,omitempty" tf:"pool,omitempty"`

	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// QuotaV2Spec defines the desired state of QuotaV2
type QuotaV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QuotaV2Parameters `json:"forProvider"`
}

// QuotaV2Status defines the observed state of QuotaV2.
type QuotaV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QuotaV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QuotaV2 is the Schema for the QuotaV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type QuotaV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QuotaV2Spec   `json:"spec"`
	Status            QuotaV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QuotaV2List contains a list of QuotaV2s
type QuotaV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuotaV2 `json:"items"`
}

// Repository type metadata.
var (
	QuotaV2_Kind             = "QuotaV2"
	QuotaV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QuotaV2_Kind}.String()
	QuotaV2_KindAPIVersion   = QuotaV2_Kind + "." + CRDGroupVersion.String()
	QuotaV2_GroupVersionKind = CRDGroupVersion.WithKind(QuotaV2_Kind)
)

func init() {
	SchemeBuilder.Register(&QuotaV2{}, &QuotaV2List{})
}
