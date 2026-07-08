// Copyright (c) 2026 Tigera, Inc. All rights reserved.

package v3

// L2BridgeSpec configures a Layer 2 bridged network.  Workload interfaces
// are attached to a Linux bridge on each node as 802.1Q access ports; the
// bridge connects to the physical network via a single trunk port defined
// in HostConfig.  (HostConnections is currently capped at one entry per
// HostConfig; multiple trunk ports per host is a future extension.)
//
// Calico-managed bridges run in tagged-only mode: the bridge has
// vlan_filtering=1 and vlan_default_pvid=0, so untagged frames are dropped
// on the wire by default.  Operators that need untagged-on-wire delivery
// for a trunk opt in by setting L2HostTrunkPort.NativeVLAN to the VLAN ID
// that should be carried untagged on that trunk.
type L2BridgeSpec struct {
	// VLANs is the authoritative list of 802.1Q VLAN segments carried by
	// this network.  Each entry defines a single VLAN ID or a contiguous
	// range of VLAN IDs, plus the subnets associated with that segment.
	//
	// The set of VLAN IDs the network may carry is the union of all IDs and
	// ranges in this list.  Entries are not required to be disjoint: if the
	// same VLAN ID appears in more than one entry the network still simply
	// carries that VLAN, so overlap is permitted and not validated here.
	//
	// Workload attachments select a single entry via the CNI config
	// "vlan" field; if spec.vlans has exactly one entry that resolves to
	// a single VLAN ID, the CNI config "vlan" field may be omitted.
	// +kubebuilder:validation:MinItems=1
	// +listType=atomic
	VLANs []L2VLANSpec `json:"vlans" validate:"required,min=1,dive"`

	// HostConfig defines per-node-group configuration for this network.
	// When multiple entries are present, the first entry whose
	// nodeSelector matches a given node is used; all other entries are
	// ignored for that node.  Entries with no nodeSelector match all
	// nodes.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=100
	// +listType=map
	// +listMapKey=nodeSelector
	HostConfig []L2HostConfig `json:"hostConfig" validate:"required,min=1,dive"`
}

// L2VLANSpec defines a VLAN segment and its optional IP configuration.
type L2VLANSpec struct {
	// VLAN identifies this segment: a single VLAN ID or a contiguous
	// range of VLAN IDs.
	VLAN L2VLANMatch `json:"vlan" validate:"required"`

	// Subnets are the IP subnets for this VLAN segment.  They filter which
	// IPPools the IPAM plugin considers for this VLAN (only pools that fall
	// within one of these subnets are eligible) and supply the prefix
	// length programmed on the pod interface.  Multiple entries support
	// dual-stack or multi-subnet VLANs.
	//
	// IPAM allocates the address from the eligible pools without preferring
	// any particular subnet; the allocated address then falls within
	// exactly one of these subnets, and that subnet's CIDR prefix length is
	// the one programmed on the pod interface.  For a dual-stack VLAN this
	// resolves independently per family (one IPv4 and one IPv6 subnet).
	// +optional
	// +listType=atomic
	Subnets []L2Subnet `json:"subnets,omitempty" validate:"omitempty,dive"`
}

// L2VLANMatch identifies a VLAN segment.  Exactly one field must be set.
//
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type L2VLANMatch struct {
	// ID selects a single 802.1Q VLAN (1-4094).
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4094
	ID *uint16 `json:"id,omitempty" validate:"omitempty,gte=1,lte=4094"`

	// Range selects a contiguous range of VLAN IDs (inclusive).
	// +optional
	Range *L2VLANRange `json:"range,omitempty"`
}

// L2VLANRange represents a contiguous, inclusive range of 802.1Q VLAN IDs.
// Start must be ≤ End.
//
// +kubebuilder:validation:XValidation:rule="self.start <= self.end",message="start must be <= end"
type L2VLANRange struct {
	// Start is the first VLAN ID in the range (1-4094).
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4094
	Start uint16 `json:"start" validate:"required,gte=1,lte=4094"`

	// End is the last VLAN ID in the range (1-4094, must be ≥ Start).
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4094
	End uint16 `json:"end" validate:"required,gte=1,lte=4094,gtefield=Start"`
}

// L2Subnet defines an IP subnet associated with a VLAN segment.
type L2Subnet struct {
	// CIDR is the subnet in CIDR notation (e.g. "10.100.0.0/24",
	// "fd00::/64").
	CIDR string `json:"cidr" validate:"required,cidr"`

	// Routes are programmed inside the pod's per-interface routing table
	// for workloads attached to this subnet, in addition to the connected
	// route for CIDR itself.  Each next hop must be an L2 neighbor on this
	// subnet, reachable via the bridge; routes are installed onlink.  The
	// common "default gateway" case is a single entry with destination
	// "0.0.0.0/0" (or "::/0") whose nextHop is the upstream router.
	// +optional
	// +listType=map
	// +listMapKey=destination
	Routes []L2Route `json:"routes,omitempty" validate:"omitempty,dive"`
}

// L2Route defines a route programmed in a pod's per-interface routing table
// for an L2 bridge subnet.  Parallel to VRFStaticRoute, but with its own
// action type so L2 and VRF route semantics can evolve independently.
type L2Route struct {
	// Destination is the CIDR prefix for this route.  Use "0.0.0.0/0" or
	// "::/0" for a default route.
	// +required
	Destination string `json:"destination" validate:"required,cidr"`

	// Action determines how traffic matching this route is handled.
	// Exactly one action field must be set.
	// +required
	Action L2RouteAction `json:"action" validate:"required"`
}

