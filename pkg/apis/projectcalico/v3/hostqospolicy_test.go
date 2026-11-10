// Copyright (c) 2026 Tigera, Inc. All rights reserved.

package v3_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	apiv3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/resource"
)

func bps(v int64) resource.Quantity { return *resource.NewQuantity(v, resource.DecimalSI) }

// WithinSafetyMargin is the pure helper backing the HostQoSPolicy reservation
// safety-margin admission check (sum(reservations) <= linkCapacity * margin).
var _ = DescribeTable("HostQoSPolicy WithinSafetyMargin",
	func(sum, linkCap resource.Quantity, margin float64, expected bool) {
		Expect(apiv3.WithinSafetyMargin(sum, linkCap, margin)).To(Equal(expected))
	},
	// Default margin 0.75 on a 10G link: cap is 7.5G.
	Entry("exactly at the 0.75 cap is accepted", bps(7_500_000_000), bps(10_000_000_000), 0.75, true),
	Entry("one over the 0.75 cap is rejected", bps(7_500_000_001), bps(10_000_000_000), 0.75, false),
	Entry("well under the 0.75 cap is accepted", bps(1_000_000_000), bps(10_000_000_000), 0.75, true),
	// margin 1.0 degenerates to the hard cap (sum <= linkCapacity).
	Entry("margin 1.0 accepts sum == linkCapacity", bps(10_000_000_000), bps(10_000_000_000), 1.0, true),
	Entry("margin 1.0 rejects sum > linkCapacity", bps(10_000_000_001), bps(10_000_000_000), 1.0, false),
	// Small margin.
	Entry("margin 0.5 rejects sum above half", bps(6_000_000_000), bps(10_000_000_000), 0.5, false),
	Entry("margin 0.5 accepts sum at half", bps(5_000_000_000), bps(10_000_000_000), 0.5, true),
	// Zero reservations always fit.
	Entry("zero reservations always fit", bps(0), bps(10_000_000_000), 0.75, true),
	// Sums beyond int64 must still be compared honestly; Quantity.Value() would report 0.
	Entry("sum overflowing int64 is rejected",
		resource.MustParse("15E"), resource.MustParse("5E"), 0.75, false),
	Entry("capacity overflowing int64 still caps",
		resource.MustParse("15E"), resource.MustParse("100E"), 0.75, true),
)
