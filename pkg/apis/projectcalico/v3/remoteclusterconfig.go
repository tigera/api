// Copyright (c) 2018,2020-2021 Tigera, Inc. All rights reserved.

package v3

import (
	k8sv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindRemoteClusterConfiguration     = "RemoteClusterConfiguration"
	KindRemoteClusterConfigurationList = "RemoteClusterConfigurationList"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RemoteClusterConfiguration contains the configuration for remote clusters.
type RemoteClusterConfiguration struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the RemoteClusterConfiguration.
	Spec RemoteClusterConfigurationSpec `json:"spec,omitempty"`
}

// It's desirable to keep the list of things configurable here in sync with the other mechanism in apiconfig.go

// RemoteClusterConfigurationSpec contains the values of describing the cluster.
type RemoteClusterConfigurationSpec struct {
	// Indicates the datastore to use. If unspecified, defaults to etcdv3
	DatastoreType string `json:"datastoreType,omitempty" validate:"omitempty,datastoreType"`

	// Specifies a Secret to read for the RemoteClusterconfiguration.
	// If defined all datastore configuration in this struct will be cleared
	// and overwritten with the appropriate fields in the Secret.
	ClusterAccessSecret *k8sv1.ObjectReference `json:"clusterAccessSecret,omitempty" validate:"omitempty,clusterAccessSecret"`

	// Inline the ectd config fields
	EtcdConfig `json:",inline"`

	// Inline the k8s config fields.
	KubeConfig `json:",inline"`

	// Configuration options that do not relate to the underlying datastore connection. These fields relate to the
	// syncing of resources once the connection is established. These fields can be set independent of the other
	// connection-oriented fields, e.g. they can be set when ClusterAccessSecret is non-nil.
	// +kubebuilder:default={overlayRoutingMode: "Disabled"}
	SyncOptions RemoteClusterSyncOptions `json:"syncOptions,omitempty"`
}

type RemoteClusterSyncOptions struct {
	// Determines whether overlay routing will be established between federated clusters. If unspecified during create or
	// update of RemoteClusterConfiguration, this field will default based on the encapsulation mode of the local cluster
	// at the time of RemoteClusterConfiguration application: "Enabled" if VXLAN, "Disabled" otherwise. If upgrading from
	// a version that predates this field, this field will default to "Disabled".
	// +kubebuilder:default=Disabled
	OverlayRoutingMode OverlayRoutingMode `json:"overlayRoutingMode,omitempty" validate:"omitempty,oneof=Enabled Disabled"`
}

type OverlayRoutingMode string

const (
	OverlayRoutingModeEnabled  OverlayRoutingMode = "Enabled"
	OverlayRoutingModeDisabled OverlayRoutingMode = "Disabled"
)

type EtcdConfig struct {
	// A comma separated list of etcd endpoints. Valid if DatastoreType is etcdv3.  [Default: ]
	EtcdEndpoints string `json:"etcdEndpoints,omitempty" validate:"omitempty,etcdEndpoints"`
	// User name for RBAC. Valid if DatastoreType is etcdv3.
	EtcdUsername string `json:"etcdUsername,omitempty" validate:"omitempty"`
	// Password for the given user name. Valid if DatastoreType is etcdv3.
	EtcdPassword string `json:"etcdPassword,omitempty" validate:"omitempty"`
	// Path to the etcd key file. Valid if DatastoreType is etcdv3.
	EtcdKeyFile string `json:"etcdKeyFile,omitempty" validate:"omitempty,file"`
	// Path to the etcd client certificate. Valid if DatastoreType is etcdv3.
	EtcdCertFile string `json:"etcdCertFile,omitempty" validate:"omitempty,file"`
	// Path to the etcd Certificate Authority file. Valid if DatastoreType is etcdv3.
	EtcdCACertFile string `json:"etcdCACertFile,omitempty" validate:"omitempty,file"`
	// These config file parameters are to support inline certificates, keys and CA / Trusted certificate.
	EtcdKey    string `json:"etcdKey,omitempty" ignored:"true"`
	EtcdCert   string `json:"etcdCert,omitempty" ignored:"true"`
	EtcdCACert string `json:"etcdCACert,omitempty" ignored:"true"`
}

type KubeConfig struct {
	// When using the Kubernetes datastore, the location of a kubeconfig file. Valid if DatastoreType is kubernetes.
	Kubeconfig string `json:"kubeconfig,omitempty" validate:"omitempty,file"`
	// Location of the Kubernetes API. Not required if using kubeconfig. Valid if DatastoreType is kubernetes.
	K8sAPIEndpoint string `json:"k8sAPIEndpoint,omitempty" validate:"omitempty,k8sEndpoint"`
	// Location of a client key for accessing the Kubernetes API. Valid if DatastoreType is kubernetes.
	K8sKeyFile string `json:"k8sKeyFile,omitempty" validate:"omitempty,file"`
	// Location of a client certificate for accessing the Kubernetes API. Valid if DatastoreType is kubernetes.
	K8sCertFile string `json:"k8sCertFile,omitempty" validate:"omitempty,file"`
	// Location of a CA for accessing the Kubernetes API. Valid if DatastoreType is kubernetes.
	K8sCAFile string `json:"k8sCAFile,omitempty" validate:"omitempty,file"`
	// Token to be used for accessing the Kubernetes API. Valid if DatastoreType is kubernetes.
	K8sAPIToken              string `json:"k8sAPIToken,omitempty" validate:"omitempty"`
	K8sInsecureSkipTLSVerify bool   `json:"k8sInsecureSkipTLSVerify,omitempty" validate:"omitempty"`
	// This is an alternative to Kubeconfig and if specified overrides Kubeconfig.
	// This contains the contents that would normally be in the file pointed at by Kubeconfig.
	KubeconfigInline string `json:"kubeconfigInline,omitempty" ignored:"true"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RemoteClusterConfigurationList contains a list of RemoteClusterConfiguration resources
type RemoteClusterConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []RemoteClusterConfiguration `json:"items"`
}

// New RemoteClusterConfiguration creates a new (zeroed) RemoteClusterConfiguration struct with the TypeMetadata
// initialized to the current version.
func NewRemoteClusterConfiguration() *RemoteClusterConfiguration {
	return &RemoteClusterConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindRemoteClusterConfiguration,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewRemoteClusterConfigurationList creates a new (zeroed) RemoteClusterConfigurationList struct with the TypeMetadata
// initialized to the current version.
func NewRemoteClusterConfigurationList() *RemoteClusterConfigurationList {
	return &RemoteClusterConfigurationList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindRemoteClusterConfigurationList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
