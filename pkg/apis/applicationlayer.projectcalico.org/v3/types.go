// Copyright (c) 2026 Tigera, Inc. All rights reserved.

// Package v3 contains API Schema definitions for the applicationlayer.projectcalico.org v3 API group
// +kubebuilder:object:generate=true
// +groupName=applicationlayer.projectcalico.org
package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GroupVersionCurrent is the API version string for this group.
const GroupVersionCurrent = GroupName + "/v3"

// NewWAFPolicy creates a new (zeroed) WAFPolicy with TypeMeta initialised.
func NewWAFPolicy() *WAFPolicy {
	return &WAFPolicy{TypeMeta: metav1.TypeMeta{Kind: "WAFPolicy", APIVersion: GroupVersionCurrent}}
}

// NewWAFPolicyList creates a new (zeroed) WAFPolicyList with TypeMeta initialised.
func NewWAFPolicyList() *WAFPolicyList {
	return &WAFPolicyList{TypeMeta: metav1.TypeMeta{Kind: "WAFPolicyList", APIVersion: GroupVersionCurrent}}
}

// NewGlobalWAFPolicy creates a new (zeroed) GlobalWAFPolicy with TypeMeta initialised.
func NewGlobalWAFPolicy() *GlobalWAFPolicy {
	return &GlobalWAFPolicy{TypeMeta: metav1.TypeMeta{Kind: "GlobalWAFPolicy", APIVersion: GroupVersionCurrent}}
}

// NewGlobalWAFPolicyList creates a new (zeroed) GlobalWAFPolicyList with TypeMeta initialised.
func NewGlobalWAFPolicyList() *GlobalWAFPolicyList {
	return &GlobalWAFPolicyList{TypeMeta: metav1.TypeMeta{Kind: "GlobalWAFPolicyList", APIVersion: GroupVersionCurrent}}
}

// NewWAFPlugin creates a new (zeroed) WAFPlugin with TypeMeta initialised.
func NewWAFPlugin() *WAFPlugin {
	return &WAFPlugin{TypeMeta: metav1.TypeMeta{Kind: "WAFPlugin", APIVersion: GroupVersionCurrent}}
}

// NewWAFPluginList creates a new (zeroed) WAFPluginList with TypeMeta initialised.
func NewWAFPluginList() *WAFPluginList {
	return &WAFPluginList{TypeMeta: metav1.TypeMeta{Kind: "WAFPluginList", APIVersion: GroupVersionCurrent}}
}

// NewGlobalWAFPlugin creates a new (zeroed) GlobalWAFPlugin with TypeMeta initialised.
func NewGlobalWAFPlugin() *GlobalWAFPlugin {
	return &GlobalWAFPlugin{TypeMeta: metav1.TypeMeta{Kind: "GlobalWAFPlugin", APIVersion: GroupVersionCurrent}}
}

// NewGlobalWAFPluginList creates a new (zeroed) GlobalWAFPluginList with TypeMeta initialised.
func NewGlobalWAFPluginList() *GlobalWAFPluginList {
	return &GlobalWAFPluginList{TypeMeta: metav1.TypeMeta{Kind: "GlobalWAFPluginList", APIVersion: GroupVersionCurrent}}
}

// NewGlobalWAFValidationPolicy creates a new (zeroed) GlobalWAFValidationPolicy with TypeMeta initialised.
func NewGlobalWAFValidationPolicy() *GlobalWAFValidationPolicy {
	return &GlobalWAFValidationPolicy{TypeMeta: metav1.TypeMeta{Kind: "GlobalWAFValidationPolicy", APIVersion: GroupVersionCurrent}}
}

// NewGlobalWAFValidationPolicyList creates a new (zeroed) GlobalWAFValidationPolicyList with TypeMeta initialised.
func NewGlobalWAFValidationPolicyList() *GlobalWAFValidationPolicyList {
	return &GlobalWAFValidationPolicyList{TypeMeta: metav1.TypeMeta{Kind: "GlobalWAFValidationPolicyList", APIVersion: GroupVersionCurrent}}
}

// NewWAFValidationPolicy creates a new (zeroed) WAFValidationPolicy with TypeMeta initialised.
func NewWAFValidationPolicy() *WAFValidationPolicy {
	return &WAFValidationPolicy{TypeMeta: metav1.TypeMeta{Kind: "WAFValidationPolicy", APIVersion: GroupVersionCurrent}}
}

// NewWAFValidationPolicyList creates a new (zeroed) WAFValidationPolicyList with TypeMeta initialised.
func NewWAFValidationPolicyList() *WAFValidationPolicyList {
	return &WAFValidationPolicyList{TypeMeta: metav1.TypeMeta{Kind: "WAFValidationPolicyList", APIVersion: GroupVersionCurrent}}
}

