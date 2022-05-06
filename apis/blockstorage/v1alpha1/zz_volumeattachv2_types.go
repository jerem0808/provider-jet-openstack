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

type VolumeAttachV2Observation struct {
	DriverVolumeType *string `json:"driverVolumeType,omitempty" tf:"driver_volume_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MountPointBase *string `json:"mountPointBase,omitempty" tf:"mount_point_base,omitempty"`
}

type VolumeAttachV2Parameters struct {

	// +kubebuilder:validation:Optional
	AttachMode *string `json:"attachMode,omitempty" tf:"attach_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	// +kubebuilder:validation:Required
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	Initiator *string `json:"initiator,omitempty" tf:"initiator,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	Multipath *bool `json:"multipath,omitempty" tf:"multipath,omitempty"`

	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// +kubebuilder:validation:Optional
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	VolumeID *string `json:"volumeId" tf:"volume_id,omitempty"`

	// +kubebuilder:validation:Optional
	Wwnn *string `json:"wwnn,omitempty" tf:"wwnn,omitempty"`

	// +kubebuilder:validation:Optional
	Wwpn []*string `json:"wwpn,omitempty" tf:"wwpn,omitempty"`
}

// VolumeAttachV2Spec defines the desired state of VolumeAttachV2
type VolumeAttachV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeAttachV2Parameters `json:"forProvider"`
}

// VolumeAttachV2Status defines the observed state of VolumeAttachV2.
type VolumeAttachV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeAttachV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeAttachV2 is the Schema for the VolumeAttachV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type VolumeAttachV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeAttachV2Spec   `json:"spec"`
	Status            VolumeAttachV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeAttachV2List contains a list of VolumeAttachV2s
type VolumeAttachV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeAttachV2 `json:"items"`
}

// Repository type metadata.
var (
	VolumeAttachV2_Kind             = "VolumeAttachV2"
	VolumeAttachV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeAttachV2_Kind}.String()
	VolumeAttachV2_KindAPIVersion   = VolumeAttachV2_Kind + "." + CRDGroupVersion.String()
	VolumeAttachV2_GroupVersionKind = CRDGroupVersion.WithKind(VolumeAttachV2_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeAttachV2{}, &VolumeAttachV2List{})
}