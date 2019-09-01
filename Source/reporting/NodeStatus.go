/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package reporting

import (
	"agent/provisioning"
)

// NodeStatus represents a node within a location
type NodeStatus struct {
	provisioning.Node
	Configuration interface{} `json:"Configuration,omitempty"`
	Metrics       map[string]float32
	Infos         map[string]string
}
