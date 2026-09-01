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
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	apiv3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These tests guard the RBACSync field added to KubeControllersConfiguration:
// the generated DeepCopy must produce an independent copy of the optional
// stanza and its optional sub-field, so mutating one cannot affect the other.
var _ = Describe("KubeControllersConfiguration RBACSync deep-copy", func() {
	It("round-trips with a nil RBACSync stanza", func() {
		src := apiv3.KubeControllersConfigurationSpec{
			Controllers: apiv3.ControllersConfig{},
		}
		dst := src.DeepCopy()
		Expect(dst.Controllers.RBACSync).To(BeNil())
	})

	It("round-trips with a non-nil stanza and nil ReconcilerPeriod", func() {
		src := apiv3.KubeControllersConfigurationSpec{
			Controllers: apiv3.ControllersConfig{
				RBACSync: &apiv3.RBACSyncControllerConfig{},
			},
		}
		dst := src.DeepCopy()
		Expect(dst.Controllers.RBACSync).NotTo(BeNil())
		Expect(dst.Controllers.RBACSync).NotTo(BeIdenticalTo(src.Controllers.RBACSync),
			"DeepCopy must allocate a new RBACSyncControllerConfig")
		Expect(dst.Controllers.RBACSync.ReconcilerPeriod).To(BeNil())
	})

	It("round-trips with a non-nil ReconcilerPeriod and the periods are independent", func() {
		src := apiv3.KubeControllersConfigurationSpec{
			Controllers: apiv3.ControllersConfig{
				RBACSync: &apiv3.RBACSyncControllerConfig{
					ReconcilerPeriod: &metav1.Duration{Duration: 47 * time.Second},
				},
			},
		}
		dst := src.DeepCopy()
		Expect(dst.Controllers.RBACSync).NotTo(BeNil())
		Expect(dst.Controllers.RBACSync.ReconcilerPeriod).NotTo(BeNil())
		Expect(dst.Controllers.RBACSync.ReconcilerPeriod).NotTo(BeIdenticalTo(src.Controllers.RBACSync.ReconcilerPeriod),
			"DeepCopy must allocate a new ReconcilerPeriod pointer")
		Expect(dst.Controllers.RBACSync.ReconcilerPeriod.Duration).To(Equal(47 * time.Second))

		// Mutate the source after the copy; the destination must be unaffected.
		src.Controllers.RBACSync.ReconcilerPeriod.Duration = 99 * time.Second
		Expect(dst.Controllers.RBACSync.ReconcilerPeriod.Duration).To(Equal(47 * time.Second))
	})
})
