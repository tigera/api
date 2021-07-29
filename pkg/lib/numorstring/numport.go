// Copyright (c) 2021 Tigera, Inc. All rights reserved.
//
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

package numorstring

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-openapi/spec"
	openapi "k8s.io/kube-openapi/pkg/common"
)

// NumPort represents either a range of numeric ports or a single port.
//
//     - For a port range, set MinPort and MaxPort to the (inclusive) port numbers.
//     - For a single port, set MinPort = MaxPort.
type NumPort struct {
	MinPort uint16 `json:"minPort,omitempty"`
	MaxPort uint16 `json:"maxPort,omitempty"`
}

// SingleNumPort creates a NUmPort struct representing a single port.
func SingleNumPort(port uint16) NumPort {
	return NumPort{MinPort: port, MaxPort: port}
}

// NumPortFromRange creates a Port struct representing a range of ports.
func NumPortFromRange(minPort, maxPort uint16) (NumPort, error) {
	port := NumPort{MinPort: minPort, MaxPort: maxPort}
	if minPort > maxPort {
		msg := fmt.Sprintf("minimum port number (%d) is greater than maximum port number (%d) in port range", minPort, maxPort)
		return port, errors.New(msg)
	}
	return port, nil
}

// NumPortFromString creates a Port struct from its string representation.  A port
// may either be single value "1234", a range of values "100:200".
func NumPortFromString(s string) (NumPort, error) {
	if allDigits.MatchString(s) {
		// Port is all digits, it should parse as a single port.
		num, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			msg := fmt.Sprintf("invalid port format (%s)", s)
			return NumPort{}, errors.New(msg)
		}
		return SingleNumPort(uint16(num)), nil
	}

	if groups := portRange.FindStringSubmatch(s); len(groups) > 0 {
		// Port matches <digits>:<digits>, it should parse as a range of ports.
		if pmin, err := strconv.ParseUint(groups[1], 10, 16); err != nil {
			msg := fmt.Sprintf("invalid minimum port number in range (%s)", s)
			return NumPort{}, errors.New(msg)
		} else if pmax, err := strconv.ParseUint(groups[2], 10, 16); err != nil {
			msg := fmt.Sprintf("invalid maximum port number in range (%s)", s)
			return NumPort{}, errors.New(msg)
		} else {
			return NumPortFromRange(uint16(pmin), uint16(pmax))
		}
	}

	msg := fmt.Sprintf("invalid value for a numerical port (%s)", s)
	return NumPort{}, errors.New(msg)
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (p *NumPort) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}

		if v, err := NumPortFromString(s); err != nil {
			return err
		} else {
			*p = v
			return nil
		}
	}

	// It's not a string, it must be a single int.
	var i uint16
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	v := SingleNumPort(i)
	*p = v
	return nil
}

// MarshalJSON implements the json.Marshaller interface.
func (p NumPort) MarshalJSON() ([]byte, error) {
	if p.MinPort == p.MaxPort {
		return json.Marshal(p.MinPort)
	} else {
		return json.Marshal(p.String())
	}
}

// String returns the string value.  If the min and max port are the same
// this returns a single string representation of the port number, otherwise
// if returns a colon separated range of ports.
func (p NumPort) String() string {
	if p.MinPort == p.MaxPort {
		return strconv.FormatUint(uint64(p.MinPort), 10)
	} else {
		return fmt.Sprintf("%d:%d", p.MinPort, p.MaxPort)
	}
}

func (_ NumPort) OpenAPIDefinition() openapi.OpenAPIDefinition {
	return openapi.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:   []string{"string"},
				Format: "int-or-string",
			},
		},
	}
}