// WAFAction defines the action to take when a rule matches
// +kubebuilder:validation:Enum=Detect;Block
type WAFAction string

const (
	WAFActionDetect WAFAction = "Detect"
	WAFActionBlock  WAFAction = "Block"
)

// CoreRuleSetState enables or disables the OWASP CRS baseline.
// +kubebuilder:validation:Enum=Enabled;Disabled
type CoreRuleSetState string

const (
	CoreRuleSetStateEnabled  CoreRuleSetState = "Enabled"
	CoreRuleSetStateDisabled CoreRuleSetState = "Disabled"
)

// CoreRuleSetConfig defines OWASP Core Rule Set configuration
type CoreRuleSetConfig struct {
	// State enables or disables the OWASP CRS baseline.
	// +kubebuilder:default=Enabled
	State CoreRuleSetState `json:"state,omitempty"`

	// Enabled is the deprecated boolean form. Read for backward compatibility
	// with v0.2.x CRs; new CRs should set State. If both fields are set,
	// State wins.
	//
	// Deprecated: use State.
	Enabled *bool `json:"enabled,omitempty"`

	// ParanoiaLevel sets the CRS paranoia level (1-4)
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4
	// +kubebuilder:default=1
	ParanoiaLevel int `json:"paranoiaLevel,omitempty"`
}

// EffectiveState returns the resolved CoreRuleSetState given State and the
// deprecated Enabled field. State wins if explicitly set; otherwise the
// deprecated Enabled is mapped (true -> Enabled, false -> Disabled).
// Both unset -> Enabled (kubebuilder default).
func (c *CoreRuleSetConfig) EffectiveState() CoreRuleSetState {
	if c.State != "" {
		return c.State
	}
	if c.Enabled != nil && !*c.Enabled {
		return CoreRuleSetStateDisabled
	}
	return CoreRuleSetStateEnabled
}

// PluginRef references a plugin by kind + name. Kind makes the
// reference target explicit; without it, "is this naming a Global or a
// namespace plugin?" had to be inferred from which policy's spec.plugins[]
// the reference appeared in. The default is WAFPlugin (namespace-scoped) for
// backward-compat with v0.2.x CRs that omit Kind.
//
// Note: in this release the resolved scope still follows the parent policy
// (GlobalWAFPolicy.spec.plugins[] resolves as Global; WAFPolicy.spec.plugins[]
// resolves as namespace-scoped). Cross-scope lookup (e.g. WAFPolicy
// referencing a GlobalWAFPlugin) is a future addition; the Kind field
// here is type-level disclosure of intent that the reconciler can validate
// against in a later release.
type PluginRef struct {
	// Kind selects between namespace-scoped (WAFPlugin) and cluster-scoped
	// (GlobalWAFPlugin) plugin types.
	// +kubebuilder:default=WAFPlugin
	// +kubebuilder:validation:Enum=WAFPlugin;GlobalWAFPlugin
	Kind string `json:"kind,omitempty"`

	// Name is the name of the plugin resource.
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`
}

// PolicyRef identifies a policy that references a plugin. Used by plugin
// status.referencedBy to give kubectl-readable reverse-index entries
// without packing identifiers into a single string.
//
// Current scope-binding (this release): the plugin reconciler resolves
// references from the parent policy's scope — WAFPolicy.spec.plugins[]
// resolves namespace-scoped, GlobalWAFPolicy.spec.plugins[] resolves
// cluster-scoped. Cross-scope lookup is a future addition; the Kind field
// below is type-level disclosure of intent for a later release.
type PolicyRef struct {
	// Name is the name of the referencing policy.
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`

	// Namespace is the namespace of the referencing policy. Unset for
	// cluster-scoped GlobalWAFPolicy entries (pointer so unset is
	// distinguishable from explicit empty, per the optional-fields-are-
	// pointers convention).
	// +optional
	// +kubebuilder:validation:MaxLength=253
	Namespace *string `json:"namespace,omitempty"`

	// Kind identifies the referencing policy kind.
	// +kubebuilder:validation:Enum=WAFPolicy;GlobalWAFPolicy
	Kind string `json:"kind"`
}