// L2RouteAction defines the forwarding behaviour for an L2 route.  Exactly
// one field must be set.
//
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type L2RouteAction struct {
	// NextHop forwards matching traffic to the specified gateway IP.  The
	// address must be an L2 neighbor on this subnet, reachable via the
	// bridge; the route is installed onlink.
	// +optional
	NextHop *string `json:"nextHop,omitempty" validate:"omitempty,ip"`
}

// L2HostConfig provides node-specific L2 bridge settings.  Different nodes
// may have different bridge types or trunk interface names; each entry
// applies to the nodes matched by its NodeSelector (first-match wins).
type L2HostConfig struct {
	// NodeSelector is a Calico selector expression that determines which
	// nodes this configuration applies to.  If omitted, the entry
	// applies to all nodes.  When multiple HostConfig entries are
	// present, the first entry whose selector matches a given node wins;
	// subsequent entries are ignored for that node.
	// +optional
	// +kubebuilder:default=""
	NodeSelector string `json:"nodeSelector,omitempty" validate:"omitempty,selector"`

	// Bridge selects the bridge device for these nodes: either a
	// Calico-managed bridge or a pre-existing (BYO) bridge.
	Bridge L2BridgeDevice `json:"bridge" validate:"required"`

	// HostConnections defines how the bridge connects to the physical
	// network on these nodes.  Each entry is a typed connection
	// (currently only trunkPort).  If omitted, no trunk is attached by
	// Calico — the user is responsible for providing external
	// connectivity (e.g. on a BYO bridge with a pre-configured trunk).
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +listType=atomic
	HostConnections []L2HostConnection `json:"hostConnections,omitempty" validate:"omitempty,dive"`
}

// L2BridgeDevice selects between a Calico-managed bridge and a
// pre-existing user-managed bridge.  Exactly one field must be set.
//
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type L2BridgeDevice struct {
	// ManagedBridge instructs Calico to create and fully manage the
	// bridge device.  The bridge name is derived automatically from the
	// Network name so it does not clash with other Networks or with
	// Calico workload veth names.
	// +optional
	ManagedBridge *L2ManagedBridge `json:"managedBridge,omitempty"`

	// ExistingBridge instructs Calico to attach to a pre-existing bridge
	// created and managed externally (e.g. via netplan).  Calico does
	// not create, reconfigure, or delete the bridge; it does not add or
	// remove IP addresses on it; and it does not remove interfaces it
	// did not create.  Calico still attaches workload veths, configures
	// their VLAN membership, and connects trunk interfaces (if
	// hostConnections is specified) unless the trunk is already
	// connected.
	// +optional
	ExistingBridge *L2ExistingBridge `json:"existingBridge,omitempty"`
}

// L2ManagedBridge configures a bridge device that Calico creates and
// manages.
type L2ManagedBridge struct {
	// STP controls whether Spanning Tree Protocol is active on the
	// bridge.  "Disabled" (default) is appropriate for datacenter
	// topologies where loop prevention is handled by the upstream
	// switch; "Enabled" causes the bridge to exchange BPDU frames on
	// trunk ports.  Workload veth ports always operate in edge/portfast
	// mode regardless of this setting — they never participate in STP.
	// +optional
	// +kubebuilder:validation:Enum=Enabled;Disabled
	// +kubebuilder:default=Disabled
	STP string `json:"stp,omitempty" validate:"omitempty,oneof=Enabled Disabled"`
}

// L2ExistingBridge configures attachment to a pre-existing bridge device.
type L2ExistingBridge struct {
	// Name is the Linux bridge device name (e.g. "br0").
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=15
	Name string `json:"name" validate:"required,interface"`
}

// L2HostConnection defines a single host-side connection to the bridge.
// Exactly one connection type must be set.
//
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type L2HostConnection struct {
	// TrunkPort enslaves a host interface to the bridge as an 802.1Q
	// trunk.
	// +optional
	TrunkPort *L2HostTrunkPort `json:"trunkPort,omitempty"`
}

// L2HostTrunkPort configures a trunk port on the bridge.
//
// +kubebuilder:validation:XValidation:rule="!has(self.nativeVLAN) || self.nativeVLAN >= 1 && self.nativeVLAN <= 4094",message="nativeVLAN must be between 1 and 4094"
type L2HostTrunkPort struct {
	// Interface identifies the trunk interface on this node group.
	// Heterogeneous clusters use separate HostConfig entries with
	// different NodeSelectors rather than per-rule selectors.
	Interface InterfaceMatch `json:"interface" validate:"required"`

	// NativeVLAN, when set, designates one VLAN ID as the trunk's native
	// VLAN: frames in that VLAN are sent untagged on the wire and
	// untagged frames received on the trunk are tagged with this VLAN.
	// This should be one of the VLAN IDs the trunk carries (i.e. present
	// in spec.vlans); if it names a VLAN the trunk does not carry it is
	// ignored and the trunk stays strictly tagged for every VLAN.  When
	// unset (default) the trunk is strictly tagged: only 802.1Q-tagged
	// frames are accepted and transmitted.
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4094
	NativeVLAN *uint16 `json:"nativeVLAN,omitempty" validate:"omitempty,gte=1,lte=4094"`
}
