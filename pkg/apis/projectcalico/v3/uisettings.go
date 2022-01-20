// Copyright (c) 2021 Tigera, Inc. All rights reserved.

package v3

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	KindUISettings     = "UISettings"
	KindUISettingsList = "UISettingsList"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UISettings contains UI settings.
type UISettings struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the UISettings.
	Spec UISettingsSpec `json:"spec,omitempty"`
}

// UISettingsSpec contains the specification for a UISettings resource.
type UISettingsSpec struct {
	// The settings group. Once configured this cannot be modified. The group must exist.
	Group string `json:"group" validate:"name"`

	// This description is displayed by the UI.
	Description string `json:"description" validate:"uiDescription"`

	// View data. One of View, Layer or Dashboard should be specified.
	View *UIGraphView `json:"view,omitempty" validate:"omitempty"`

	// Layer data. One of View, Layer or Dashboard should be specified.
	Layer *UIGraphLayer `json:"layer,omitempty" validate:"omitempty"`

	// Dashboard data. One of View, Layer or Dashboard should be specified.
	Dashboard *UIDashboard `json:"dashboard,omitempty" validate:"omitempty"`
}

// UIGraphView contains the data for a UI graph view.
type UIGraphView struct {
	// Deprecated: Use Nodes.
	// The set of nodes that are the focus of the graph. All nodes returned by the graph query will be connected to at
	// least one of these nodes. If this is empty, then all nodes will be returned.
	Focus []UIGraphNode `json:"focus,omitempty" validate:"omitempty,dive"`

	// Deprecated: Use Nodes.
	// The set of nodes that are expanded to the next level of expansion.
	Expanded []UIGraphNode `json:"expanded,omitempty" validate:"omitempty,dive"`

	// Whether ports are expanded. If false, port information is aggregated.
	ExpandPorts bool `json:"expandPorts" validate:"omitempty"`

	// Whether or not to automatically follow directly connected nodes.
	FollowConnectionDirection bool `json:"followConnectionDirection" validate:"omitempty"`

	// Whether to split HostEndpoints, NetworkSets and Networks into separate ingress and egress nodes or to combine
	// them. In a service-centric view, splitting these makes the graph clearer. This never splits pods which represent
	// a true microservice which has ingress and egress connections.
	SplitIngressEgress bool `json:"splitIngressEgress" validate:"omitempty"`

	// The set of selectors used to aggregate hosts (Kubernetes nodes). Nodes are aggregated based on the supplied set
	// of selectors. In the case of overlapping selectors, the order specified in the slice is the order checked and so
	// the first selector to match is used.  The nodes will be aggregated into a graph node with the name specified in
	// the NamedSelector.
	HostAggregationSelectors []NamedSelector `json:"hostAggregationSelectors,omitempty" validate:"omitempty,dive"`

	// Deprecated: Use Nodes.
	// Followed nodes. These are nodes on the periphery of the graph that we follow further out of the scope of the
	// graph focus. For example a Node N may have egress connections to X and Y, but neither X nor Y are displayed in
	// the graph because they are not explicitly in focus. The service graph response will indicate that Node N has
	// egress connections that may be followed.  If Node N is added to this "FollowedEgress" then the response will
	// include the egress connections to X and Y.
	FollowedEgress []UIGraphNode `json:"followedEgress,omitempty" validate:"omitempty,dive"`
	// Deprecated: Use Nodes.
	FollowedIngress []UIGraphNode `json:"followedIngress,omitempty" validate:"omitempty,dive"`

	// Layout type. Semi-arbitrary value used to specify the layout-type/algorithm. For example could specify
	// different layout algorithms, or click-to-grid.  Mostly here for future use.
	LayoutType string `json:"layoutType" validate:"omitempty"`

	// Positions of graph nodes.
	Positions []Position `json:"positions" validate:"omitempty,dive"`

	// The set of layer names that are active in this view.  Note that layers may be defined, but it is not necessary
	// to have each layer "active". Corresponds directly to the name of the UISettings resource that contains a layer
	// definition.
	Layers []string `json:"layers" validate:"omitempty,dive"`

	// View information for grpah nodes. This provides information about what is in focus, expanded, hidden,
	// deemphasized etc. at a per-node level.
	Nodes []UIGraphNodeView `json:"graphConfigurationByNode" validate:"omitempty,dive"`
}

