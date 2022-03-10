// Copyright (c) 2019-2022 Tigera, Inc. All rights reserved.

package v3

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	// Add non-generated conversion functions
	err := scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "NetworkPolicy"},
		func(label, value string) (string, string, error) {
			switch label {
			case "spec.tier", "metadata.name", "metadata.namespace":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label not supported: %s", label)
			}
		},
	)
	if err != nil {
		return err
	}

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "GlobalNetworkPolicy"},
		func(label, value string) (string, string, error) {
			switch label {
			case "spec.tier", "metadata.name":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label not supported: %s", label)
			}
		},
	)
	if err != nil {
		return err
	}

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "GlobalNetworkSet"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "NetworkSet"},
		func(label, value string) (string, string, error) {
			switch label {
			case "metadata.name", "metadata.namespace":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label not supported: %s", label)
			}
		},
	)
	if err != nil {
		return err
	}

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "HostEndpoint"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "IPPool"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "BGPConfiguration"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "BGPPeer"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "Profile"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "FelixConfiguration"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "KubeControllersConfiguration"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "ClusterInformation"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "GlobalAlert"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "GlobalAlertTemplate"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "GlobalReport"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "GlobalReportType"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "GlobalThreatFeed"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "LicenseKey"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "ManagedCluster"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "PacketCapture"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "RemoteClusterConfiguration"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "StagedGlobalNetworkPolicy"},
		func(label, value string) (string, string, error) {
			switch label {
			case "spec.tier", "metadata.name":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label not supported: %s", label)
			}
		},
	)
	if err != nil {
		return err
	}

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "StagedKubernetesNetworkPolicy"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "StagedNetworkPolicy"},
		func(label, value string) (string, string, error) {
			switch label {
			case "spec.tier", "metadata.name", "metadata.namespace":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label not supported: %s", label)
			}
		},
	)
	if err != nil {
		return err
	}

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "Tier"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "UISettingsGroup"},
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

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "UISettings"},
		func(label, value string) (string, string, error) {
			switch label {
			case "spec.group", "metadata.name":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label not supported: %s", label)
			}
		},
	)
	if err != nil {
		return err
	}

	err = scheme.AddFieldLabelConversionFunc(schema.GroupVersionKind{"projectcalico.org", "v3", "AlertException"},
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
