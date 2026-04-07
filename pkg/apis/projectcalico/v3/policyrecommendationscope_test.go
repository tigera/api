// Copyright (c) 2026 Tigera, Inc. All rights reserved.

package v3_test

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPolicyRecommendationScopeHostEndpointSpecSerialization(t *testing.T) {
	g := NewGomegaWithT(t)

	scope := v3.PolicyRecommendationScope{
		TypeMeta: metav1.TypeMeta{
			Kind:       v3.KindPolicyRecommendationScope,
			APIVersion: v3.GroupVersionCurrent,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "default",
		},
		Spec: v3.PolicyRecommendationScopeSpec{
			Interval: &metav1.Duration{Duration: 150 * time.Second},
			NamespaceSpec: &v3.PolicyRecommendationScopeNamespaceSpec{
				RecStatus: v3.PolicyRecommendationEnabled,
				Selector:  "all()",
			},
			HostEndpointSpec: &v3.PolicyRecommendationScopeHostEndpointSpec{
				RecommendationStatus: v3.PolicyRecommendationEnabled,
				Selector:             "hostendpoint.projectcalico.org/type == 'nonclusterhost'",
				TierName:             "hostendpoint-isolation",
			},
		},
	}

	// Marshal to JSON.
	data, err := json.Marshal(scope)
	g.Expect(err).NotTo(HaveOccurred())

	// Verify JSON contains the hostEndpointSpec fields.
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	g.Expect(err).NotTo(HaveOccurred())
	spec := raw["spec"].(map[string]interface{})
	g.Expect(spec).To(HaveKey("hostEndpointSpec"))
	hostEndpointSpec := spec["hostEndpointSpec"].(map[string]interface{})
	g.Expect(hostEndpointSpec["recommendationStatus"]).To(Equal("Enabled"))
	g.Expect(hostEndpointSpec["selector"]).To(Equal("hostendpoint.projectcalico.org/type == 'nonclusterhost'"))
	g.Expect(hostEndpointSpec["tierName"]).To(Equal("hostendpoint-isolation"))

	// Unmarshal back and verify round-trip.
	var decoded v3.PolicyRecommendationScope
	err = json.Unmarshal(data, &decoded)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(decoded.Spec.HostEndpointSpec).NotTo(BeNil())
	g.Expect(decoded.Spec.HostEndpointSpec.RecommendationStatus).To(Equal(v3.PolicyRecommendationEnabled))
	g.Expect(decoded.Spec.HostEndpointSpec.Selector).To(Equal("hostendpoint.projectcalico.org/type == 'nonclusterhost'"))
	g.Expect(decoded.Spec.HostEndpointSpec.TierName).To(Equal("hostendpoint-isolation"))
}

func TestPolicyRecommendationScopeHostEndpointSpecOmittedWhenNil(t *testing.T) {
	g := NewGomegaWithT(t)

	scope := v3.PolicyRecommendationScope{
		Spec: v3.PolicyRecommendationScopeSpec{
			NamespaceSpec: &v3.PolicyRecommendationScopeNamespaceSpec{
				Selector: "all()",
			},
		},
	}

	data, err := json.Marshal(scope)
	g.Expect(err).NotTo(HaveOccurred())

	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	g.Expect(err).NotTo(HaveOccurred())

	spec := raw["spec"].(map[string]interface{})
	// hostEndpointSpec is a pointer with omitempty, so a nil value is omitted from JSON.
	g.Expect(spec).NotTo(HaveKey("hostEndpointSpec"))
}

func TestPolicyRecommendationScopeHostEndpointSpecDisabledStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	scope := v3.PolicyRecommendationScope{
		Spec: v3.PolicyRecommendationScopeSpec{
			HostEndpointSpec: &v3.PolicyRecommendationScopeHostEndpointSpec{
				RecommendationStatus: v3.PolicyRecommendationDisabled,
				Selector:             "all()",
			},
		},
	}

	data, err := json.Marshal(scope)
	g.Expect(err).NotTo(HaveOccurred())

	var decoded v3.PolicyRecommendationScope
	err = json.Unmarshal(data, &decoded)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(decoded.Spec.HostEndpointSpec).NotTo(BeNil())
	g.Expect(decoded.Spec.HostEndpointSpec.RecommendationStatus).To(Equal(v3.PolicyRecommendationDisabled))
}