// UI screen position.
type Position struct {
	ID   string `json:"id" validate:"servicegraphId"`
	XPos int    `json:"xPos"`
	YPos int    `json:"yPos"`
	ZPos int    `json:"zPos"`
}

// A Calico format label selector with an associated name.
type NamedSelector struct {
	Name     string `json:"name" validate:"uiDescription"`
	Selector string `json:"selector" validate:"selector"`
}

// UIGraphLayer contains the data for a UI graph layer.
type UIGraphLayer struct {
	// The nodes that are aggregated into a single layer.
	Nodes []UIGraphNode `json:"nodes" validate:"omitempty,dive"`

	// A user-configurable icon. If not specified, the default layer icon is used for this layer node.
	Icon string `json:"icon,omitempty" validate:"omitempty,icon"`

	// The color used to represent the layer when an Icon has not been specified.
	Color string `json:"color,omitempty" validate:"omitempty,color"`
}

type UIGraphNodeView struct {
	UIGraphNode `json:",inline"`

	// This node is a primary focus of the graph (i.e. the graph contains this node and connected nodes).
	InFocus bool `json:"inFocus,omitempty" validate:"omitempty"`

	// This node is expanded to the next level.  This node can, for example, be a layer that is expanded into its
	// constituent parts.
	Expanded bool `json:"expanded,omitempty" validate:"omitempty"`

	// Whether the ingress/egress connections to/from this node are included in the graph.  This effectively brings
	// more nodes into focus.
	FollowIngress bool `json:"followIngress,omitempty" validate:"omitempty"`
	FollowEgress  bool `json:"followEgress,omitempty" validate:"omitempty"`

	// Whether the UI should de-emphasize the node when visible. This is just a UI directive and does not correspond to
	// a backend parameter.
	Deemphasize bool `json:"deemphasize,omitempty" validate:"omitempty"`

	// Whether the UI should hide the node. This is just a UI directive and does not correspond to a backend parameter.
	Hide bool `json:"hide,omitempty" validate:"omitempty"`

	// Whether the UI should hide unrelated nodes. This is just a UI directive and does not correspond to a backend
	// parameter.
	HideUnrelated bool `json:"hideUnrelated,omitempty" validate:"omitempty"`
}

// UIGraphNode contains details about a graph node so that the UI can render it correctly.  Mostly this provides the
// UI a mechanism for caching name/type information associated with a node ID.  However, it also provides the UI a way
// to provide an alternative name, icon or color for the node.
type UIGraphNode struct {
	// The node ID.
	ID string `json:"id" validate:"servicegraphId"`

	// The node type.
	Type string `json:"type" validate:"servicegraphNodeType"`

	// The node name.
	Name string `json:"name" validate:"name"`

	// The node namespace.
	Namespace string `json:"namespace,omitempty" validate:"omitempty,name"`
}

// UIDashboard contains the data for a UI dashboard.
type UIDashboard struct {
	// TBD
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UISettingsList contains a list of UISettings resources.
type UISettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []UISettings `json:"items"`
}

// NewUISettings creates a new (zeroed) UISettings struct with the TypeMetadata
// initialised to the current version.
func NewUISettings() *UISettings {
	return &UISettings{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindUISettings,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewUISettingsList creates a new (zeroed) UISettingsList struct with the
// TypeMetadata initialised to the current version.
func NewUISettingsList() *UISettingsList {
	return &UISettingsList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindUISettingsList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
