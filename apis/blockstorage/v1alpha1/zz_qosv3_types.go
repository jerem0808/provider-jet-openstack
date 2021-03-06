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

type QosV3Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type QosV3Parameters struct {

	// +kubebuilder:validation:Optional
	Consumer *string `json:"consumer,omitempty" tf:"consumer,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Specs map[string]*string `json:"specs,omitempty" tf:"specs,omitempty"`
}

// QosV3Spec defines the desired state of QosV3
type QosV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QosV3Parameters `json:"forProvider"`
}

// QosV3Status defines the observed state of QosV3.
type QosV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QosV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QosV3 is the Schema for the QosV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type QosV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QosV3Spec   `json:"spec"`
	Status            QosV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QosV3List contains a list of QosV3s
type QosV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QosV3 `json:"items"`
}

// Repository type metadata.
var (
	QosV3_Kind             = "QosV3"
	QosV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QosV3_Kind}.String()
	QosV3_KindAPIVersion   = QosV3_Kind + "." + CRDGroupVersion.String()
	QosV3_GroupVersionKind = CRDGroupVersion.WithKind(QosV3_Kind)
)

func init() {
	SchemeBuilder.Register(&QosV3{}, &QosV3List{})
}