// PolicyTargetReference identifies a Gateway API resource
type PolicyTargetReference struct {
	// Group is the group of the target resource
	// +kubebuilder:default="gateway.networking.k8s.io"
	// +kubebuilder:validation:MaxLength=253
	Group string `json:"group,omitempty"`

	// Kind is the kind of the target resource
	// +kubebuilder:validation:Enum=Gateway;HTTPRoute
	// +kubebuilder:validation:MaxLength=253
	Kind string `json:"kind"`

	// Name is the name of the target resource
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`

	// SectionName scopes the WAF to a single named section within the target
	// (GEP-713): an HTTPRoute rule name (spec.rules[].name) or a Gateway listener
	// name (spec.listeners[].name). When unset, the policy attaches to the entire
	// target. The named rule/listener must exist, or the policy reports
	// Programmed=False/TargetNotFound. A rule-scoped policy takes precedence over
	// a whole-route one for that rule (Envoy Gateway GEP-713 precedence). See
	// designs#25 §01.
	// +optional
	// +kubebuilder:validation:MaxLength=253
	SectionName *string `json:"sectionName,omitempty"`
}

// =============================================================================
// GlobalWAFPolicy - Cluster-scoped WAF policy
// =============================================================================

// GlobalWAFPolicySpec defines the desired state
type GlobalWAFPolicySpec struct {
	// DefaultAction specifies what happens when a rule matches
	// +kubebuilder:default=Detect
	DefaultAction WAFAction `json:"defaultAction,omitempty"`

	// CoreRuleSet configures the OWASP Core Rule Set
	CoreRuleSet CoreRuleSetConfig `json:"coreRuleSet,omitempty"`

	// Plugins references GlobalWAFPlugins
	// These plugins are applied to all namespaces
	// +kubebuilder:validation:MaxItems=64
	Plugins []PluginRef `json:"plugins,omitempty"`
}

// GlobalWAFPolicyStatus defines the observed state.
//
// Conditions emitted by the controller:
//   - Licensed   (LicenseValid / LicenseBlocked)
//   - Accepted   (Accepted / Invalid / PluginNotFound)
//   - Ready      (Ready / NotReady / ConflictingGlobalPolicy)
type GlobalWAFPolicyStatus struct {
	// Conditions represent the latest available observations
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// NamespaceCount is the number of namespaces in which a WAFPolicy targets
	// resources covered by this Global policy. Different denominator from
	// GlobalWAFPluginStatus.NamespaceCount (which counts namespaces where a
	// policy references the plugin).
	NamespaceCount int `json:"namespaceCount,omitempty"`

	// LastApplied is the timestamp of the last successful application
	LastApplied *metav1.Time `json:"lastApplied,omitempty"`

	// RenderedConfigMapRefs enumerates ConfigMaps the controller has emitted
	// from this GlobalWAFPolicy, one per namespace that consumed it. Capped
	// at 50 entries; see RenderedConfigMapRefsTruncated. When the cap is
	// exceeded, the surfaced entries are the lexically-first 50 namespace
	// names (sorted ascending by namespace), so the truncation is
	// deterministic across reconciles; use NamespaceCount to recover the
	// untruncated total.
	RenderedConfigMapRefs []RenderedConfigMapRef `json:"renderedConfigMapRefs,omitempty"`

	// RenderedConfigMapRefsTruncated is true when the full set exceeded the
	// 50-entry status cap. Consumers should use NamespaceCount for totals.
	RenderedConfigMapRefsTruncated bool `json:"renderedConfigMapRefsTruncated,omitempty"`

	// InheritedFromGlobal is always nil for GlobalWAFPolicy. The field is
	// declared for symmetry with WAFPolicyStatus so generated clients (Go,
	// TS) can share a single attribution shape across scopes.
	// +optional
	InheritedFromGlobal *string `json:"inheritedFromGlobal,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,shortName=gwafp
// +kubebuilder:printcolumn:name="Action",type=string,JSONPath=`.spec.defaultAction`
// +kubebuilder:printcolumn:name="CRS",type=string,JSONPath=`.spec.coreRuleSet.state`
// +kubebuilder:printcolumn:name="Paranoia",type=integer,JSONPath=`.spec.coreRuleSet.paranoiaLevel`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:validation:XValidation:rule="self.metadata.name == 'default'",message="GlobalWAFPolicy is a singleton; the only permitted name is 'default'"

// GlobalWAFPolicy is the Schema for cluster-wide WAF configuration
type GlobalWAFPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalWAFPolicySpec   `json:"spec,omitempty"`
	Status GlobalWAFPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlobalWAFPolicyList contains a list of GlobalWAFPolicy
type GlobalWAFPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalWAFPolicy `json:"items"`
}

// =============================================================================
// WAFPolicy - Namespace-scoped WAF policy
// =============================================================================

