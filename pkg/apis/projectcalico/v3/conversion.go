// Copyright (c) 2019 Tigera, Inc. All rights reserved.

package v3

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	err := scheme.AddFieldLabelConversionFunc("projectcalico.org/v3", "GlobalReportType",
		func(label, value string) (string, string, error) {
			switch label {
			case "metadata.name":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label not supported: %s", label)
			}
		},
	)
	if err != nil {
		return err
	}

	return nil
}
