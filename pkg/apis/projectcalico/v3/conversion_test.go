// Copyright (c) 2026 Tigera, Inc. All rights reserved.

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

package v3_test

import (
	"io/fs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tigera/api/config/crd"
	apiv3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

// The aggregated apiserver rejects any field label it hasn't been told about, so
// this whitelist has to stay in step with UISettingsToSelectableFields and with
// the CRD's selectableFields.
var _ = Describe("UISettings field label conversion", func() {
	gvk := schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: apiv3.KindUISettings}

	var scheme *runtime.Scheme

	BeforeEach(func() {
		scheme = runtime.NewScheme()
		Expect(apiv3.AddToScheme(scheme)).To(Succeed())
	})

	for _, label := range []string{"metadata.name", "spec.group", "spec.user"} {
		It("accepts "+label, func() {
			gotLabel, gotValue, err := scheme.ConvertFieldLabel(gvk, label, "some-value")
			Expect(err).NotTo(HaveOccurred())
			Expect(gotLabel).To(Equal(label))
			Expect(gotValue).To(Equal("some-value"))
		})
	}

	It("rejects a label that isn't a selectable field", func() {
		_, _, err := scheme.ConvertFieldLabel(gvk, "spec.description", "whatever")
		Expect(err).To(MatchError(ContainSubstring("field label not supported: spec.description")))
	})
})

// Clients look up the HostEndpoints of a node by spec.node, so the field has to be
// selectable whether the API is served by the aggregated apiserver or by the v3 CRDs.
var _ = Describe("HostEndpoint spec.node field selector", func() {
	It("is accepted by the aggregated apiserver", func() {
		scheme := runtime.NewScheme()
		Expect(apiv3.AddToScheme(scheme)).To(Succeed())

		gvk := schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: apiv3.KindHostEndpoint}
		label, value, err := scheme.ConvertFieldLabel(gvk, "spec.node", "node-1")
		Expect(err).NotTo(HaveOccurred())
		Expect(label).To(Equal("spec.node"))
		Expect(value).To(Equal("node-1"))
	})

	It("is a selectable field of the CRD", func() {
		raw, err := fs.ReadFile(crd.FS(), "projectcalico.org_hostendpoints.yaml")
		Expect(err).NotTo(HaveOccurred())

		var def apiextensionsv1.CustomResourceDefinition
		Expect(yaml.Unmarshal(raw, &def)).To(Succeed())
		Expect(def.Spec.Versions).NotTo(BeEmpty())
		for _, v := range def.Spec.Versions {
			Expect(v.SelectableFields).To(ContainElement(apiextensionsv1.SelectableField{JSONPath: ".spec.node"}), "version %s", v.Name)
		}
	})
})