// WAFPolicySpec defines the desired state
// +kubebuilder:validation:XValidation:rule="!has(self.targetRefs) || self.targetRefs.all(ref, ref.group == 'gateway.networking.k8s.io')",message="only group gateway.networking.k8s.io is supported for targetRefs"
type WAFPolicySpec struct {
	// TargetRefs specifies Gateway API references (Gateway, HTTPRoute)
	// +kubebuilder:validation:MaxItems=16
	// +kubebuilder:validation:MinItems=1
	TargetRefs []PolicyTargetReference `json:"targetRefs"`

	// Action overrides the global default action for this namespace.
	// Nil means inherit from GlobalWAFPolicy; an explicit value overrides
	// the global default. Subject to validation policy enforcement.
	// +optional
	Action *WAFAction `json:"action,omitempty"`

	// CoreRuleSet configures the OWASP Core Rule Set for this namespace
	// If not specified, inherits from GlobalWAFPolicy
	CoreRuleSet *CoreRuleSetConfig `json:"coreRuleSet,omitempty"`

	// Plugins references WAFPlugins in this namespace. A cross-scope reference
	// (kind: GlobalWAFPlugin) is rejected at admission by the CEL rule below and
	// the validating webhook, because the reconciler resolves plugin scope from
	// the parent policy and does not yet honor Kind. See EV-6753.
	//
	// Only this namespaced direction is guarded. The mirror case, a
	// GlobalWAFPolicy referencing a namespaced WAFPlugin, is intentionally not
	// guarded: PluginRef.Kind defaults to WAFPlugin (see above), so a symmetric
	// rule would reject that default on nearly every GlobalWAFPolicy. The
	// cross-scope-resolution follow-up owns the reverse guard, the Kind default,
	// and an upgrade migration; see the WAF controllers README.
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:XValidation:rule="self.all(p, !has(p.kind) || p.kind != 'GlobalWAFPlugin')",message="cross-scope plugin references are not yet supported: a namespaced WAFPolicy cannot reference a GlobalWAFPlugin"
	Plugins []PluginRef `json:"plugins,omitempty"` // rule's !has(p.kind) is defensive: Kind is defaulted to WAFPlugin upstream, so an empty Kind never reaches CEL today.
}

// RenderedConfigMapRef identifies the ConfigMap into which the controller
// materialised the merged WAF rule set for a policy. CLI tools and status
// consumers resolve this reference instead of reconstructing the name.
type RenderedConfigMapRef struct {
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`
	// +kubebuilder:validation:MaxLength=253
	Namespace string `json:"namespace"`
	// +kubebuilder:validation:MaxLength=253
	ResourceVersion string `json:"resourceVersion,omitempty"`
}

// ValidationState represents validation results decorated onto a WAF policy
type ValidationState struct {
	// Status is the overall validation outcome. Denormalized from the
	// Validated x Programmed conditions so the `kubectl get wafpolicy`
	// printcolumn can show a single readable scalar; CRD printcolumns can
	// only read one JSONPath, not compute from two conditions. The
	// conditions are authoritative — if this field and the conditions ever
	// diverge in a future change, the conditions win and this field is the
	// rendered projection. Valid = Validated/True + Programmed/True;
	// Audited = Validated/False + Programmed/True (Audit mode);
	// Rejected = Validated/False + Programmed/False (Enforce mode blocked).
	// +kubebuilder:validation:Enum=Valid;Audited;Rejected
	Status string `json:"status"`

	// SecurityPosture is the security posture based on all validation results
	SecurityPosture SecurityPosture `json:"securityPosture,omitempty"`

	// Failures contains details about failed validation rules
	Failures []ValidationFailure `json:"failures,omitempty"`

	// LastEvaluated is when validation was last performed
	LastEvaluated *metav1.Time `json:"lastEvaluated,omitempty"`
}

// ValidationFailure describes a single validation rule failure
type ValidationFailure struct {
	// PolicyName is the name of the validation policy that failed
	// +kubebuilder:validation:MaxLength=253
	PolicyName string `json:"policyName"`

	// PolicyKind is Global or Namespace-scoped
	// +kubebuilder:validation:Enum=GlobalWAFValidationPolicy;WAFValidationPolicy
	// +kubebuilder:validation:MaxLength=253
	PolicyKind string `json:"policyKind"`

	// Rule is the name of the specific rule that failed
	// +kubebuilder:validation:MaxLength=63
	Rule string `json:"rule"`

	// Severity of the failure
	Severity ValidationSeverity `json:"severity"`

	// Message explains why validation failed
	// +kubebuilder:validation:MaxLength=1024
	Message string `json:"message,omitempty"`
}

// WAFPolicyStatus defines the observed state.
//
// Conditions emitted by the controller:
//   - Licensed   (LicenseValid / LicenseInGracePeriod / LicenseExpired /
//     LicenseInvalid / LicenseBlocked)
//   - Accepted   (Accepted / Invalid / Conflicted / PluginNotFound /
//     ReplicaUnmanaged / ValidationFailed)
//   - Programmed (ConfigurationApplied / WASMUnavailable / CheckerError /
//     NotAttempted)
//   - Ready      (Ready / NotReady)
type WAFPolicyStatus struct {
	// Conditions represent the latest available observations
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// PluginCount is the number of plugins applied
	PluginCount int `json:"pluginCount,omitempty"`

	// LastApplied is the timestamp of the last successful application
	LastApplied *metav1.Time `json:"lastApplied,omitempty"`

	// Validation contains the results of validation policy evaluation.
	// Shows whether the policy passed, was audited (warnings logged), or rejected.
	Validation *ValidationState `json:"validation,omitempty"`

	// RenderedConfigMapRef points to the ConfigMap the controller emitted
	// for this namespace. Unset until the first successful reconcile.
	RenderedConfigMapRef *RenderedConfigMapRef `json:"renderedConfigMapRef,omitempty"`

	// InheritedFromGlobal names the GlobalWAFPolicy whose plugin merge
	// contributed an inherited error (e.g. an Invalid reason on the Accepted
	// condition) to this policy. Nil when the failure is self-authored or
	// no failure is present. UI consumers use this to render a "Inherited
	// from GlobalWAFPolicy/<name>" attribution badge that deep-links to the
	// offending Global resource.
	// +optional
	InheritedFromGlobal *string `json:"inheritedFromGlobal,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=wafp
// +kubebuilder:printcolumn:name="Action",type=string,JSONPath=`.spec.action`
// +kubebuilder:printcolumn:name="Plugins",type=integer,JSONPath=`.status.pluginCount`
// +kubebuilder:printcolumn:name="Validation",type=string,JSONPath=`.status.validation.status`
// +kubebuilder:printcolumn:name="Posture",type=string,JSONPath=`.status.validation.securityPosture`,priority=1
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// WAFPolicy is the Schema for namespace-scoped WAF configuration
type WAFPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WAFPolicySpec   `json:"spec,omitempty"`
	Status WAFPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WAFPolicyList contains a list of WAFPolicy
type WAFPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WAFPolicy `json:"items"`
}

