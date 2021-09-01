// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"github.com/tigera/api/pkg/client/clientset_generated/clientset/scheme"
	rest "k8s.io/client-go/rest"
)

type ProjectcalicoV3Interface interface {
	RESTClient() rest.Interface
	AuthenticationReviewsGetter
	AuthorizationReviewsGetter
	BGPConfigurationsGetter
	BGPPeersGetter
	ClusterInformationsGetter
	DeepPacketInspectionsGetter
	FelixConfigurationsGetter
	GlobalAlertsGetter
	GlobalAlertTemplatesGetter
	GlobalNetworkPoliciesGetter
	GlobalNetworkSetsGetter
	GlobalReportsGetter
	GlobalReportTypesGetter
	GlobalThreatFeedsGetter
	HostEndpointsGetter
	IPPoolsGetter
	KubeControllersConfigurationsGetter
	LicenseKeysGetter
	ManagedClustersGetter
	NetworkPoliciesGetter
	NetworkSetsGetter
	PacketCapturesGetter
	ProfilesGetter
	RemoteClusterConfigurationsGetter
	StagedGlobalNetworkPoliciesGetter
	StagedKubernetesNetworkPoliciesGetter
	StagedNetworkPoliciesGetter
	TiersGetter
	UISettingsGroupsGetter
}

// ProjectcalicoV3Client is used to interact with features provided by the projectcalico.org group.
type ProjectcalicoV3Client struct {
	restClient rest.Interface
}

func (c *ProjectcalicoV3Client) AuthenticationReviews() AuthenticationReviewInterface {
	return newAuthenticationReviews(c)
}

func (c *ProjectcalicoV3Client) AuthorizationReviews() AuthorizationReviewInterface {
	return newAuthorizationReviews(c)
}

func (c *ProjectcalicoV3Client) BGPConfigurations() BGPConfigurationInterface {
	return newBGPConfigurations(c)
}

func (c *ProjectcalicoV3Client) BGPPeers() BGPPeerInterface {
	return newBGPPeers(c)
}

func (c *ProjectcalicoV3Client) ClusterInformations() ClusterInformationInterface {
	return newClusterInformations(c)
}

func (c *ProjectcalicoV3Client) DeepPacketInspections(namespace string) DeepPacketInspectionInterface {
	return newDeepPacketInspections(c, namespace)
}

func (c *ProjectcalicoV3Client) FelixConfigurations() FelixConfigurationInterface {
	return newFelixConfigurations(c)
}

func (c *ProjectcalicoV3Client) GlobalAlerts() GlobalAlertInterface {
	return newGlobalAlerts(c)
}

func (c *ProjectcalicoV3Client) GlobalAlertTemplates() GlobalAlertTemplateInterface {
	return newGlobalAlertTemplates(c)
}

func (c *ProjectcalicoV3Client) GlobalNetworkPolicies() GlobalNetworkPolicyInterface {
	return newGlobalNetworkPolicies(c)
}

func (c *ProjectcalicoV3Client) GlobalNetworkSets() GlobalNetworkSetInterface {
	return newGlobalNetworkSets(c)
}

func (c *ProjectcalicoV3Client) GlobalReports() GlobalReportInterface {
	return newGlobalReports(c)
}

func (c *ProjectcalicoV3Client) GlobalReportTypes() GlobalReportTypeInterface {
	return newGlobalReportTypes(c)
}

func (c *ProjectcalicoV3Client) GlobalThreatFeeds() GlobalThreatFeedInterface {
	return newGlobalThreatFeeds(c)
}

func (c *ProjectcalicoV3Client) HostEndpoints() HostEndpointInterface {
	return newHostEndpoints(c)
}

func (c *ProjectcalicoV3Client) IPPools() IPPoolInterface {
	return newIPPools(c)
}

func (c *ProjectcalicoV3Client) KubeControllersConfigurations() KubeControllersConfigurationInterface {
	return newKubeControllersConfigurations(c)
}

func (c *ProjectcalicoV3Client) LicenseKeys() LicenseKeyInterface {
	return newLicenseKeys(c)
}

func (c *ProjectcalicoV3Client) ManagedClusters() ManagedClusterInterface {
	return newManagedClusters(c)
}

func (c *ProjectcalicoV3Client) NetworkPolicies(namespace string) NetworkPolicyInterface {
	return newNetworkPolicies(c, namespace)
}

func (c *ProjectcalicoV3Client) NetworkSets(namespace string) NetworkSetInterface {
	return newNetworkSets(c, namespace)
}

func (c *ProjectcalicoV3Client) PacketCaptures(namespace string) PacketCaptureInterface {
	return newPacketCaptures(c, namespace)
}

func (c *ProjectcalicoV3Client) Profiles() ProfileInterface {
	return newProfiles(c)
}

func (c *ProjectcalicoV3Client) RemoteClusterConfigurations() RemoteClusterConfigurationInterface {
	return newRemoteClusterConfigurations(c)
}

func (c *ProjectcalicoV3Client) StagedGlobalNetworkPolicies() StagedGlobalNetworkPolicyInterface {
	return newStagedGlobalNetworkPolicies(c)
}

func (c *ProjectcalicoV3Client) StagedKubernetesNetworkPolicies(namespace string) StagedKubernetesNetworkPolicyInterface {
	return newStagedKubernetesNetworkPolicies(c, namespace)
}

func (c *ProjectcalicoV3Client) StagedNetworkPolicies(namespace string) StagedNetworkPolicyInterface {
	return newStagedNetworkPolicies(c, namespace)
}

func (c *ProjectcalicoV3Client) Tiers() TierInterface {
	return newTiers(c)
}

func (c *ProjectcalicoV3Client) UISettingsGroups() UISettingsGroupInterface {
	return newUISettingsGroups(c)
}

// NewForConfig creates a new ProjectcalicoV3Client for the given config.
func NewForConfig(c *rest.Config) (*ProjectcalicoV3Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ProjectcalicoV3Client{client}, nil
}

// NewForConfigOrDie creates a new ProjectcalicoV3Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ProjectcalicoV3Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ProjectcalicoV3Client for the given RESTClient.
func New(c rest.Interface) *ProjectcalicoV3Client {
	return &ProjectcalicoV3Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v3.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ProjectcalicoV3Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
