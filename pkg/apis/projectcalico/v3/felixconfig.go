// Copyright (c) 2017-2022 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v3

import (
	"github.com/tigera/api/pkg/lib/numorstring"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FelixConfigurationList contains a list of FelixConfiguration object.
type FelixConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []FelixConfiguration `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FelixConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec FelixConfigurationSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

const (
	KindFelixConfiguration     = "FelixConfiguration"
	KindFelixConfigurationList = "FelixConfigurationList"
)

type IptablesBackend string

const (
	IptablesBackendLegacy   IptablesBackend = "Legacy"
	IptablesBackendNFTables IptablesBackend = "NFT"
	IptablesBackendAuto     IptablesBackend = "Auto"
)

// NFTablesMode is the enum used to enable/disable nftables mode.
// +enum
type NFTablesMode string

const (
	NFTablesModeEnabled  = "Enabled"
	NFTablesModeDisabled = "Disabled"
)

// +kubebuilder:validation:Enum=DoNothing;Enable;Disable
type AWSSrcDstCheckOption string

const (
	AWSSrcDstCheckOptionDoNothing AWSSrcDstCheckOption = "DoNothing"
	AWSSrcDstCheckOptionEnable    AWSSrcDstCheckOption = "Enable"
	AWSSrcDstCheckOptionDisable   AWSSrcDstCheckOption = "Disable"
)

const (
	AWSSecondaryIPEnabled               = "Enabled"
	AWSSecondaryIPDisabled              = "Disabled"
	AWSSecondaryIPEnabledENIPerWorkload = "EnabledENIPerWorkload"
)

// +kubebuilder:validation:Enum=NoDelay;DelayDeniedPacket;DelayDNSResponse
type DNSPolicyMode string

const (
	DNSPolicyModeNoDelay           DNSPolicyMode = "NoDelay"
	DNSPolicyModeDelayDeniedPacket DNSPolicyMode = "DelayDeniedPacket"
	DNSPolicyModeDelayDNSResponse  DNSPolicyMode = "DelayDNSResponse"
)

type BPFDNSPolicyMode string

const (
	BPFDNSPolicyModeNoDelay BPFDNSPolicyMode = "NoDelay"
	BPFDNSPolicyModeInline  BPFDNSPolicyMode = "Inline"
)

// +kubebuilder:validation:Enum=Enabled;Disabled
type FloatingIPType string

const (
	FloatingIPsEnabled  FloatingIPType = "Enabled"
	FloatingIPsDisabled FloatingIPType = "Disabled"
)

type BPFHostNetworkedNATType string

const (
	BPFHostNetworkedNATEnabled  BPFHostNetworkedNATType = "Enabled"
	BPFHostNetworkedNATDisabled BPFHostNetworkedNATType = "Disabled"
)

// +kubebuilder:validation:Enum=TCP;Enabled;Disabled
type BPFConnectTimeLBType string

const (
	BPFConnectTimeLBTCP      BPFConnectTimeLBType = "TCP"
	BPFConnectTimeLBEnabled  BPFConnectTimeLBType = "Enabled"
	BPFConnectTimeLBDisabled BPFConnectTimeLBType = "Disabled"
)

// +kubebuilder:validation:Enum=Auto;Userspace;BPFProgram
type BPFConntrackMode string

const (
	BPFConntrackModeAuto       BPFConntrackMode = "Auto"
	BPFConntrackModeUserspace  BPFConntrackMode = "Userspace"
	BPFConntrackModeBPFProgram BPFConntrackMode = "BPFProgram"
)

// +kubebuilder:validation:Enum=Enabled;Disabled
type WindowsManageFirewallRulesMode string

const (
	WindowsManageFirewallRulesEnabled  WindowsManageFirewallRulesMode = "Enabled"
	WindowsManageFirewallRulesDisabled WindowsManageFirewallRulesMode = "Disabled"
)

// FelixConfigurationSpec contains the values of the Felix configuration.
type FelixConfigurationSpec struct {
	// UseInternalDataplaneDriver, if true, Felix will use its internal dataplane programming logic.  If false, it
	// will launch an external dataplane driver and communicate with it over protobuf.
	UseInternalDataplaneDriver *bool `json:"useInternalDataplaneDriver,omitempty"`

	// DataplaneDriver filename of the external dataplane driver to use.  Only used if UseInternalDataplaneDriver
	// is set to false.
	DataplaneDriver string `json:"dataplaneDriver,omitempty"`

	// DataplaneWatchdogTimeout is the readiness/liveness timeout used for Felix's (internal) dataplane driver.
	// Deprecated: replaced by the generic HealthTimeoutOverrides.
	DataplaneWatchdogTimeout *metav1.Duration `json:"dataplaneWatchdogTimeout,omitempty" configv1timescale:"seconds"`

	// IPv6Support controls whether Felix enables support for IPv6 (if supported by the in-use dataplane).
	IPv6Support *bool `json:"ipv6Support,omitempty" confignamev1:"Ipv6Support"`

	// RouteRefreshInterval is the period at which Felix re-checks the routes
	// in the dataplane to ensure that no other process has accidentally broken Calico's rules.
	// Set to 0 to disable route refresh. [Default: 90s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	RouteRefreshInterval *metav1.Duration `json:"routeRefreshInterval,omitempty" configv1timescale:"seconds"`

	// InterfaceRefreshInterval is the period at which Felix rescans local interfaces to verify their state.
	// The rescan can be disabled by setting the interval to 0.
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	InterfaceRefreshInterval *metav1.Duration `json:"interfaceRefreshInterval,omitempty" configv1timescale:"seconds"`

	// IptablesRefreshInterval is the period at which Felix re-checks the IP sets
	// in the dataplane to ensure that no other process has accidentally broken Calico's rules.
	// Set to 0 to disable IP sets refresh. Note: the default for this value is lower than the
	// other refresh intervals as a workaround for a Linux kernel bug that was fixed in kernel
	// version 4.11. If you are using v4.11 or greater you may want to set this to, a higher value
	// to reduce Felix CPU usage. [Default: 10s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	IptablesRefreshInterval *metav1.Duration `json:"iptablesRefreshInterval,omitempty" configv1timescale:"seconds"`

	// IptablesPostWriteCheckInterval is the period after Felix has done a write
	// to the dataplane that it schedules an extra read back in order to check the write was not
	// clobbered by another process. This should only occur if another application on the system
	// doesn't respect the iptables lock. [Default: 1s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	IptablesPostWriteCheckInterval *metav1.Duration `json:"iptablesPostWriteCheckInterval,omitempty" configv1timescale:"seconds" confignamev1:"IptablesPostWriteCheckIntervalSecs"`

	// IptablesLockFilePath is the location of the iptables lock file. You may need to change this
	// if the lock file is not in its standard location (for example if you have mapped it into Felix's
	// container at a different path). [Default: /run/xtables.lock]
	IptablesLockFilePath string `json:"iptablesLockFilePath,omitempty"`

	// IptablesLockTimeout is the time that Felix itself will wait for the iptables lock (rather than delegating the
	// lock handling to the `iptables` command).
	//
	// Deprecated: `iptables-restore` v1.8+ always takes the lock, so enabling this feature results in deadlock.
	// [Default: 0s disabled]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	IptablesLockTimeout *metav1.Duration `json:"iptablesLockTimeout,omitempty" configv1timescale:"seconds" confignamev1:"IptablesLockTimeoutSecs"`

	// IptablesLockProbeInterval when IptablesLockTimeout is enabled: the time that Felix will wait between
	// attempts to acquire the iptables lock if it is not available. Lower values make Felix more
	// responsive when the lock is contended, but use more CPU. [Default: 50ms]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	IptablesLockProbeInterval *metav1.Duration `json:"iptablesLockProbeInterval,omitempty" configv1timescale:"milliseconds" confignamev1:"IptablesLockProbeIntervalMillis"`

	// FeatureDetectOverride is used to override feature detection based on auto-detected platform
	// capabilities.  Values are specified in a comma separated list with no spaces, example;
	// "SNATFullyRandom=true,MASQFullyRandom=false,RestoreSupportsLock=". A value of "true" or "false" will
	// force enable/disable feature, empty or omitted values fall back to auto-detection.
	// +kubebuilder:validation:Pattern=`^([a-zA-Z0-9-_]+=(true|false|),)*([a-zA-Z0-9-_]+=(true|false|))?$`
	FeatureDetectOverride string `json:"featureDetectOverride,omitempty" validate:"omitempty,keyValueList"`

	// FeatureGates is used to enable or disable tech-preview Calico features.
	// Values are specified in a comma separated list with no spaces, example;
	// "BPFConnectTimeLoadBalancingWorkaround=enabled,XyZ=false". This is
	// used to enable features that are not fully production ready.
	// +kubebuilder:validation:Pattern=`^([a-zA-Z0-9-_]+=([^=]+),)*([a-zA-Z0-9-_]+=([^=]+))?$`
	FeatureGates string `json:"featureGates,omitempty" validate:"omitempty,keyValueList"`

	// IpsetsRefreshInterval controls the period at which Felix re-checks all IP sets to look for discrepancies.
	// Set to 0 to disable the periodic refresh. [Default: 90s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	IpsetsRefreshInterval *metav1.Duration `json:"ipsetsRefreshInterval,omitempty" configv1timescale:"seconds"`

	// MaxIpsetSize is the maximum number of IP addresses that can be stored in an IP set. Not applicable
	// if using the nftables backend.
	MaxIpsetSize *int `json:"maxIpsetSize,omitempty"`

	// IptablesBackend controls which backend of iptables will be used. The default is `Auto`.
	//
	// Warning: changing this on a running system can leave "orphaned" rules in the "other" backend. These
	// should be cleaned up to avoid confusing interactions.
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^(?i)(Auto|Legacy|NFT)?$`
	IptablesBackend *IptablesBackend `json:"iptablesBackend,omitempty" validate:"omitempty,iptablesBackend"`

	// XDPRefreshInterval is the period at which Felix re-checks all XDP state to ensure that no
	// other process has accidentally broken Calico's BPF maps or attached programs. Set to 0 to
	// disable XDP refresh. [Default: 90s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	XDPRefreshInterval *metav1.Duration `json:"xdpRefreshInterval,omitempty" configv1timescale:"seconds"`

	// NetlinkTimeout is the timeout when talking to the kernel over the netlink protocol, used for programming
	// routes, rules, and other kernel objects. [Default: 10s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	NetlinkTimeout *metav1.Duration `json:"netlinkTimeout,omitempty" configv1timescale:"seconds" confignamev1:"NetlinkTimeoutSecs"`

	// MetadataAddr is the IP address or domain name of the server that can answer VM queries for
	// cloud-init metadata. In OpenStack, this corresponds to the machine running nova-api (or in
	// Ubuntu, nova-api-metadata). A value of none (case-insensitive) means that Felix should not
	// set up any NAT rule for the metadata path. [Default: 127.0.0.1]
	MetadataAddr string `json:"metadataAddr,omitempty"`

	// MetadataPort is the port of the metadata server. This, combined with global.MetadataAddr (if
	// not 'None'), is used to set up a NAT rule, from 169.254.169.254:80 to MetadataAddr:MetadataPort.
	// In most cases this should not need to be changed [Default: 8775].
	MetadataPort *int `json:"metadataPort,omitempty"`

	// OpenstackRegion is the name of the region that a particular Felix belongs to. In a multi-region
	// Calico/OpenStack deployment, this must be configured somehow for each Felix (here in the datamodel,
	// or in felix.cfg or the environment on each compute node), and must match the [calico]
	// openstack_region value configured in neutron.conf on each node. [Default: Empty]
	OpenstackRegion string `json:"openstackRegion,omitempty"`

	// InterfacePrefix is the interface name prefix that identifies workload endpoints and so distinguishes
	// them from host endpoint interfaces. Note: in environments other than bare metal, the orchestrators
	// configure this appropriately. For example our Kubernetes and Docker integrations set the 'cali' value,
	// and our OpenStack integration sets the 'tap' value. [Default: cali]
	InterfacePrefix string `json:"interfacePrefix,omitempty"`

	// InterfaceExclude A comma-separated list of interface names that should be excluded when Felix is resolving
	// host endpoints. The default value ensures that Felix ignores Kubernetes' internal `kube-ipvs0` device. If you
	// want to exclude multiple interface names using a single value, the list supports regular expressions. For
	// regular expressions you must wrap the value with `/`. For example having values `/^kube/,veth1` will exclude
	// all interfaces that begin with `kube` and also the interface `veth1`. [Default: kube-ipvs0]
	InterfaceExclude string `json:"interfaceExclude,omitempty"`

	// ChainInsertMode controls whether Felix hooks the kernel's top-level iptables chains by inserting a rule
	// at the top of the chain or by appending a rule at the bottom. insert is the safe default since it prevents
	// Calico's rules from being bypassed. If you switch to append mode, be sure that the other rules in the chains
	// signal acceptance by falling through to the Calico rules, otherwise the Calico policy will be bypassed.
	// [Default: insert]
	// +kubebuilder:validation:Pattern=`^(?i)(Insert|Append)?$`
	ChainInsertMode string `json:"chainInsertMode,omitempty"`

	// DefaultEndpointToHostAction controls what happens to traffic that goes from a workload endpoint to the host
	// itself (after the endpoint's egress policy is applied). By default, Calico blocks traffic from workload
	// endpoints to the host itself with an iptables "DROP" action. If you want to allow some or all traffic from
	// endpoint to host, set this parameter to RETURN or ACCEPT. Use RETURN if you have your own rules in the iptables
	// "INPUT" chain; Calico will insert its rules at the top of that chain, then "RETURN" packets to the "INPUT" chain
	// once it has completed processing workload endpoint egress policy. Use ACCEPT to unconditionally accept packets
	// from workloads after processing workload endpoint egress policy. [Default: Drop]
	// +kubebuilder:validation:Pattern=`^(?i)(Drop|Accept|Return)?$`
	DefaultEndpointToHostAction string `json:"defaultEndpointToHostAction,omitempty" validate:"omitempty,dropAcceptReturn"`

	// IptablesFilterAllowAction controls what happens to traffic that is accepted by a Felix policy chain in the
	// iptables filter table (which is used for "normal" policy). The default will immediately `Accept` the traffic. Use
	// `Return` to send the traffic back up to the system chains for further processing.
	// +kubebuilder:validation:Pattern=`^(?i)(Accept|Return)?$`
	IptablesFilterAllowAction string `json:"iptablesFilterAllowAction,omitempty" validate:"omitempty,acceptReturn"`

	// IptablesMangleAllowAction controls what happens to traffic that is accepted by a Felix policy chain in the
	// iptables mangle table (which is used for "pre-DNAT" policy). The default will immediately `Accept` the traffic.
	// Use `Return` to send the traffic back up to the system chains for further processing.
	// +kubebuilder:validation:Pattern=`^(?i)(Accept|Return)?$`
	IptablesMangleAllowAction string `json:"iptablesMangleAllowAction,omitempty" validate:"omitempty,acceptReturn"`

	// IptablesFilterDenyAction controls what happens to traffic that is denied by network policy. By default Calico blocks traffic
	// with an iptables "DROP" action. If you want to use "REJECT" action instead you can configure it in here.
	// +kubebuilder:validation:Pattern=`^(?i)(Drop|Reject)?$`
	IptablesFilterDenyAction string `json:"iptablesFilterDenyAction,omitempty" validate:"omitempty,dropReject"`

	// LogPrefix is the log prefix that Felix uses when rendering LOG rules. [Default: calico-packet]
	LogPrefix string `json:"logPrefix,omitempty"`

	// LogDropActionOverride specifies whether or not to include the DropActionOverride in the logs when it is triggered.
	LogDropActionOverride *bool `json:"logDropActionOverride,omitempty"`

	// LogFilePath is the full path to the Felix log. Set to none to disable file logging. [Default: /var/log/calico/felix.log]
	LogFilePath string `json:"logFilePath,omitempty"`

	// LogSeverityFile is the log severity above which logs are sent to the log file. [Default: Info]
	// +kubebuilder:validation:Pattern=`^(?i)(Debug|Info|Warning|Error|Fatal)?$`
	LogSeverityFile string `json:"logSeverityFile,omitempty" validate:"omitempty,logLevel"`

	// LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: Info]
	// +kubebuilder:validation:Pattern=`^(?i)(Debug|Info|Warning|Error|Fatal)?$`
	LogSeverityScreen string `json:"logSeverityScreen,omitempty" validate:"omitempty,logLevel"`

	// LogSeveritySys is the log severity above which logs are sent to the syslog. Set to None for no logging to syslog.
	// [Default: Info]
	// +kubebuilder:validation:Pattern=`^(?i)(Debug|Info|Warning|Error|Fatal)?$`
	LogSeveritySys string `json:"logSeveritySys,omitempty" validate:"omitempty,logLevel"`

	// LogDebugFilenameRegex controls which source code files have their Debug log output included in the logs.
	// Only logs from files with names that match the given regular expression are included.  The filter only applies
	// to Debug level logs.
	LogDebugFilenameRegex string `json:"logDebugFilenameRegex,omitempty" validate:"omitempty,regexp"`

	// IPIPEnabled overrides whether Felix should configure an IPIP interface on the host. Optional as Felix
	// determines this based on the existing IP pools. [Default: nil (unset)]
	IPIPEnabled *bool `json:"ipipEnabled,omitempty" confignamev1:"IpInIpEnabled"`

	// IPIPMTU controls the MTU to set on the IPIP tunnel device.  Optional as Felix auto-detects the MTU based on the
	// MTU of the host's interfaces. [Default: 0 (auto-detect)]
	IPIPMTU *int `json:"ipipMTU,omitempty" confignamev1:"IpInIpMtu"`

	// VXLANEnabled overrides whether Felix should create the VXLAN tunnel device for IPv4 VXLAN networking.
	// Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]
	VXLANEnabled *bool `json:"vxlanEnabled,omitempty" confignamev1:"VXLANEnabled"`

	// VXLANMTU is the MTU to set on the IPv4 VXLAN tunnel device.  Optional as Felix auto-detects the MTU based on the
	// MTU of the host's interfaces. [Default: 0 (auto-detect)]
	VXLANMTU *int `json:"vxlanMTU,omitempty"`

	// VXLANMTUV6 is the MTU to set on the IPv6 VXLAN tunnel device. Optional as Felix auto-detects the MTU based on the
	// MTU of the host's interfaces. [Default: 0 (auto-detect)]
	VXLANMTUV6 *int `json:"vxlanMTUV6,omitempty"`

	// VXLANPort is the UDP port number to use for VXLAN traffic. [Default: 4789]
	VXLANPort *int `json:"vxlanPort,omitempty"`

	// VXLANVNI is the VXLAN VNI to use for VXLAN traffic.  You may need to change this if the default value is
	// in use on your system. [Default: 4096]
	VXLANVNI *int `json:"vxlanVNI,omitempty"`

	// AllowVXLANPacketsFromWorkloads controls whether Felix will add a rule to drop VXLAN encapsulated traffic
	// from workloads. [Default: false]
	// +optional
	AllowVXLANPacketsFromWorkloads *bool `json:"allowVXLANPacketsFromWorkloads,omitempty"`

	// AllowIPIPPacketsFromWorkloads controls whether Felix will add a rule to drop IPIP encapsulated traffic
	// from workloads. [Default: false]
	// +optional
	AllowIPIPPacketsFromWorkloads *bool `json:"allowIPIPPacketsFromWorkloads,omitempty"`

	// ReportingInterval is the interval at which Felix reports its status into the datastore or 0 to disable.
	// Must be non-zero in OpenStack deployments. [Default: 30s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	ReportingInterval *metav1.Duration `json:"reportingInterval,omitempty" configv1timescale:"seconds" confignamev1:"ReportingIntervalSecs"`

	// ReportingTTL is the time-to-live setting for process-wide status reports. [Default: 90s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	ReportingTTL *metav1.Duration `json:"reportingTTL,omitempty" configv1timescale:"seconds" confignamev1:"ReportingTTLSecs"`

	// EndpointReportingEnabled controls whether Felix reports endpoint status to the datastore. This is only used
	// by the OpenStack integration. [Default: false]
	EndpointReportingEnabled *bool `json:"endpointReportingEnabled,omitempty"`

	// EndpointReportingDelay is the delay before Felix reports endpoint status to the datastore. This is only used
	// by the OpenStack integration. [Default: 1s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	EndpointReportingDelay *metav1.Duration `json:"endpointReportingDelay,omitempty" configv1timescale:"seconds" confignamev1:"EndpointReportingDelaySecs"`

	// EndpointStatusPathPrefix is the path to the directory where endpoint status will be written. Endpoint status
	// file reporting is disabled if field is left empty.
	//
	// Chosen directory should match the directory used by the CNI plugin for PodStartupDelay.
	// [Default: ""]
	EndpointStatusPathPrefix string `json:"endpointStatusPathPrefix,omitempty"`

	// IptablesMarkMask is the mask that Felix selects its IPTables Mark bits from. Should be a 32 bit hexadecimal
	// number with at least 8 bits set, none of which clash with any other mark bits in use on the system.
	// [Default: 0xffff0000]
	IptablesMarkMask *uint32 `json:"iptablesMarkMask,omitempty"`

	// DisableConntrackInvalidCheck disables the check for invalid connections in conntrack. While the conntrack
	// invalid check helps to detect malicious traffic, it can also cause issues with certain multi-NIC scenarios.
	DisableConntrackInvalidCheck *bool `json:"disableConntrackInvalidCheck,omitempty"`

	// HealthEnabled if set to true, enables Felix's health port, which provides readiness and liveness endpoints.
	// [Default: false]
	HealthEnabled *bool `json:"healthEnabled,omitempty"`
	// HealthHost is the host that the health server should bind to. [Default: localhost]
	HealthHost *string `json:"healthHost,omitempty"`
	// HealthPort is the TCP port that the health server should bind to. [Default: 9099]
	HealthPort *int `json:"healthPort,omitempty"`

	// HealthTimeoutOverrides allows the internal watchdog timeouts of individual subcomponents to be
	// overridden.  This is useful for working around "false positive" liveness timeouts that can occur
	// in particularly stressful workloads or if CPU is constrained.  For a list of active
	// subcomponents, see Felix's logs.
	HealthTimeoutOverrides []HealthTimeoutOverride `json:"healthTimeoutOverrides,omitempty" validate:"omitempty,dive"`

	// PrometheusMetricsEnabled enables the Prometheus metrics server in Felix if set to true. [Default: false]
	PrometheusMetricsEnabled *bool `json:"prometheusMetricsEnabled,omitempty"`

	// PrometheusMetricsHost is the host that the Prometheus metrics server should bind to. [Default: empty]
	PrometheusMetricsHost string `json:"prometheusMetricsHost,omitempty" validate:"omitempty,prometheusHost"`

	// PrometheusMetricsPort is the TCP port that the Prometheus metrics server should bind to. [Default: 9091]
	PrometheusMetricsPort *int `json:"prometheusMetricsPort,omitempty"`

	// PrometheusGoMetricsEnabled disables Go runtime metrics collection, which the Prometheus client does by default, when
	// set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]
	PrometheusGoMetricsEnabled *bool `json:"prometheusGoMetricsEnabled,omitempty"`

	// PrometheusProcessMetricsEnabled disables process metrics collection, which the Prometheus client does by default, when
	// set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]
	PrometheusProcessMetricsEnabled *bool `json:"prometheusProcessMetricsEnabled,omitempty"`

	// PrometheusWireGuardMetricsEnabled disables wireguard metrics collection, which the Prometheus client does by default, when
	// set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]
	PrometheusWireGuardMetricsEnabled *bool `json:"prometheusWireGuardMetricsEnabled,omitempty"`

	// PrometheusMetricsCertFile is the path to the TLS certificate file for the Prometheus metrics server. [Default: empty]
	PrometheusMetricsCertFile string `json:"prometheusMetricsCertFile,omitempty"`
	// PrometheusMetricsKeyFile is the path to the TLS private key file for the Prometheus metrics server. [Default: empty]
	PrometheusMetricsKeyFile string `json:"prometheusMetricsKeyFile,omitempty"`
	// PrometheusMetricsCAFile is the path to the TLS CA file for the Prometheus metrics server. [Default: empty]
	PrometheusMetricsCAFile string `json:"prometheusMetricsCAFile,omitempty"`

	// FailsafeInboundHostPorts is a list of ProtoPort struct objects including UDP/TCP/SCTP ports and CIDRs that Felix will
	// allow incoming traffic to host endpoints on irrespective of the security policy. This is useful to avoid accidentally
	// cutting off a host with incorrect configuration. For backwards compatibility, if the protocol is not specified,
	// it defaults to "tcp". If a CIDR is not specified, it will allow traffic from all addresses. To disable all inbound host ports,
	// use the value "[]". The default value allows ssh access, DHCP, BGP, etcd and the Kubernetes API.
	// [Default: tcp:22, udp:68, tcp:179, tcp:2379, tcp:2380, tcp:5473, tcp:6443, tcp:6666, tcp:6667 ]
	FailsafeInboundHostPorts *[]ProtoPort `json:"failsafeInboundHostPorts,omitempty"`

	// FailsafeOutboundHostPorts is a list of PortProto struct objects including UDP/TCP/SCTP ports and CIDRs that Felix
	// will allow outgoing traffic from host endpoints to irrespective of the security policy. This is useful to avoid accidentally
	// cutting off a host with incorrect configuration. For backwards compatibility, if the protocol is not specified, it defaults
	// to "tcp". If a CIDR is not specified, it will allow traffic from all addresses. To disable all outbound host ports,
	// use the value "[]". The default value opens etcd's standard ports to ensure that Felix does not get cut off from etcd
	// as well as allowing DHCP, DNS, BGP and the Kubernetes API.
	// [Default: udp:53, udp:67, tcp:179, tcp:2379, tcp:2380, tcp:5473, tcp:6443, tcp:6666, tcp:6667 ]
	FailsafeOutboundHostPorts *[]ProtoPort `json:"failsafeOutboundHostPorts,omitempty"`

	// KubeMasqueradeBit should be set to the same value as --iptables-masquerade-bit of kube-proxy
	// when TPROXY is used. The default is the same as kube-proxy default thus only needs a change
	// if kube-proxy is using a non-standard setting. Must be within the range of 0-31.  [Default: 14]
	KubeMasqueradeBit *int `json:"kubeMasqueradeBit,omitempty" validate:"omitempty,gte=0,lte=31"`

	// KubeNodePortRanges holds list of port ranges used for service node ports. Only used if felix detects kube-proxy running in ipvs mode.
	// Felix uses these ranges to separate host and workload traffic. [Default: 30000:32767].
	KubeNodePortRanges *[]numorstring.Port `json:"kubeNodePortRanges,omitempty" validate:"omitempty,dive"`

	// PolicySyncPathPrefix is used to by Felix to communicate policy changes to external services,
	// like Application layer policy. [Default: Empty]
	PolicySyncPathPrefix string `json:"policySyncPathPrefix,omitempty"`

	// UsageReportingEnabled is unused in Calico Enterprise, usage reporting is permanently disabled.
	UsageReportingEnabled *bool `json:"usageReportingEnabled,omitempty"`

	// UsageReportingInitialDelay is unused in Calico Enterprise, usage reporting is permanently disabled. [Default: 300s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	UsageReportingInitialDelay *metav1.Duration `json:"usageReportingInitialDelay,omitempty" configv1timescale:"seconds" confignamev1:"UsageReportingInitialDelaySecs"`

	// UsageReportingInterval is unused in Calico Enterprise, usage reporting is permanently disabled. [Default: 86400s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	UsageReportingInterval *metav1.Duration `json:"usageReportingInterval,omitempty" configv1timescale:"seconds" confignamev1:"UsageReportingIntervalSecs"`

	// NATPortRange specifies the range of ports that is used for port mapping when doing outgoing NAT. When unset the default behavior of the
	// network stack is used.
	NATPortRange *numorstring.Port `json:"natPortRange,omitempty"`

	// NATOutgoingAddress specifies an address to use when performing source NAT for traffic in a natOutgoing pool that
	// is leaving the network. By default the address used is an address on the interface the traffic is leaving on
	// (i.e. it uses the iptables MASQUERADE target).
	NATOutgoingAddress string `json:"natOutgoingAddress,omitempty"`

	// DeviceRouteSourceAddress IPv4 address to set as the source hint for routes programmed by Felix. When not set
	// the source address for local traffic from host to workload will be determined by the kernel.
	DeviceRouteSourceAddress string `json:"deviceRouteSourceAddress,omitempty"`

	// DeviceRouteSourceAddressIPv6 IPv6 address to set as the source hint for routes programmed by Felix. When not set
	// the source address for local traffic from host to workload will be determined by the kernel.
	DeviceRouteSourceAddressIPv6 string `json:"deviceRouteSourceAddressIPv6,omitempty"`

	// DeviceRouteProtocol controls the protocol to set on routes programmed by Felix. The protocol is an 8-bit label
	// used to identify the owner of the route.
	DeviceRouteProtocol *int `json:"deviceRouteProtocol,omitempty"`

	// RemoveExternalRoutes Controls whether Felix will remove unexpected routes to workload interfaces. Felix will
	// always clean up expected routes that use the configured DeviceRouteProtocol.  To add your own routes, you must
	// use a distinct protocol (in addition to setting this field to false).
	RemoveExternalRoutes *bool `json:"removeExternalRoutes,omitempty"`

	// IPForwarding controls whether Felix sets the host sysctls to enable IP forwarding.  IP forwarding is required
	// when using Calico for workload networking.  This should be disabled only on hosts where Calico is used solely for
	// host protection. In BPF mode, due to a kernel interaction, either IPForwarding must be enabled or BPFEnforceRPF
	// must be disabled. [Default: Enabled]
	// +kubebuilder:validation:Enum=Enabled;Disabled
	IPForwarding string `json:"ipForwarding,omitempty"`

	// ExternalNodesCIDRList is a list of CIDR's of external, non-Calico nodes from which VXLAN/IPIP overlay traffic
	// will be allowed.  By default, external tunneled traffic is blocked to reduce attack surface.
	ExternalNodesCIDRList *[]string `json:"externalNodesList,omitempty"`

	// NfNetlinkBufSize controls the size of NFLOG messages that the kernel will try to send to Felix.  NFLOG messages
	// are used to report flow verdicts from the kernel.  Warning: currently increasing the value may cause errors
	// due to a bug in the netlink library.
	NfNetlinkBufSize string `json:"nfNetlinkBufSize,omitempty"`

	// StatsDumpFilePath is the path to write a diagnostic flow logs statistics dump to when triggered by signal.
	StatsDumpFilePath string `json:"statsDumpFilePath,omitempty"`

	// PrometheusReporterEnabled controls whether the Prometheus per-flow metrics reporter is enabled. This is
	// used to show real-time flow metrics in the UI.
	PrometheusReporterEnabled *bool `json:"prometheusReporterEnabled,omitempty"`
	// PrometheusReporterPort is the port that the Prometheus per-flow metrics reporter should bind to.
	PrometheusReporterPort *int `json:"prometheusReporterPort,omitempty"`
	// PrometheusReporterCertFile is the path to the TLS certificate file for the Prometheus per-flow metrics reporter.
	PrometheusReporterCertFile string `json:"prometheusReporterCertFile,omitempty"`
	// PrometheusReporterKeyFile is the path to the TLS private key file for the Prometheus per-flow metrics reporter.
	PrometheusReporterKeyFile string `json:"prometheusReporterKeyFile,omitempty"`
	// PrometheusReporterCAFile is the path to the TLS CA file for the Prometheus per-flow metrics reporter.
	PrometheusReporterCAFile string `json:"prometheusReporterCAFile,omitempty"`
	// DeletedMetricsRetentionSecs controls how long metrics are retianed after the flow is gone.
	DeletedMetricsRetentionSecs *int `json:"deletedMetricsRetentionSecs,omitempty"`

	// DropActionOverride overrides the Drop action in Felix, optionally changing the behavior to Accept, and optionally adding Log.
	// Possible values are Drop, LogAndDrop, Accept, LogAndAccept. [Default: Drop]
	// +kubebuilder:validation:Pattern=`^(?i)(Drop|LogAndDrop|Accept|LogAndAccept)?$`
	DropActionOverride string `json:"dropActionOverride,omitempty" validate:"omitempty,dropActionOverride"`

	// DebugMemoryProfilePath is the path to write the memory profile to when triggered by signal.
	DebugMemoryProfilePath string `json:"debugMemoryProfilePath,omitempty"`

	// DebugDisableLogDropping disables the dropping of log messages when the log buffer is full.  This can
	// significantly impact performance if log write-out is a bottleneck. [Default: false]
	DebugDisableLogDropping *bool `json:"debugDisableLogDropping,omitempty"`

	// DebugSimulateCalcGraphHangAfter is used to simulate a hang in the calculation graph after the specified duration.
	// This is useful in tests of the watchdog system only!
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	DebugSimulateCalcGraphHangAfter *metav1.Duration `json:"debugSimulateCalcGraphHangAfter,omitempty" configv1timescale:"seconds"`

	// DebugSimulateDataplaneHangAfter is used to simulate a hang in the dataplane after the specified duration.
	// This is useful in tests of the watchdog system only!
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	DebugSimulateDataplaneHangAfter *metav1.Duration `json:"debugSimulateDataplaneHangAfter,omitempty" configv1timescale:"seconds"`

	// DebugSimulateDataplaneApplyDelay adds an artificial delay to every dataplane operation.  This is useful for
	// simulating a heavily loaded system for test purposes only.
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	DebugSimulateDataplaneApplyDelay *metav1.Duration `json:"debugSimulateDataplaneApplyDelay,omitempty" configv1timescale:"seconds"`

	// DebugHost is the host IP or hostname to bind the debug port to.  Only used
	// if DebugPort is set. [Default:localhost]
	DebugHost *string `json:"debugHost,omitempty"`

	// DebugPort if set, enables Felix's debug HTTP port, which allows memory and CPU profiles
	// to be retrieved.  The debug port is not secure, it should not be exposed to the internet.
	DebugPort *int `json:"debugPort,omitempty" validate:"omitempty,gte=0,lte=65535"`

	// This parameter can be used to limit the host interfaces on which Calico will apply SNAT to traffic leaving a
	// Calico IPAM pool with "NAT outgoing" enabled. This can be useful if you have a main data interface, where
	// traffic should be SNATted and a secondary device (such as the docker bridge) which is local to the host and
	// doesn't require SNAT. This parameter uses the iptables interface matching syntax, which allows + as a
	// wildcard. Most users will not need to set this. Example: if your data interfaces are eth0 and eth1 and you
	// want to exclude the docker bridge, you could set this to eth+
	IptablesNATOutgoingInterfaceFilter string `json:"iptablesNATOutgoingInterfaceFilter,omitempty" validate:"omitempty,ifaceFilter"`

	// SidecarAccelerationEnabled enables experimental sidecar acceleration [Default: false]
	SidecarAccelerationEnabled *bool `json:"sidecarAccelerationEnabled,omitempty"`

	// XDPEnabled enables XDP acceleration for suitable untracked incoming deny rules. [Default: true]
	XDPEnabled *bool `json:"xdpEnabled,omitempty" confignamev1:"XDPEnabled"`

	// GenericXDPEnabled enables Generic XDP so network cards that don't support XDP offload or driver
	// modes can use XDP. This is not recommended since it doesn't provide better performance than
	// iptables. [Default: false]
	GenericXDPEnabled *bool `json:"genericXDPEnabled,omitempty" confignamev1:"GenericXDPEnabled"`

	// NFTablesMode configures nftables support in Felix. [Default: Disabled]
	// +kubebuilder:validation:Enum=Disabled;Enabled;Auto
	NFTablesMode *NFTablesMode `json:"nftablesMode,omitempty"`

	// NftablesRefreshInterval controls the interval at which Felix periodically refreshes the nftables rules. [Default: 90s]
	NftablesRefreshInterval *metav1.Duration `json:"nftablesRefreshInterval,omitempty" configv1timescale:"seconds"`

	// NftablesFilterAllowAction controls the nftables action that Felix uses to represent the "allow" policy verdict
	// in the filter table. The default is to `ACCEPT` the traffic, which is a terminal action.  Alternatively,
	// `RETURN` can be used to return the traffic back to the top-level chain for further processing by your rules.
	// +kubebuilder:validation:Pattern=`^(?i)(Accept|Return)?$`
	NftablesFilterAllowAction string `json:"nftablesFilterAllowAction,omitempty" validate:"omitempty,acceptReturn"`

	// NftablesMangleAllowAction controls the nftables action that Felix uses to represent the "allow" policy verdict
	// in the mangle table. The default is to `ACCEPT` the traffic, which is a terminal action.  Alternatively,
	// `RETURN` can be used to return the traffic back to the top-level chain for further processing by your rules.
	// +kubebuilder:validation:Pattern=`^(?i)(Accept|Return)?$`
	NftablesMangleAllowAction string `json:"nftablesMangleAllowAction,omitempty" validate:"omitempty,acceptReturn"`

	// NftablesFilterDenyAction controls what happens to traffic that is denied by network policy. By default, Calico
	// blocks traffic with a "drop" action. If you want to use a "reject" action instead you can configure it here.
	// +kubebuilder:validation:Pattern=`^(?i)(Drop|Reject)?$`
	NftablesFilterDenyAction string `json:"nftablesFilterDenyAction,omitempty" validate:"omitempty,dropReject"`

	// NftablesMarkMask is the mask that Felix selects its nftables Mark bits from. Should be a 32 bit hexadecimal
	// number with at least 8 bits set, none of which clash with any other mark bits in use on the system.
	// [Default: 0xffff0000]
	NftablesMarkMask *uint32 `json:"nftablesMarkMask,omitempty"`

	// BPFEnabled, if enabled Felix will use the BPF dataplane. [Default: false]
	BPFEnabled *bool `json:"bpfEnabled,omitempty" validate:"omitempty"`

	// BPFDisableUnprivileged, if enabled, Felix sets the kernel.unprivileged_bpf_disabled sysctl to disable
	// unprivileged use of BPF.  This ensures that unprivileged users cannot access Calico's BPF maps and
	// cannot insert their own BPF programs to interfere with Calico's. [Default: true]
	BPFDisableUnprivileged *bool `json:"bpfDisableUnprivileged,omitempty" validate:"omitempty"`

	// BPFLogLevel controls the log level of the BPF programs when in BPF dataplane mode.  One of "Off", "Info", or
	// "Debug".  The logs are emitted to the BPF trace pipe, accessible with the command `tc exec bpf debug`.
	// [Default: Off].
	// +optional
	// +kubebuilder:validation:Pattern=`^(?i)(Off|Info|Debug)?$`
	BPFLogLevel string `json:"bpfLogLevel,omitempty" validate:"omitempty,bpfLogLevel"`

	// BPFConntrackLogLevel controls the log level of the BPF conntrack cleanup program, which runs periodically
	// to clean up expired BPF conntrack entries.
	// [Default: Off].
	// +optional
	// +kubebuilder:validation:Enum=Off;Debug
	BPFConntrackLogLevel string `json:"bpfConntrackLogLevel,omitempty" validate:"omitempty,oneof=Off Debug"`

	// BPFConntrackCleanupMode controls how BPF conntrack entries are cleaned up.  `Auto` will use a BPF program if supported,
	// falling back to userspace if not.  `Userspace` will always use the userspace cleanup code.  `BPFProgram` will
	// always use the BPF program (failing if not supported).
	// [Default: Auto]
	BPFConntrackCleanupMode *BPFConntrackMode `json:"bpfConntrackMode,omitempty" validate:"omitempty,oneof=Auto Userspace BPFProgram"`

	// BPFLogFilters is a map of key=values where the value is
	// a pcap filter expression and the key is an interface name with 'all'
	// denoting all interfaces, 'weps' all workload endpoints and 'heps' all host
	// endpoints.
	//
	// When specified as an env var, it accepts a comma-separated list of
	// key=values.
	// [Default: unset - means all debug logs are emitted]
	// +optional
	BPFLogFilters *map[string]string `json:"bpfLogFilters,omitempty" validate:"omitempty,bpfLogFilters"`

	// BPFCTLBLogFilter specifies, what is logged by connect time load balancer when BPFLogLevel is
	// debug. Currently has to be specified as 'all' when BPFLogFilters is set
	// to see CTLB logs.
	// [Default: unset - means logs are emitted when BPFLogLevel id debug and BPFLogFilters not set.]
	// +optional
	BPFCTLBLogFilter string `json:"bpfCTLBLogFilter,omitempty" validate:"omitempty"`

	// BPFDataIfacePattern is a regular expression that controls which interfaces Felix should attach BPF programs to
	// in order to catch traffic to/from the network.  This needs to match the interfaces that Calico workload traffic
	// flows over as well as any interfaces that handle incoming traffic to nodeports and services from outside the
	// cluster.  It should not match the workload interfaces (usually named cali...).
	BPFDataIfacePattern string `json:"bpfDataIfacePattern,omitempty" validate:"omitempty,regexp"`

	// BPFL3IfacePattern is a regular expression that allows to list tunnel devices like wireguard or vxlan (i.e., L3 devices)
	// in addition to BPFDataIfacePattern. That is, tunnel interfaces not created by Calico, that Calico workload traffic flows
	// over as well as any interfaces that handle incoming traffic to nodeports and services from outside the cluster.
	BPFL3IfacePattern string `json:"bpfL3IfacePattern,omitempty" validate:"omitempty,regexp"`

	// BPFConnectTimeLoadBalancingEnabled when in BPF mode, controls whether Felix installs the connection-time load
	// balancer.  The connect-time load balancer is required for the host to be able to reach Kubernetes services
	// and it improves the performance of pod-to-service connections.  The only reason to disable it is for debugging
	// purposes.
	//
	// Deprecated: Use BPFConnectTimeLoadBalancing [Default: true]
	BPFConnectTimeLoadBalancingEnabled *bool `json:"bpfConnectTimeLoadBalancingEnabled,omitempty" validate:"omitempty"`

	// BPFConnectTimeLoadBalancing when in BPF mode, controls whether Felix installs the connect-time load
	// balancer. The connect-time load balancer is required for the host to be able to reach Kubernetes services
	// and it improves the performance of pod-to-service connections.When set to TCP, connect time load balancing
	// is available only for services with TCP ports. [Default: TCP]
	BPFConnectTimeLoadBalancing *BPFConnectTimeLBType `json:"bpfConnectTimeLoadBalancing,omitempty" validate:"omitempty,oneof=TCP Enabled Disabled"`

	// BPFHostNetworkedNATWithoutCTLB when in BPF mode, controls whether Felix does a NAT without CTLB. This along with BPFConnectTimeLoadBalancing
	// determines the CTLB behavior. [Default: Enabled]
	BPFHostNetworkedNATWithoutCTLB *BPFHostNetworkedNATType `json:"bpfHostNetworkedNATWithoutCTLB,omitempty" validate:"omitempty,oneof=Enabled Disabled"`

	// BPFExternalServiceMode in BPF mode, controls how connections from outside the cluster to services (node ports
	// and cluster IPs) are forwarded to remote workloads.  If set to "Tunnel" then both request and response traffic
	// is tunneled to the remote node.  If set to "DSR", the request traffic is tunneled but the response traffic
	// is sent directly from the remote node.  In "DSR" mode, the remote node appears to use the IP of the ingress
	// node; this requires a permissive L2 network.  [Default: Tunnel]
	// +kubebuilder:validation:Pattern=`^(?i)(Tunnel|DSR)?$`
	BPFExternalServiceMode string `json:"bpfExternalServiceMode,omitempty" validate:"omitempty,bpfServiceMode"`

	// BPFDSROptoutCIDRs is a list of CIDRs which are excluded from DSR. That is, clients
	// in those CIDRs will access service node ports as if BPFExternalServiceMode was set to
	// Tunnel.
	BPFDSROptoutCIDRs *[]string `json:"bpfDSROptoutCIDRs,omitempty" validate:"omitempty,cidrs"`

	// BPFExtToServiceConnmark in BPF mode, controls a 32bit mark that is set on connections from an
	// external client to a local service. This mark allows us to control how packets of that
	// connection are routed within the host and how is routing interpreted by RPF check. [Default: 0]
	BPFExtToServiceConnmark *int `json:"bpfExtToServiceConnmark,omitempty" validate:"omitempty,gte=0,lte=4294967295"`

	// BPFKubeProxyIptablesCleanupEnabled, if enabled in BPF mode, Felix will proactively clean up the upstream
	// Kubernetes kube-proxy's iptables chains.  Should only be enabled if kube-proxy is not running.  [Default: true]
	BPFKubeProxyIptablesCleanupEnabled *bool `json:"bpfKubeProxyIptablesCleanupEnabled,omitempty" validate:"omitempty"`

	// BPFKubeProxyMinSyncPeriod, in BPF mode, controls the minimum time between updates to the dataplane for Felix's
	// embedded kube-proxy.  Lower values give reduced set-up latency.  Higher values reduce Felix CPU usage by
	// batching up more work.  [Default: 1s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	BPFKubeProxyMinSyncPeriod *metav1.Duration `json:"bpfKubeProxyMinSyncPeriod,omitempty" validate:"omitempty" configv1timescale:"seconds"`

	// BPFKubeProxyEndpointSlicesEnabled is deprecated and has no effect. BPF
	// kube-proxy always accepts endpoint slices. This option will be removed in
	// the next release.
	BPFKubeProxyEndpointSlicesEnabled *bool `json:"bpfKubeProxyEndpointSlicesEnabled,omitempty" validate:"omitempty"`

	// BPFPSNATPorts sets the range from which we randomly pick a port if there is a source port
	// collision. This should be within the ephemeral range as defined by RFC 6056 (1024–65535) and
	// preferably outside the  ephemeral ranges used by common operating systems. Linux uses
	// 32768–60999, while others mostly use the IANA defined range 49152–65535. It is not necessarily
	// a problem if this range overlaps with the operating systems. Both ends of the range are
	// inclusive. [Default: 20000:29999]
	BPFPSNATPorts *numorstring.Port `json:"bpfPSNATPorts,omitempty"`

	// BPFMapSizeNATFrontend sets the size for NAT front end map.
	// FrontendMap should be large enough to hold an entry for each nodeport,
	// external IP and each port in each service.
	BPFMapSizeNATFrontend *int `json:"bpfMapSizeNATFrontend,omitempty"`

	// BPFMapSizeNATBackend sets the size for NAT back end map.
	// This is the total number of endpoints. This is mostly
	// more than the size of the number of services.
	BPFMapSizeNATBackend *int `json:"bpfMapSizeNATBackend,omitempty"`

	// BPFMapSizeNATAffinity sets the size of the BPF map that stores the affinity of a connection (for services that
	// enable that feature.
	BPFMapSizeNATAffinity *int `json:"bpfMapSizeNATAffinity,omitempty"`

	// BPFMapSizeRoute sets the size for the routes map.  The routes map should be large enough
	// to hold one entry per workload and a handful of entries per host (enough to cover its own IPs and
	// tunnel IPs).
	BPFMapSizeRoute *int `json:"bpfMapSizeRoute,omitempty"`

	// BPFMapSizeConntrack sets the size for the conntrack map.  This map must be large enough to hold
	// an entry for each active connection.  Warning: changing the size of the conntrack map can cause disruption.
	BPFMapSizeConntrack *int `json:"bpfMapSizeConntrack,omitempty"`

	// BPFMapSizeConntrackCleanupQueue sets the size for the map used to hold NAT conntrack entries that are queued
	// for cleanup.  This should be big enough to hold all the NAT entries that expire within one cleanup interval.
	// +kubebuilder:validation:Minimum=1
	BPFMapSizeConntrackCleanupQueue *int `json:"bpfMapSizeConntrackCleanupQueue,omitempty"  validate:"omitempty,gte=1"`

	// BPFMapSizeIPSets sets the size for ipsets map.  The IP sets map must be large enough to hold an entry
	// for each endpoint matched by every selector in the source/destination matches in network policy.  Selectors
	// such as "all()" can result in large numbers of entries (one entry per endpoint in that case).
	BPFMapSizeIPSets *int `json:"bpfMapSizeIPSets,omitempty"`

	// BPFMapSizeIfState sets the size for ifstate map.  The ifstate map must be large enough to hold an entry
	// for each device (host + workloads) on a host.
	BPFMapSizeIfState *int `json:"bpfMapSizeIfState,omitempty"`

	// BPFHostConntrackBypass Controls whether to bypass Linux conntrack in BPF mode for
	// workloads and services. [Default: true - bypass Linux conntrack]
	BPFHostConntrackBypass *bool `json:"bpfHostConntrackBypass,omitempty"`

	// BPFEnforceRPF enforce strict RPF on all host interfaces with BPF programs regardless of
	// what is the per-interfaces or global setting. Possible values are Disabled, Strict
	// or Loose. [Default: Loose]
	// +kubebuilder:validation:Pattern=`^(?i)(Disabled|Strict|Loose)?$`
	BPFEnforceRPF string `json:"bpfEnforceRPF,omitempty"`

	// BPFPolicyDebugEnabled when true, Felix records detailed information
	// about the BPF policy programs, which can be examined with the calico-bpf command-line tool.
	BPFPolicyDebugEnabled *bool `json:"bpfPolicyDebugEnabled,omitempty"`

	// BPFForceTrackPacketsFromIfaces in BPF mode, forces traffic from these interfaces
	// to skip Calico's iptables NOTRACK rule, allowing traffic from those interfaces to be
	// tracked by Linux conntrack.  Should only be used for interfaces that are not used for
	// the Calico fabric.  For example, a docker bridge device for non-Calico-networked
	// containers. [Default: docker+]
	BPFForceTrackPacketsFromIfaces *[]string `json:"bpfForceTrackPacketsFromIfaces,omitempty" validate:"omitempty,ifaceFilterSlice"`

	// BPFDisableGROForIfaces is a regular expression that controls which interfaces Felix should disable the
	// Generic Receive Offload [GRO] option.  It should not match the workload interfaces (usually named cali...).
	BPFDisableGROForIfaces string `json:"bpfDisableGROForIfaces,omitempty" validate:"omitempty,regexp"`

	// BPFExcludeCIDRsFromNAT is a list of CIDRs that are to be excluded from NAT
	// resolution so that host can handle them. A typical usecase is node local
	// DNS cache.
	BPFExcludeCIDRsFromNAT *[]string `json:"bpfExcludeCIDRsFromNAT,omitempty" validate:"omitempty,cidrs"`
	// BPFExportBufferSizeMB in BPF mode, controls the buffer size used for sending BPF events to felix.
	// [Default: 1]
	BPFExportBufferSizeMB *int `json:"bpfExportBufferSizeMB,omitempty" validate:"omitempty,cidrs"`

	// SyslogReporterEnabled turns on the feature to write logs to Syslog. Please note that this can incur significant
	// disk space usage when running felix on non-cluster hosts.
	SyslogReporterEnabled *bool `json:"syslogReporterEnabled,omitempty" validate:"omitempty"`
	// SyslogReporterNetwork is the network to dial to when writing to Syslog. Known networks are "tcp", "tcp4"
	// (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only), "ip6"
	// (IPv6-only), "unix", "unixgram" and "unixpacket". For more, see: https://pkg.go.dev/net#Dial
	SyslogReporterNetwork string `json:"syslogReporterNetwork,omitempty"`
	// SyslogReporterAddress is the address to dial to when writing to Syslog. For TCP and UDP networks, the address has
	// the form "host:port". The host must be a literal IP address, or a host name that can be resolved to IP addresses.
	// The port must be a literal port number or a service name. For more, see: https://pkg.go.dev/net#Dial
	SyslogReporterAddress string `json:"syslogReporterAddress,omitempty"`

	// IPSecMode controls which mode IPSec is operating on.
	// Default value means IPSec is not enabled. [Default: ""]
	IPSecMode string `json:"ipsecMode,omitempty" validate:"omitempty,ipsecMode"`
	// IPSecAllowUnsecuredTraffic controls whether non-IPsec traffic is allowed in addition to IPsec traffic. Enabling this
	// negates the anti-spoofing protections of IPsec but it is useful when migrating to/from IPsec. [Default: false]
	IPSecAllowUnsecuredTraffic *bool `json:"ipsecAllowUnsecuredTraffic,omitempty"`
	// IPSecIKEAlgorithm sets IPSec IKE algorithm. Default is NIST suite B recommendation. [Default: aes128gcm16-prfsha256-ecp256]
	IPSecIKEAlgorithm string `json:"ipsecIKEAlgorithm,omitempty"`
	// IPSecESAlgorithm sets IPSec ESP algorithm. Default is NIST suite B recommendation. [Default: aes128gcm16-ecp256]
	IPSecESPAlgorithm string `json:"ipsecESPAlgorithm,omitempty"`
	// IPSecLogLevel controls log level for IPSec components. Set to None for no logging.
	// A generic log level terminology is used [None, Notice, Info, Debug, Verbose].
	// [Default: Info]
	// +kubebuilder:validation:Pattern=`^(?i)(None|Notice|Info|Debug|Verbose)?$`
	IPSecLogLevel string `json:"ipsecLogLevel,omitempty" validate:"omitempty,ipsecLogLevel"`
	// IPSecPolicyRefreshInterval is the interval at which Felix will check the kernel's IPsec policy tables and
	// repair any inconsistencies. [Default: 600s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	IPSecPolicyRefreshInterval *metav1.Duration `json:"ipsecPolicyRefreshInterval,omitempty" configv1timescale:"seconds"`

	// FlowLogsFlushInterval configures the interval at which Felix exports flow logs.
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	FlowLogsFlushInterval *metav1.Duration `json:"flowLogsFlushInterval,omitempty" configv1timescale:"seconds"`
	// FlowLogsEnableHostEndpoint enables Flow logs reporting for HostEndpoints.
	FlowLogsEnableHostEndpoint *bool `json:"flowLogsEnableHostEndpoint,omitempty"`
	// FlowLogsEnableNetworkSets enables Flow logs reporting for GlobalNetworkSets.
	FlowLogsEnableNetworkSets *bool `json:"flowLogsEnableNetworkSets,omitempty"`
	// FlowLogsMaxOriginalIPsIncluded specifies the number of unique IP addresses (if relevant) that should be included in Flow logs.
	FlowLogsMaxOriginalIPsIncluded *int `json:"flowLogsMaxOriginalIPsIncluded,omitempty"`
	// FlowLogsCollectProcessInfo, if enabled Felix will load the kprobe BPF programs to collect process info. [Default: false]
	FlowLogsCollectProcessInfo *bool `json:"flowLogsCollectProcessInfo,omitempty" validate:"omitempty"`
	// FlowLogsCollectTcpStats enables flow logs reporting TCP socket stats
	FlowLogsCollectTcpStats *bool `json:"flowLogsCollectTcpStats,omitempty" validate:"omitempty"`
	// When FlowLogsCollectProcessPath and FlowLogsCollectProcessInfo are
	// both enabled, each flow log will include information about the process
	// that is sending or receiving the packets in that flow: the
	// `process_name` field will contain the full path of the process
	// executable, and the `process_args` field will have the arguments with
	// which the executable was invoked.  Process information will not be
	// reported for connections which use raw sockets.
	FlowLogsCollectProcessPath *bool `json:"flowLogsCollectProcessPath,omitempty" validate:"omitempty"`
	// When FlowLogsCollectorDebugTrace is set to true, enables the logs in the collector to be
	// printed in their entirety.
	FlowLogsCollectorDebugTrace *bool `json:"flowLogsCollectorDebugTrace,omitempty"`
	// FlowLogsDestDomainsByClient is used to configure if the source IP is used in the mapping of top
	// level destination domains. [Default: true]
	FlowLogsDestDomainsByClient *bool `json:"flowLogsDestDomainsByClient,omitempty"`
	// FlowLogsFileEnabled when set to true, enables logging flow logs to a file. If false no flow logging to file will occur.
	FlowLogsFileEnabled *bool `json:"flowLogsFileEnabled,omitempty"`
	// FlowLogsFileMaxFiles sets the number of log files to keep.
	FlowLogsFileMaxFiles *int `json:"flowLogsFileMaxFiles,omitempty"`
	// FlowLogsFileMaxFileSizeMB sets the max size in MB of flow logs files before rotation.
	FlowLogsFileMaxFileSizeMB *int `json:"flowLogsFileMaxFileSizeMB,omitempty"`
	// FlowLogsFileDirectory sets the directory where flow logs files are stored.
	FlowLogsFileDirectory *string `json:"flowLogsFileDirectory,omitempty"`
	// FlowLogsFileIncludeLabels is used to configure if endpoint labels are included in a Flow log entry written to file.
	FlowLogsFileIncludeLabels *bool `json:"flowLogsFileIncludeLabels,omitempty"`
	// FlowLogsFileIncludePolicies is used to configure if policy information are included in a Flow log entry written to file.
	FlowLogsFileIncludePolicies *bool `json:"flowLogsFileIncludePolicies,omitempty"`
	// FlowLogsFileIncludeService is used to configure if the destination service is included in a Flow log entry written to file.
	// The service information can only be included if the flow was explicitly determined to be directed at the service (e.g.
	// when the pre-DNAT destination corresponds to the service ClusterIP and port).
	FlowLogsFileIncludeService *bool `json:"flowLogsFileIncludeService,omitempty"`
	// FlowLogsFileAggregationKindForAllowed is used to choose the type of aggregation for flow log entries created for
	// allowed connections. [Default: 2 - pod prefix name based aggregation].
	// Accepted values are 0, 1 and 2.
	// 0 - No aggregation.
	// 1 - Source port based aggregation.
	// 2 - Pod prefix name based aggreagation.
	// +kubebuilder:validation:Enum=0;1;2
	FlowLogsFileAggregationKindForAllowed *int `json:"flowLogsFileAggregationKindForAllowed,omitempty" validate:"omitempty,flowLogAggregationKind"`
	// FlowLogsFileAggregationKindForDenied is used to choose the type of aggregation for flow log entries created for
	// denied connections. [Default: 1 - source port based aggregation].
	// Accepted values are 0, 1 and 2.
	// 0 - No aggregation.
	// 1 - Source port based aggregation.
	// 2 - Pod prefix name based aggregation.
	// 3 - No destination ports based aggregation.
	// +kubebuilder:validation:Enum=0;1;2;3
	FlowLogsFileAggregationKindForDenied *int `json:"flowLogsFileAggregationKindForDenied,omitempty" validate:"omitempty,flowLogAggregationKind"`
	// FlowLogsFileEnabledForAllowed is used to enable/disable flow logs entries created for allowed connections. Default is true.
	// This parameter only takes effect when FlowLogsFileReporterEnabled is set to true.
	FlowLogsFileEnabledForAllowed *bool `json:"flowLogsFileEnabledForAllowed,omitempty"`
	// FlowLogsFileEnabledForDenied is used to enable/disable flow logs entries created for denied flows. Default is true.
	// This parameter only takes effect when FlowLogsFileReporterEnabled is set to true.
	FlowLogsFileEnabledForDenied *bool `json:"flowLogsFileEnabledForDenied,omitempty"`
	// FlowLogsDynamicAggregationEnabled is used to enable/disable dynamically changing aggregation levels. Default is true.
	FlowLogsDynamicAggregationEnabled *bool `json:"flowLogsDynamicAggregationEnabled,omitempty"`
	// FlowLogsPositionFilePath is used specify the position of the external pipeline that reads flow logs. Default is /var/log/calico/flows.log.pos.
	// This parameter only takes effect when FlowLogsDynamicAggregationEnabled is set to true.
	FlowLogsPositionFilePath *string `json:"flowLogsPositionFilePath,omitempty"`
	// FlowLogsAggregationThresholdBytes is used specify how far behind the external pipeline that reads flow logs can be. Default is 8192 bytes.
	// This parameter only takes effect when FlowLogsDynamicAggregationEnabled is set to true.
	FlowLogsAggregationThresholdBytes *int `json:"flowLogsAggregationThresholdBytes,omitempty"`
	// FlowLogsFilePerFlowProcessLimit, is used to specify the maximum number of flow log entries with distinct process information
	// beyond which process information will be aggregated. [Default: 2]
	FlowLogsFilePerFlowProcessLimit *int `json:"flowLogsFilePerFlowProcessLimit,omitempty" validate:"omitempty"`
	// FlowLogsFilePerFlowProcessArgsLimit is used to specify the maximum number of distinct process args that will appear in the flowLogs.
	// Default value is 5
	FlowLogsFilePerFlowProcessArgsLimit *int `json:"flowLogsFilePerFlowProcessArgsLimit,omitempty" validate:"omitempty"`
	// FlowLogsFileNatOutgoingPortLimit is used to specify the maximum number of distinct post SNAT ports that will appear
	// in the flowLogs. Default value is 3
	FlowLogsFileNatOutgoingPortLimit *int `json:"flowLogsFileNatOutgoingPortLimit,omitempty" validate:"omitempty"`
	// FlowLogsFileDomainsLimit is used to configure the number of (destination) domains to include in the flow log.
	// These are not included for workload or host endpoint destinations.
	// [Default: 5]
	FlowLogsFileDomainsLimit *int `json:"flowLogsFileDomainsLimit,omitempty" validate:"omitempty"`

	// WindowsFlowLogsFileDirectory sets the directory where flow logs files are stored on Windows nodes. [Default: "c:\\TigeraCalico\\flowlogs"].
	WindowsFlowLogsFileDirectory string `json:"windowsFlowLogsFileDirectory,omitempty"`
	// WindowsFlowLogsPositionFilePath is used to specify the position of the external pipeline that reads flow logs on Windows nodes.
	// [Default: "c:\\TigeraCalico\\flowlogs\\flows.log.pos"].
	// This parameter only takes effect when FlowLogsDynamicAggregationEnabled is set to true.
	WindowsFlowLogsPositionFilePath string `json:"windowsFlowLogsPositionFilePath,omitempty"`
	// WindowsStatsDumpFilePath is used to specify the path of the stats dump file on Windows nodes. [Default: "c:\\TigeraCalico\\stats\\dump"]
	WindowsStatsDumpFilePath string `json:"windowsStatsDumpFilePath,omitempty"`
	// The name of the file that Felix uses to preserve learnt DNS information when restarting. [Default:
	// "c:\\TigeraCalico\\felix-dns-cache.txt"].
	WindowsDNSCacheFile string `json:"windowsDnsCacheFile,omitempty"`
	// Extra time to keep IPs and alias names that are learnt from DNS, in addition to each name
	// or IP's advertised TTL. The default value is 120s which is same as the default value of
	// ServicePointManager.DnsRefreshTimeout on .net framework. [Default: 120s].
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	WindowsDNSExtraTTL *metav1.Duration `json:"windowsDnsExtraTTL,omitempty" configv1timescale:"seconds"`

	// The DNS servers that Felix should trust. Each entry here must be `<ip>[:<port>]` - indicating an
	// explicit DNS server IP - or `k8s-service:[<namespace>/]<name>[:port]` - indicating a Kubernetes DNS
	// service. `<port>` defaults to the first service port, or 53 for an IP, and `<namespace>` to
	// `kube-system`. An IPv6 address with a port must use the square brackets convention, for example
	// `[fd00:83a6::12]:5353`.Note that Felix (calico-node) will need RBAC permission to read the details of
	// each service specified by a `k8s-service:...` form. [Default: "k8s-service:kube-dns"].
	DNSTrustedServers *[]string `json:"dnsTrustedServers,omitempty" validate:"omitempty,dive,ipOrK8sService"`
	// The name of the file that Felix uses to preserve learnt DNS information when restarting. [Default:
	// "/var/run/calico/felix-dns-cache.txt"].
	DNSCacheFile string `json:"dnsCacheFile,omitempty"`
	// The periodic interval at which Felix saves learnt DNS information to the cache file. [Default:
	// 60s].
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	DNSCacheSaveInterval *metav1.Duration `json:"dnsCacheSaveInterval,omitempty" configv1timescale:"seconds"`
	// An arbitrary number that can be changed, at runtime, to tell Felix to discard all its
	// learnt DNS information. [Default: 0].
	DNSCacheEpoch *int `json:"dnsCacheEpoch,omitempty"`
	// Extra time to keep IPs and alias names that are learnt from DNS, in addition to each name
	// or IP's advertised TTL. [Default: 0s].
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	DNSExtraTTL *metav1.Duration `json:"dnsExtraTTL,omitempty" configv1timescale:"seconds"`

	// DNSLogsFlushInterval configures the interval at which Felix exports DNS logs.
	// [Default: 300s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	DNSLogsFlushInterval *metav1.Duration `json:"dnsLogsFlushInterval,omitempty" configv1timescale:"seconds"`
	// DNSLogsFileEnabled controls logging DNS logs to a file. If false no DNS logging to file will occur.
	// [Default: false]
	DNSLogsFileEnabled *bool `json:"dnsLogsFileEnabled,omitempty"`
	// DNSLogsFileMaxFiles sets the number of DNS log files to keep.
	// [Default: 5]
	DNSLogsFileMaxFiles *int `json:"dnsLogsFileMaxFiles,omitempty"`
	// DNSLogsFileMaxFileSizeMB sets the max size in MB of DNS log files before rotation.
	// [Default: 100]
	DNSLogsFileMaxFileSizeMB *int `json:"dnsLogsFileMaxFileSizeMB,omitempty"`
	// DNSLogsFileDirectory sets the directory where DNS log files are stored.
	// [Default: /var/log/calico/dnslogs]
	DNSLogsFileDirectory *string `json:"dnsLogsFileDirectory,omitempty"`
	// DNSLogsFileIncludeLabels is used to configure if endpoint labels are included in a DNS log entry written to file.
	// [Default: true]
	DNSLogsFileIncludeLabels *bool `json:"dnsLogsFileIncludeLabels,omitempty"`
	// DNSLogsFileAggregationKind is used to choose the type of aggregation for DNS log entries.
	// [Default: 1 - client name prefix aggregation].
	// Accepted values are 0 and 1.
	// 0 - No aggregation.
	// 1 - Aggregate over clients with the same name prefix.
	// +kubebuilder:validation:Enum=0;1
	DNSLogsFileAggregationKind *int `json:"dnsLogsFileAggregationKind,omitempty" validate:"omitempty,dnsAggregationKind"`
	// Limit on the number of DNS logs that can be emitted within each flush interval.  When
	// this limit has been reached, Felix counts the number of unloggable DNS responses within
	// the flush interval, and emits a WARNING log with that count at the same time as it
	// flushes the buffered DNS logs.  [Default: 0, meaning no limit]
	DNSLogsFilePerNodeLimit *int `json:"dnsLogsFilePerNodeLimit,omitempty"`
	// DNSLogsLatency indicates to include measurements of DNS request/response latency in each DNS log.
	// [Default: true]
	DNSLogsLatency *bool `json:"dnsLogsLatency,omitempty"`
	// DNSPolicyMode specifies how DNS policy programming will be handled.
	// DelayDeniedPacket - Felix delays any denied packet that traversed a policy that included egress domain matches,
	// but did not match. The packet is released after a fixed time, or after the destination IP address was programmed.
	// DelayDNSResponse - Felix delays any DNS response until related IPSets are programmed. This introduces some
	// latency to all DNS packets (even when no IPSet programming is required), but it ensures policy hit statistics
	// are accurate. This is the recommended setting when you are making use of staged policies or policy rule hit
	// statistics.
	// NoDelay - Felix does not introduce any delay to the packets. DNS rules may not have been programmed by the time
	// the first packet traverses the policy rules. Client applications need to handle reconnection attempts if initial
	// connection attempts fail. This may be problematic for some applications or for very low DNS TTLs.
	//
	// This setting is ignored on Windows and "NoDelay" is always used.
	//
	// This setting is ignored by eBPF and BPFDNSPolicyMode is used instead.
	//
	// [Default: DelayDeniedPacket]
	DNSPolicyMode *DNSPolicyMode `json:"dnsPolicyMode,omitempty" validate:"omitempty,oneof=NoDelay DelayDeniedPacket DelayDNSResponse"`
	// BPFDNSPolicyMode specifies how DNS policy programming will be handled.
	// Inline - BPF parses DNS response inline with DNS response packet
	// processing. This guarantees the DNS rules reflect any change immediately.
	// NoDelay - Felix does not introduce any delay to the packets. DNS rules may not have been programmed by the time
	// the first packet traverses the policy rules. Client applications need to handle reconnection attempts if initial
	// connection attempts fail. This may be problematic for some applications or for very low DNS TTLs.
	// [Default: Inline]
	BPFDNSPolicyMode *BPFDNSPolicyMode `json:"bpfDNSPolicyMode,omitempty" validate:"omitempty,oneof=NoDelay Inline"`
	// DNSPolicyNfqueueID is the NFQUEUE ID to use for DNS Policy re-evaluation when the domains IP hasn't been programmed
	// to ipsets yet. Used when DNSPolicyMode is DelayDeniedPacket. [Default: 100]
	DNSPolicyNfqueueID *int `json:"dnsPolicyNfqueueID,omitempty" validate:"omitempty,gte=0,lte=65535"`
	// DNSPolicyNfqueueID is the size of the NFQUEUE for DNS policy re-evaluation. This is the maximum number of denied
	// packets that may be queued up pending re-evaluation.
	// Used when DNSPolicyMode is DelayDeniedPacket. [Default: 100]
	DNSPolicyNfqueueSize *int `json:"dnsPolicyNfqueueSize,omitempty" validate:"omitempty,gte=0,lte=65535"`
	// DNSPacketsNfqueueID is the NFQUEUE ID to use for capturing DNS packets to ensure programming IPSets occurs before
	// the response is released. Used when DNSPolicyMode is DelayDNSResponse. [Default: 101]
	DNSPacketsNfqueueID *int `json:"dnsPacketsNfqueueID,omitempty" validate:"omitempty,gte=0,lte=65535"`
	// DNSPacketsNfqueueSize is the size of the NFQUEUE for captured DNS packets. This is the maximum number of DNS
	// packets that may be queued awaiting programming in the dataplane. Used when DNSPolicyMode is DelayDNSResponse.
	// [Default: 100]
	DNSPacketsNfqueueSize *int `json:"dnsPacketsNfqueueSize,omitempty" validate:"omitempty,gte=0,lte=65535"`
	// DNSPacketsNfqueueMaxHoldDuration is the max length of time to hold on to a DNS response while waiting for the
	// the dataplane to be programmed. Used when DNSPolicyMode is DelayDNSResponse.
	// [Default: 3s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	DNSPacketsNfqueueMaxHoldDuration *metav1.Duration `json:"dnsPacketsNfqueueMaxHoldDuration,omitempty"`

	// L7LogsFlushInterval configures the interval at which Felix exports L7 logs.
	// [Default: 300s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	L7LogsFlushInterval *metav1.Duration `json:"l7LogsFlushInterval,omitempty" configv1timescale:"seconds"`
	// L7LogsFileEnabled controls logging L7 logs to a file. If false no L7 logging to file will occur.
	// [Default: true]
	L7LogsFileEnabled *bool `json:"l7LogsFileEnabled,omitempty"`
	// L7LogsFileMaxFiles sets the number of L7 log files to keep.
	// [Default: 5]
	L7LogsFileMaxFiles *int `json:"l7LogsFileMaxFiles,omitempty"`
	// L7LogsFileMaxFileSizeMB sets the max size in MB of L7 log files before rotation.
	// [Default: 100]
	L7LogsFileMaxFileSizeMB *int `json:"l7LogsFileMaxFileSizeMB,omitempty"`
	// L7LogsFileDirectory sets the directory where L7 log files are stored.
	// [Default: /var/log/calico/l7logs]
	L7LogsFileDirectory *string `json:"l7LogsFileDirectory,omitempty"`
	// L7LogsFileAggregationHTTPHeaderInfo is used to choose the type of aggregation for HTTP header data on L7 log entries.
	// [Default: ExcludeL7HTTPHeaderInfo - http header info removal].
	// Accepted values are IncludeL7HTTPHeaderInfo and ExcludeL7HTTPHeaderInfo.
	// IncludeL7HTTPHeaderInfo - Include HTTP header data in the logs.
	// ExcludeL7HTTPHeaderInfo - Aggregate over all other fields ignoring the user agent and log type.
	// +kubebuilder:validation:Pattern=`^(?i)(IncludeL7HTTPHeaderInfo|ExcludeL7HTTPHeaderInfo)?$`
	L7LogsFileAggregationHTTPHeaderInfo *string `json:"l7LogsFileAggregationHTTPHeaderInfo,omitempty" validate:"omitempty,l7HTTPHeaderAggregation"`
	// L7LogsFileAggregationHTTPMethod is used to choose the type of aggregation for the HTTP request method on L7 log entries.
	// [Default: IncludeL7HTTPMethod - include the HTTP method].
	// Accepted values are IncludeL7HTTPMethod and ExcludeL7HTTPMethod.
	// IncludeL7HTTPMethod - Include HTTP method in the logs.
	// ExcludeL7HTTPMethod - Aggregate over all other fields ignoring the HTTP method.
	// +kubebuilder:validation:Pattern=`^(?i)(IncludeL7HTTPMethod|ExcludeL7HTTPMethod)?$`
	L7LogsFileAggregationHTTPMethod *string `json:"l7LogsFileAggregationHTTPMethod,omitempty" validate:"omitempty,l7HTTPMethodAggregation"`
	// L7LogsFileAggregationServiceInfo is used to choose the type of aggregation for the service data on L7 log entries.
	// [Default: IncludeL7ServiceInfo - include service data].
	// Accepted values are IncludeL7ServiceInfo and ExcludeL7ServiceInfo.
	// IncludeL7ServiceInfo - Include service data in the logs.
	// ExcludeL7ServiceInfo - Aggregate over all other fields ignoring the service name, namespace, and port.
	// +kubebuilder:validation:Pattern=`^(?i)(IncludeL7ServiceInfo|ExcludeL7ServiceInfo)?$`
	L7LogsFileAggregationServiceInfo *string `json:"l7LogsFileAggregationServiceInfo,omitempty" validate:"omitempty,l7ServiceAggregation"`
	// L7LogsFileAggregationDestinationInfo is used to choose the type of aggregation for the destination metadata on L7 log entries.
	// [Default: IncludeL7DestinationInfo - include destination metadata].
	// Accepted values are IncludeL7DestinationInfo and ExcludeL7DestinationInfo.
	// IncludeL7DestinationInfo - Include destination metadata in the logs.
	// ExcludeL7DestinationInfo - Aggregate over all other fields ignoring the destination aggregated name, namespace, and type.
	// +kubebuilder:validation:Pattern=`^(?i)(IncludeL7DestinationInfo|ExcludeL7DestinationInfo)?$`
	L7LogsFileAggregationDestinationInfo *string `json:"l7LogsFileAggregationDestinationInfo,omitempty" validate:"omitempty,l7DestinationAggregation"`
	// L7LogsFileAggregationExcludeSourceInfo is used to choose the type of aggregation for the source metadata on L7 log entries.
	// [Default: IncludeL7SourceInfoNoPort - include all source metadata except for the source port].
	// Accepted values are IncludeL7SourceInfo, IncludeL7SourceInfoNoPort, and ExcludeL7SourceInfo.
	// IncludeL7SourceInfo - Include source metadata in the logs.
	// IncludeL7SourceInfoNoPort - Include source metadata in the logs excluding the source port.
	// ExcludeL7SourceInfo - Aggregate over all other fields ignoring the source aggregated name, namespace, and type.
	// +kubebuilder:validation:Pattern=`^(?i)(IncludeL7SourceInfo|IncludeL7SourceInfoNoPort|ExcludeL7SourceInfo)?$`
	L7LogsFileAggregationSourceInfo *string `json:"l7LogsFileAggregationSourceInfo,omitempty" validate:"omitempty,l7SourceAggregation"`
	// L7LogsFileAggregationResponseCode is used to choose the type of aggregation for the response code on L7 log entries.
	// [Default: IncludeL7ResponseCode - include the response code].
	// Accepted values are IncludeL7ResponseCode and ExcludeL7ResponseCode.
	// IncludeL7ResponseCode - Include the response code in the logs.
	// ExcludeL7ResponseCode - Aggregate over all other fields ignoring the response code.
	// +kubebuilder:validation:Pattern=`^(?i)(IncludeL7ResponseCode|ExcludeL7ResponseCode)?$`
	L7LogsFileAggregationResponseCode *string `json:"l7LogsFileAggregationResponseCode,omitempty" validate:"omitempty,l7ResponseCodeAggregation"`
	// L7LogsFileAggregationTrimURL is used to choose the type of aggregation for the url on L7 log entries.
	// [Default: IncludeL7FullURL - include the full URL up to however many path components are allowed by L7LogsFileAggregationNumURLPath].
	// Accepted values:
	// IncludeL7FullURL - Include the full URL up to however many path components are allowed by L7LogsFileAggregationNumURLPath.
	// TrimURLQuery - Aggregate over all other fields ignoring the query parameters on the URL.
	// TrimURLQueryAndPath - Aggregate over all other fields and the base URL only.
	// ExcludeL7URL - Aggregate over all other fields ignoring the URL entirely.
	// +kubebuilder:validation:Pattern=`^(?i)(IncludeL7FullURL|TrimURLQuery|TrimURLQueryAndPath|ExcludeL7URL)?$`
	L7LogsFileAggregationTrimURL *string `json:"l7LogsFileAggregationTrimURL,omitempty" validate:"omitempty,l7URLAggregation"`
	// L7LogsFileAggregationNumURLPath is used to choose the number of components in the url path to display.
	// This allows for the url to be truncated in case parts of the path provide no value. Setting this value
	// to negative will allow all parts of the path to be displayed.
	// [Default: 5].
	L7LogsFileAggregationNumURLPath *int `json:"l7LogsFileAggregationNumURLPath,omitempty"`
	// Limit on the length of the URL collected in L7 logs. When a URL length reaches this limit
	// it is sliced off, and the sliced URL is sent to log storage. [Default: 250]
	L7LogsFileAggregationURLCharLimit *int `json:"l7LogsFileAggregationURLCharLimit,omitempty"`
	// Limit on the number of L7 logs that can be emitted within each flush interval.  When
	// this limit has been reached, Felix counts the number of unloggable L7 responses within
	// the flush interval, and emits a WARNING log with that count at the same time as it
	// flushes the buffered L7 logs. A value of 0 means no limit. [Default: 1500]
	L7LogsFilePerNodeLimit *int `json:"l7LogsFilePerNodeLimit,omitempty"`

	// WindowsNetworkName specifies which Windows HNS networks Felix should operate on.  The default is to match
	// networks that start with "calico".  Supports regular expression syntax.
	WindowsNetworkName *string `json:"windowsNetworkName,omitempty"`

	// BPFRedirectToPeer controls which whether it is allowed to forward straight to the
	// peer side of the workload devices. It is allowed for any host L2 devices by default
	// (L2Only), but it breaks TCP dump on the host side of workload device as it bypasses
	// it on ingress. Value of Enabled also allows redirection from L3 host devices like
	// IPIP tunnel or Wireguard directly to the peer side of the workload's device. This
	// makes redirection faster, however, it breaks tools like tcpdump on the peer side.
	// Use Enabled with caution. [Default: Disabled]
	// +kubebuilder:validation:Enum=Enabled;Disabled;L2Only
	BPFRedirectToPeer string `json:"bpfRedirectToPeer,omitempty"`

	// RouteSource configures where Felix gets its routing information.
	// - WorkloadIPs: use workload endpoints to construct routes.
	// - CalicoIPAM: the default - use IPAM data to construct routes.
	// +kubebuilder:validation:Pattern=`^(?i)(WorkloadIPs|CalicoIPAM)?$`
	RouteSource string `json:"routeSource,omitempty" validate:"omitempty,routeSource"`

	// Calico programs additional Linux route tables for various purposes.
	// RouteTableRanges specifies a set of table index ranges that Calico should use.
	// Deprecates`RouteTableRange`, overrides `RouteTableRange`.
	RouteTableRanges *RouteTableRanges `json:"routeTableRanges,omitempty" validate:"omitempty,dive"`

	// Deprecated in favor of RouteTableRanges.
	// Calico programs additional Linux route tables for various purposes.
	// RouteTableRange specifies the indices of the route tables that Calico should use.
	RouteTableRange *RouteTableRange `json:"routeTableRange,omitempty" validate:"omitempty"`

	// EgressIPSupport defines three different support modes for egress IP function. [Default: Disabled]
	// - Disabled:                    Egress IP function is disabled.
	// - EnabledPerNamespace:         Egress IP function is enabled and can be configured on a per-namespace basis;
	//                                per-pod egress annotations are ignored.
	// - EnabledPerNamespaceOrPerPod: Egress IP function is enabled and can be configured per-namespace or per-pod,
	//                                with per-pod egress annotations overriding namespace annotations.
	// +kubebuilder:validation:Pattern=`^(?i)(Disabled|EnabledPerNamespace|EnabledPerNamespaceOrPerPod)?$`
	EgressIPSupport string `json:"egressIPSupport,omitempty" validate:"omitempty,oneof=Disabled EnabledPerNamespace EnabledPerNamespaceOrPerPod"`
	// EgressIPVXLANPort is the port number of vxlan tunnel device for egress traffic. [Default: 4790]
	EgressIPVXLANPort *int `json:"egressIPVXLANPort,omitempty"`
	// EgressIPVXLANVNI is the VNI ID of vxlan tunnel device for egress traffic. [Default: 4097]
	EgressIPVXLANVNI *int `json:"egressIPVXLANVNI,omitempty"`
	// EgressIPRoutingRulePriority controls the priority value to use for the egress IP routing rule. [Default: 100]
	EgressIPRoutingRulePriority *int `json:"egressIPRoutingRulePriority,omitempty" validate:"omitempty,gt=0,lt=32766"`
	// EgressGatewayPollInterval is the interval at which Felix will poll remote egress gateways to check their
	// health.  Only Egress Gateways with a named "health" port will be polled in this way.  Egress Gateways that
	// fail the health check will be taken our of use as if they have been deleted.
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	EgressGatewayPollInterval *metav1.Duration `json:"egressGatewayPollInterval,omitempty" configv1timescale:"seconds"`
	// EgressGatewayPollFailureCount is the minimum number of poll failures before a remote Egress Gateway is considered
	// to have failed.
	EgressGatewayPollFailureCount *int `json:"egressGatewayPollFailureCount,omitempty" validate:"omitempty,gt=0"`

	// RouteSyncDisabled will disable all operations performed on the route table. Set to true to
	// run in network-policy mode only.
	RouteSyncDisabled *bool `json:"routeSyncDisabled,omitempty"`

	// ExternalNetworkSupport defines two different support modes for external network function. [Default: Disabled]
	// - Disabled:  External network function is disabled.
	// - Enabled:   External network function is enabled.
	// +kubebuilder:validation:Pattern=`^(?i)(Disabled|Enabled)?$`
	ExternalNetworkSupport string `json:"externalNetworkSupport,omitempty" validate:"omitempty,oneof=Disabled Enabled"`
	// ExternalNetworkRoutingRulePriority controls the priority value to use for the external network routing rule. [Default: 102]
	ExternalNetworkRoutingRulePriority *int `json:"externalNetworkRoutingRulePriority,omitempty" validate:"omitempty,gt=0,lt=32766"`

	// WireguardEnabled controls whether Wireguard is enabled for IPv4 (encapsulating IPv4 traffic over an IPv4 underlay network). [Default: false]
	WireguardEnabled *bool `json:"wireguardEnabled,omitempty"`

	// WireguardEnabledV6 controls whether Wireguard is enabled for IPv6 (encapsulating IPv6 traffic over an IPv6 underlay network). [Default: false]
	WireguardEnabledV6 *bool `json:"wireguardEnabledV6,omitempty"`
	// WireguardThreadingEnabled controls whether Wireguard has NAPI threading enabled. [Default: false]
	WireguardThreadingEnabled *bool `json:"wireguardThreadingEnabled,omitempty"`
	// WireguardListeningPort controls the listening port used by IPv4 Wireguard. [Default: 51820]
	WireguardListeningPort *int `json:"wireguardListeningPort,omitempty" validate:"omitempty,gt=0,lte=65535"`

	// WireguardListeningPortV6 controls the listening port used by IPv6 Wireguard. [Default: 51821]
	WireguardListeningPortV6 *int `json:"wireguardListeningPortV6,omitempty" validate:"omitempty,gt=0,lte=65535"`

	// WireguardRoutingRulePriority controls the priority value to use for the Wireguard routing rule. [Default: 99]
	WireguardRoutingRulePriority *int `json:"wireguardRoutingRulePriority,omitempty" validate:"omitempty,gt=0,lt=32766"`

	// WireguardInterfaceName specifies the name to use for the IPv4 Wireguard interface. [Default: wireguard.cali]
	WireguardInterfaceName string `json:"wireguardInterfaceName,omitempty" validate:"omitempty,interface"`

	// WireguardInterfaceNameV6 specifies the name to use for the IPv6 Wireguard interface. [Default: wg-v6.cali]
	WireguardInterfaceNameV6 string `json:"wireguardInterfaceNameV6,omitempty" validate:"omitempty,interface"`

	// WireguardMTU controls the MTU on the IPv4 Wireguard interface. See Configuring MTU [Default: 1440]
	WireguardMTU *int `json:"wireguardMTU,omitempty"`

	// WireguardMTUV6 controls the MTU on the IPv6 Wireguard interface. See Configuring MTU [Default: 1420]
	WireguardMTUV6 *int `json:"wireguardMTUV6,omitempty"`

	// WireguardHostEncryptionEnabled controls whether Wireguard host-to-host encryption is enabled. [Default: false]
	WireguardHostEncryptionEnabled *bool `json:"wireguardHostEncryptionEnabled,omitempty"`

	// WireguardPersistentKeepAlive controls Wireguard PersistentKeepalive option. Set 0 to disable. [Default: 0]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	WireguardPersistentKeepAlive *metav1.Duration `json:"wireguardKeepAlive,omitempty"`

	// +kubebuilder:validation:MinLength=1
	// CaptureDir controls directory to store file capture. [Default: /var/log/calico/pcap]
	CaptureDir *string `json:"captureDir,omitempty" validate:"omitempty,gt=0"`

	// +kubebuilder:validation:Minimum=1
	// CaptureMaxSizeBytes controls the max size of a file capture. [Default: 10000000]
	CaptureMaxSizeBytes *int `json:"captureMaxSizeBytes,omitempty" validate:"omitempty,gt=0"`

	// +kubebuilder:validation:Minimum=1
	// CaptureRotationSeconds controls the time rotation of a packet capture. [Default: 3600]
	CaptureRotationSeconds *int `json:"captureRotationSeconds,omitempty" validate:"omitempty,gt=0"`

	// +kubebuilder:validation:Minimum=1
	// CaptureMaxFiles controls number of rotated capture file to keep. [Default: 2]
	CaptureMaxFiles *int `json:"captureMaxFiles,omitempty" validate:"omitempty,gt=0"`

	// AWSSrcDstCheck controls whether Felix will try to change the "source/dest check" setting on the EC2 instance
	// on which it is running. A value of "Disable" will try to disable the source/dest check. Disabling the check
	// allows for sending workload traffic without encapsulation within the same AWS subnet.
	// [Default: DoNothing]
	AWSSrcDstCheck *AWSSrcDstCheckOption `json:"awsSrcDstCheck,omitempty" validate:"omitempty,oneof=DoNothing Enable Disable"`

	// AWSSecondaryIPSupport controls whether Felix will try to provision AWS secondary ENIs for
	// workloads that have IPs from IP pools that are configured with an AWS subnet ID.  If the field is set to
	// "EnabledENIPerWorkload" then each workload with an AWS-backed IP will be assigned its own secondary ENI.
	// If set to "Enabled" then each workload with an AWS-backed IP pool will be allocated a secondary IP address
	// on a secondary ENI; this mode requires additional IP pools to be provisioned for the host to claim IPs for
	// the primary IP of the secondary ENIs. Accepted value must be one of "Enabled", "EnabledENIPerWorkload" or
	// "Disabled". [Default: Disabled]
	// +kubebuilder:validation:Pattern=`^(?i)(Enabled|EnabledENIPerWorkload|Disabled)?$`
	AWSSecondaryIPSupport string `json:"awsSecondaryIPSupport,omitempty" validate:"omitempty,oneof=Enabled EnabledENIPerWorkload Disabled"`
	// AWSSecondaryIPRoutingRulePriority controls the priority that Felix will use for routing rules when programming
	// them for AWS Secondary IP support. [Default: 101]
	AWSSecondaryIPRoutingRulePriority *int `json:"awsSecondaryIPRoutingRulePriority,omitempty" validate:"omitempty,gte=0,lte=4294967295"`
	// AWSRequestTimeout is the timeout on AWS API requests. [Default: 30s]
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`
	AWSRequestTimeout *metav1.Duration `json:"awsRequestTimeout,omitempty" configv1timescale:"seconds"`

	// When service IP advertisement is enabled, prevent routing loops to service IPs that are
	// not in use, by dropping or rejecting packets that do not get DNAT'd by kube-proxy.
	// Unless set to "Disabled", in which case such routing loops continue to be allowed.
	// [Default: Drop]
	// +kubebuilder:validation:Pattern=`^(?i)(Drop|Reject|Disabled)?$`
	ServiceLoopPrevention string `json:"serviceLoopPrevention,omitempty" validate:"omitempty,oneof=Drop Reject Disabled"`

	// WorkloadSourceSpoofing controls whether pods can use the allowedSourcePrefixes annotation to send traffic with a source IP
	// address that is not theirs. This is disabled by default. When set to "Any", pods can request any prefix.
	// +kubebuilder:validation:Pattern=`^(?i)(Disabled|Any)?$`
	WorkloadSourceSpoofing string `json:"workloadSourceSpoofing,omitempty" validate:"omitempty,oneof=Disabled Any"`

	// MTUIfacePattern is a regular expression that controls which interfaces Felix should scan in order
	// to calculate the host's MTU.
	// This should not match workload interfaces (usually named cali...).
	// +optional
	MTUIfacePattern string `json:"mtuIfacePattern,omitempty" validate:"omitempty,regexp"`

	// TPROXYMode sets whether traffic is directed through a transparent proxy
	// for further processing or not and how is the proxying done.
	// [Default: Disabled]
	// +kubebuilder:validation:Pattern=`^(?i)(Disabled|Enabled|EnabledAllServices)?$`
	TPROXYMode string `json:"tproxyMode,omitempty" validate:"omitempty,oneof=Disabled Enabled EnabledAllServices"`
	// TPROXYPort sets to which port proxied traffic should be redirected.
	// [Default: 16001]
	TPROXYPort *int `json:"tproxyPort,omitempty" validate:"omitempty,gt=0,lte=65535"`
	// TPROXYUpstreamConnMark tells Felix which mark is used by the proxy for its upstream
	// connections so that Felix can program the dataplane correctly.  [Default: 0x17]
	TPROXYUpstreamConnMark *uint32 `json:"tproxyUpstreamConnMark,omitempty" validate:"omitempty,gt=0"`

	// FloatingIPs configures whether or not Felix will program non-OpenStack floating IP addresses.  (OpenStack-derived
	// floating IPs are always programmed, regardless of this setting.)
	//
	// +optional
	FloatingIPs *FloatingIPType `json:"floatingIPs,omitempty" validate:"omitempty"`

	// WindowsManageFirewallRules configures whether or not Felix will program Windows Firewall rules (to allow inbound access to its own metrics ports). [Default: Disabled]
	// +optional
	WindowsManageFirewallRules *WindowsManageFirewallRulesMode `json:"windowsManageFirewallRules,omitempty" validate:"omitempty,oneof=Enabled Disabled"`

	// GoGCThreshold Sets the Go runtime's garbage collection threshold.  I.e. the percentage that the heap is
	// allowed to grow before garbage collection is triggered.  In general, doubling the value halves the CPU time
	// spent doing GC, but it also doubles peak GC memory overhead.  A special value of -1 can be used
	// to disable GC entirely; this should only be used in conjunction with the GoMemoryLimitMB setting.
	//
	// This setting is overridden by the GOGC environment variable.
	//
	// [Default: 40]
	// +optional
	GoGCThreshold *int `json:"goGCThreshold,omitempty" validate:"omitempty,gte=-1"`

	// GoMemoryLimitMB sets a (soft) memory limit for the Go runtime in MB.  The Go runtime will try to keep its memory
	// usage under the limit by triggering GC as needed.  To avoid thrashing, it will exceed the limit if GC starts to
	// take more than 50% of the process's CPU time.  A value of -1 disables the memory limit.
	//
	// Note that the memory limit, if used, must be considerably less than any hard resource limit set at the container
	// or pod level.  This is because felix is not the only process that must run in the container or pod.
	//
	// This setting is overridden by the GOMEMLIMIT environment variable.
	//
	// [Default: -1]
	// +optional
	GoMemoryLimitMB *int `json:"goMemoryLimitMB,omitempty" validate:"omitempty,gte=-1"`

	// GoMaxProcs sets the maximum number of CPUs that the Go runtime will use concurrently.  A value of -1 means
	// "use the system default"; typically the number of real CPUs on the system.
	//
	// this setting is overridden by the GOMAXPROCS environment variable.
	//
	// [Default: -1]
	// +optional
	GoMaxProcs *int `json:"goMaxProcs,omitempty" validate:"omitempty,gte=-1"`
}

type HealthTimeoutOverride struct {
	Name    string          `json:"name"`
	Timeout metav1.Duration `json:"timeout"`
}

type RouteTableRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type RouteTableIDRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type RouteTableRanges []RouteTableIDRange

func (r RouteTableRanges) NumDesignatedTables() int {
	var len int = 0
	for _, rng := range r {
		len += (rng.Max - rng.Min) + 1 // add one, since range is inclusive
	}

	return len
}

// ProtoPort is combination of protocol, port, and CIDR. Protocol and port must be specified.
type ProtoPort struct {
	Protocol string `json:"protocol,omitempty"`
	Port     uint16 `json:"port"`
	// +optional
	Net string `json:"net,omitempty"`
}

// New FelixConfiguration creates a new (zeroed) FelixConfiguration struct with the TypeMetadata
// initialized to the current version.
func NewFelixConfiguration() *FelixConfiguration {
	return &FelixConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindFelixConfiguration,
			APIVersion: GroupVersionCurrent,
		},
	}
}