// =============================================================================
// GlobalWAFPlugin - Cluster-scoped custom WAF plugins
// =============================================================================

// GlobalWAFPluginSpec defines a cluster-wide custom WAF plugin
// +kubebuilder:validation:XValidation:rule="(has(self.config) && size(self.config) > 0) || (has(self.before) && size(self.before) > 0) || (has(self.rules) && size(self.rules) > 0) || (has(self.after) && size(self.after) > 0)",message="plugin must define at least one of config, before, rules, after"
type GlobalWAFPluginSpec struct {
	// Description is a human-readable description of what this plugin does
	// +kubebuilder:validation:MaxLength=1024
	Description string `json:"description,omitempty"`

	// Config contains raw SecAction/SecRule directives for plugin configuration.
	// These run first and typically set tx.* variables.
	// Equivalent to CRS *-config.conf files.
	// +kubebuilder:validation:MaxLength=65536
	Config string `json:"config,omitempty"`

	// Before contains raw SecAction/SecRule directives that run BEFORE CRS rules.
	// Use for pre-processing, initialization, or early blocking.
	// Equivalent to CRS *-before.conf files.
	// +kubebuilder:validation:MaxLength=65536
	Before string `json:"before,omitempty"`

	// Rules contains the main SecAction/SecRule directives for this plugin.
	// These run alongside CRS rules. Ordering across plugins follows array
	// position in the consuming policy's spec.plugins[].
	// +kubebuilder:validation:MaxLength=65536
	Rules string `json:"rules,omitempty"`

	// After contains raw SecAction/SecRule directives that run AFTER CRS rules.
	// Use for post-processing, logging, or cleanup.
	// Equivalent to CRS *-after.conf files.
	// +kubebuilder:validation:MaxLength=65536
	After string `json:"after,omitempty"`
}

