// Copyright (c) 2026 Tigera, Inc. All rights reserved.

package v3_test

import (
	"encoding/json"
	"testing"

	appv3 "github.com/tigera/api/pkg/apis/applicationlayer.projectcalico.org/v3"
)

func ptrBool(b bool) *bool { return &b }

func TestCoreRuleSetConfig_EffectiveState(t *testing.T) {
	cases := []struct {
		name string
		c    appv3.CoreRuleSetConfig
		want appv3.CoreRuleSetState
	}{
		{"state-set-enabled", appv3.CoreRuleSetConfig{State: appv3.CoreRuleSetStateEnabled}, appv3.CoreRuleSetStateEnabled},
		{"state-set-disabled", appv3.CoreRuleSetConfig{State: appv3.CoreRuleSetStateDisabled}, appv3.CoreRuleSetStateDisabled},
		{"only-bool-true", appv3.CoreRuleSetConfig{Enabled: ptrBool(true)}, appv3.CoreRuleSetStateEnabled},
		{"only-bool-false", appv3.CoreRuleSetConfig{Enabled: ptrBool(false)}, appv3.CoreRuleSetStateDisabled},
		{"both-state-wins", appv3.CoreRuleSetConfig{State: appv3.CoreRuleSetStateDisabled, Enabled: ptrBool(true)}, appv3.CoreRuleSetStateDisabled},
		{"unset-defaults-enabled", appv3.CoreRuleSetConfig{}, appv3.CoreRuleSetStateEnabled},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.c.EffectiveState(); got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestPluginRefJSONRoundtrip(t *testing.T) {
	// Backward-compat: existing CRs that omit Kind serialize without the
	// field; on round-trip the zero value is recovered. The kubebuilder
	// default applies at admission, so a real cluster sees Kind populated;
	// this test covers the in-process Go semantics.
	cases := []struct {
		name string
		ref  appv3.PluginRef
		want string
	}{
		{
			name: "omitted-kind",
			ref:  appv3.PluginRef{Name: "p"},
			want: `{"name":"p"}`,
		},
		{
			name: "explicit-namespace-kind",
			ref:  appv3.PluginRef{Kind: "WAFPlugin", Name: "p"},
			want: `{"kind":"WAFPlugin","name":"p"}`,
		},
		{
			name: "explicit-global-kind",
			ref:  appv3.PluginRef{Kind: "GlobalWAFPlugin", Name: "p"},
			want: `{"kind":"GlobalWAFPlugin","name":"p"}`,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(tc.ref)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if string(got) != tc.want {
				t.Errorf("marshal: got %s, want %s", string(got), tc.want)
			}
			var back appv3.PluginRef
			if err := json.Unmarshal(got, &back); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if back != tc.ref {
				t.Errorf("roundtrip: got %+v, want %+v", back, tc.ref)
			}
		})
	}
}

func TestWAFPolicyStatusHasRenderedConfigMapRef(t *testing.T) {
	var s appv3.WAFPolicyStatus
	s.RenderedConfigMapRef = &appv3.RenderedConfigMapRef{
		Name:      "tigera-waf-config",
		Namespace: "team-a",
	}
	if s.RenderedConfigMapRef.Namespace != "team-a" {
		t.Fatalf("assignment did not stick")
	}
}

func TestGlobalWAFPolicyStatusHasRefList(t *testing.T) {
	var s appv3.GlobalWAFPolicyStatus
	s.RenderedConfigMapRefs = []appv3.RenderedConfigMapRef{
		{Name: "tigera-waf-config", Namespace: "team-a"},
		{Name: "tigera-waf-config", Namespace: "team-b"},
	}
	if len(s.RenderedConfigMapRefs) != 2 {
		t.Fatalf("expected 2 refs, got %d", len(s.RenderedConfigMapRefs))
	}
	s.RenderedConfigMapRefsTruncated = true
	if !s.RenderedConfigMapRefsTruncated {
		t.Fatal("truncation flag not set")
	}
}

func TestRenderedConfigMapRefJSON(t *testing.T) {
	ref := appv3.RenderedConfigMapRef{
		Name:            "tigera-waf-config",
		Namespace:       "team-a",
		ResourceVersion: "42891",
	}
	b, err := json.Marshal(ref)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	want := `{"name":"tigera-waf-config","namespace":"team-a","resourceVersion":"42891"}`
	if string(b) != want {
		t.Fatalf("got %s want %s", b, want)
	}
}