// GlobalWAFPluginStatus defines the observed state.
//
// Conditions emitted by the controller:
//   - Accepted   (DirectivesValid / InvalidDirectives)
type GlobalWAFPluginStatus struct {
	// Conditions represent the latest available observations
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// RuleCount is the number of SecRule directives in this plugin
	RuleCount int `json:"ruleCount,omitempty"`

	// NamespaceCount is the number of namespaces with at least one WAFPolicy
	// referencing this plugin. Kept as a bounded aggregate (not a
	// ReferencedBy list) because a Global plugin can fan out to many
	// namespaces; a list would churn the status payload as policies come
	// and go. Different denominator from GlobalWAFPolicyStatus.NamespaceCount
	// (which counts namespaces using the policy itself).
	NamespaceCount int `json:"namespaceCount,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,shortName=gwafplugin
// +kubebuilder:printcolumn:name="Description",type=string,JSONPath=`.spec.description`
// +kubebuilder:printcolumn:name="Rules",type=integer,JSONPath=`.status.ruleCount`
// +kubebuilder:printcolumn:name="Namespaces",type=integer,JSONPath=`.status.namespaceCount`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// GlobalWAFPlugin is the Schema for cluster-wide custom WAF plugins
type GlobalWAFPlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalWAFPluginSpec   `json:"spec,omitempty"`
	Status GlobalWAFPluginStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlobalWAFPluginList contains a list of GlobalWAFPlugin
type GlobalWAFPluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalWAFPlugin `json:"items"`
}

// =============================================================================
// WAFPlugin - Namespace-scoped custom WAF plugins
// =============================================================================

// WAFPluginSpec defines a custom WAF plugin with CRS 4-style structure
// +kubebuilder:validation:XValidation:rule="(has(self.config) && size(self.config) > 0) || (has(self.before) && size(self.before) > 0) || (has(self.rules) && size(self.rules) > 0) || (has(self.after) && size(self.after) > 0)",message="plugin must define at least one of config, before, rules, after"
type WAFPluginSpec struct {
	// Description is a human-readable description of what this plugin does
	// +kubebuilder:validation:MaxLength=1024
	Description string `json:"description,omitempty"`

	// Config contains raw SecAction/SecRule directives for plugin configuration.
	// These run first and typically set tx.* variables.
	// Equivalent to CRS *-config.conf files.
	// +kubebuilder:validation:MaxLength=65536
	Config string `json:"config,omitempty"`

	// Before contains raw SecAction/SecRule directives that run BEFORE CRS rules.
	// Use for pre-processing, initialization, or early blocking.
	// Equivalent to CRS *-before.conf files.
	// +kubebuilder:validation:MaxLength=65536
	Before string `json:"before,omitempty"`

	// Rules contains the main SecAction/SecRule directives for this plugin.
	// These run alongside CRS rules. Ordering across plugins follows array
	// position in the consuming policy's spec.plugins[].
	// +kubebuilder:validation:MaxLength=65536
	Rules string `json:"rules,omitempty"`

	// After contains raw SecAction/SecRule directives that run AFTER CRS rules.
	// Use for post-processing, logging, or cleanup.
	// Equivalent to CRS *-after.conf files.
	// +kubebuilder:validation:MaxLength=65536
	After string `json:"after,omitempty"`
}

// WAFPluginStatus defines the observed state.
//
// Conditions emitted by the controller:
//   - Accepted   (DirectivesValid / InvalidDirectives)
type WAFPluginStatus struct {
	// Conditions represent the latest available observations
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// RuleCount is the number of SecRule directives in this plugin
	RuleCount int `json:"ruleCount,omitempty"`

	// ReferencedByCount is the total number of policies referencing this
	// plugin — always the true population size, never bounded by the
	// 50-entry list cap on ReferencedBy. Backs the `Refs` printcolumn so
	// `kubectl get wafplugin` shows the real fan-in even when ReferencedBy
	// is truncated.
	ReferencedByCount int `json:"referencedByCount,omitempty"`

	// ReferencedBy enumerates the policies that name this plugin in their
	// spec.plugins[]. Sorted: namespace asc, name asc, kind asc. Updated
	// every reconcile so client-side filtering is unnecessary. Capped at
	// 50 entries (matching RenderedConfigMapRefs on GlobalWAFPolicy) so a
	// runaway WAFPolicy population can't blow up the status payload; when
	// the cap is hit ReferencedByTruncated is set and the scalar
	// ReferencedByCount above remains authoritative for "how many
	// actually reference me".
	// +optional
	// +kubebuilder:validation:MaxItems=50
	ReferencedBy []PolicyRef `json:"referencedBy,omitempty"`

	// ReferencedByTruncated is true when the consuming-policy set exceeded
	// the 50-entry cap on ReferencedBy. Consumers should fall back to a
	// label-selector list query rather than relying on ReferencedBy alone;
	// ReferencedByCount remains accurate regardless.
	// +optional
	ReferencedByTruncated bool `json:"referencedByTruncated,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=wafplugin
// +kubebuilder:printcolumn:name="Description",type=string,JSONPath=`.spec.description`
// +kubebuilder:printcolumn:name="Rules",type=integer,JSONPath=`.status.ruleCount`
// +kubebuilder:printcolumn:name="Refs",type=integer,JSONPath=`.status.referencedByCount`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// WAFPlugin is the Schema for namespace-scoped custom WAF plugins
type WAFPlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WAFPluginSpec   `json:"spec,omitempty"`
	Status WAFPluginStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WAFPluginList contains a list of WAFPlugin
type WAFPluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WAFPlugin `json:"items"`
}

// =============================================================================
// Validation Policy Types - Shared between Global and Namespace-scoped
// =============================================================================

// ValidationSeverity defines the severity of a validation failure
// +kubebuilder:validation:Enum=info;warning;critical
type ValidationSeverity string

const (
	ValidationSeverityInfo     ValidationSeverity = "info"
	ValidationSeverityWarning  ValidationSeverity = "warning"
	ValidationSeverityCritical ValidationSeverity = "critical"
)

// EnforcementMode defines how validation failures are handled
// +kubebuilder:validation:Enum=Audit;Enforce
type EnforcementMode string

const (
	// EnforcementModeAudit logs warnings but generates EnvoyExtensionPolicy anyway
	EnforcementModeAudit EnforcementMode = "Audit"
	// EnforcementModeEnforce blocks EnvoyExtensionPolicy generation on critical failures
	EnforcementModeEnforce EnforcementMode = "Enforce"
)

// ValidationRule defines a single Rego-based validation rule
type ValidationRule struct {
	// Name is a unique identifier for this rule
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=63
	Name string `json:"name"`

	// Rego is a Rego module that decides violations against the merged WAF
	// configuration. The module must declare `package waf` and produce a
	// `violations` set; an empty set means the rule passed. Each violation
	// may be a plain string message, or an object carrying a `msg` string.
	// The `input` document available to the module:
	//   - input.directives: []string - the final merged directive list
	//   - input.config.action: string - effective action ("Detect" or "Block")
	//   - input.config.crsState: string - CRS state ("Enabled" or "Disabled")
	//   - input.config.paranoiaLevel: int - effective paranoia level (1-4)
	//   - input.source.globalPolicy: string - name of GlobalWAFPolicy
	//   - input.source.namespacePolicy: string - name of WAFPolicy (if any)
	//   - input.source.globalPlugins: []string - names of global plugins applied
	//   - input.source.namespacePlugins: []string - names of namespace plugins applied
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=16384
	Rego string `json:"rego"`

	// Severity determines the impact of this rule failing
	// +kubebuilder:default=warning
	Severity ValidationSeverity `json:"severity,omitempty"`

	// Message is displayed when validation fails
	// +kubebuilder:validation:MaxLength=1024
	Message string `json:"message,omitempty"`
}

// ValidationResult represents the result of a single validation rule
type ValidationResult struct {
	// Rule is the name of the validation rule
	// +kubebuilder:validation:MaxLength=63
	Rule string `json:"rule"`

	// Passed indicates whether the rule passed
	Passed bool `json:"passed"`

	// Severity is the severity level of this rule
	Severity ValidationSeverity `json:"severity,omitempty"`

	// Message is the failure message (if failed)
	// +kubebuilder:validation:MaxLength=1024
	Message string `json:"message,omitempty"`
}

// SecurityPosture represents the overall security posture based on validation
// +kubebuilder:validation:Enum=Compliant;Warning;Degraded;Critical
type SecurityPosture string

const (
	SecurityPostureCompliant SecurityPosture = "Compliant"
	SecurityPostureWarning   SecurityPosture = "Warning"
	SecurityPostureDegraded  SecurityPosture = "Degraded"
	SecurityPostureCritical  SecurityPosture = "Critical"
)

// =============================================================================
// GlobalWAFValidationPolicy - Cluster-scoped validation (CO)
// =============================================================================

// GlobalWAFValidationPolicySpec defines cluster-wide validation rules
type GlobalWAFValidationPolicySpec struct {
	// EnforcementMode determines how validation failures are handled
	// - Audit: Log warnings but generate EnvoyExtensionPolicy anyway
	// - Enforce: Block EnvoyExtensionPolicy generation on critical failures
	// +kubebuilder:default=Audit
	EnforcementMode EnforcementMode `json:"enforcementMode,omitempty"`

	// Rules defines the validation rules to apply
	// +kubebuilder:validation:MinItems=1
	Rules []ValidationRule `json:"rules"`

	// NamespaceSelector limits which namespaces this validation policy applies to.
	// If empty, applies to all namespaces.
	//
	// Standard metav1.LabelSelector semantics: every namespace — including
	// `default` — is matched against the selector's labels, with no implicit
	// special-casing. Since k8s 1.21 every namespace carries an auto-applied
	// `kubernetes.io/metadata.name=<ns-name>` label, so the canonical way to
	// scope down to specific namespaces (including `default`) is a
	// matchExpressions on that label.
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`
}

// ViolatingPolicy identifies a WAF policy that failed validation
type ViolatingPolicy struct {
	// Namespace is the namespace of the violating policy
	// +kubebuilder:validation:MaxLength=253
	Namespace string `json:"namespace"`

	// Name is the name of the WAFPolicy
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`

	// FailedRules lists the rules that this policy violated
	FailedRules []string `json:"failedRules,omitempty"`

	// Severity is the highest severity of the failures
	Severity ValidationSeverity `json:"severity,omitempty"`

	// Action taken: Audited or Rejected
	// +kubebuilder:validation:MaxLength=253
	Action string `json:"action,omitempty"`
}

// GlobalWAFValidationPolicyStatus defines the observed state.
//
// Conditions emitted by the controller:
//   - Ready      (Evaluated)
type GlobalWAFValidationPolicyStatus struct {
	// Conditions represent the latest available observations
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Summary provides a quick overview of validation state
	// +kubebuilder:validation:MaxLength=1024
	Summary string `json:"summary,omitempty"`

	// EvaluatedCount is the number of WAF policies evaluated
	EvaluatedCount int `json:"evaluatedCount,omitempty"`

	// PassingCount is the number of policies passing all rules
	PassingCount int `json:"passingCount,omitempty"`

	// ViolatingPolicies lists all WAF policies that failed validation.
	// Useful for COs to centrally see which namespaces need attention.
	ViolatingPolicies []ViolatingPolicy `json:"violatingPolicies,omitempty"`

	// LastEvaluated is the timestamp of the last evaluation
	LastEvaluated *metav1.Time `json:"lastEvaluated,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,shortName=gwafvp
// +kubebuilder:printcolumn:name="Mode",type=string,JSONPath=`.spec.enforcementMode`
// +kubebuilder:printcolumn:name="Passing",type=integer,JSONPath=`.status.passingCount`
// +kubebuilder:printcolumn:name="Violating",type=integer,JSONPath=`.status.violatingPolicies`
// +kubebuilder:printcolumn:name="Summary",type=string,JSONPath=`.status.summary`,priority=1
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// GlobalWAFValidationPolicy defines cluster-wide validation rules that merged
// WAF configurations must pass. Used by Cluster Operators to enforce security
// requirements across all namespaces.
type GlobalWAFValidationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalWAFValidationPolicySpec   `json:"spec,omitempty"`
	Status GlobalWAFValidationPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlobalWAFValidationPolicyList contains a list of GlobalWAFValidationPolicy
type GlobalWAFValidationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalWAFValidationPolicy `json:"items"`
}

// =============================================================================
// WAFValidationPolicy - Namespace-scoped validation (AO)
// =============================================================================

// WAFValidationPolicySpec defines namespace validation rules.
//
// Namespace WAFValidationPolicies are ADVISORY ONLY: they self-validate an App
// Operator's namespace config and surface failures via status + SecurityPosture
// but never block EnvoyExtensionPolicy generation. The enforcement role belongs
// to the Cluster Operator's GlobalWAFValidationPolicy (which has an
// EnforcementMode); this type intentionally has no EnforcementMode. See the
// approved design (per-route/namespaced WAF, designs#25 §01/§02). A future
// CO-delegated self-enforcement model would be added separately (RFE). EV-6386.
type WAFValidationPolicySpec struct {
	// Rules defines the validation rules to apply
	// +kubebuilder:validation:MinItems=1
	Rules []ValidationRule `json:"rules"`
}

// WAFValidationPolicyStatus defines the observed state.
// No conditions are emitted by the controller for this type today; the Conditions
// field is reserved for future use.
type WAFValidationPolicyStatus struct {
	// Conditions represent the latest available observations
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// SecurityPosture is the overall posture based on validation results
	SecurityPosture SecurityPosture `json:"securityPosture,omitempty"`

	// ValidationResults contains results for each rule
	ValidationResults []ValidationResult `json:"validationResults,omitempty"`

	// LastEvaluated is the timestamp of the last evaluation
	LastEvaluated *metav1.Time `json:"lastEvaluated,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=wafvp
// +kubebuilder:printcolumn:name="Posture",type=string,JSONPath=`.status.securityPosture`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// WAFValidationPolicy defines namespace-scoped validation rules for Application
// Operators to self-validate their WAF configuration. Useful for catching
// mistakes like accidentally removing required plugins.
type WAFValidationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WAFValidationPolicySpec   `json:"spec,omitempty"`
	Status WAFValidationPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WAFValidationPolicyList contains a list of WAFValidationPolicy
type WAFValidationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WAFValidationPolicy `json:"items"`
}
