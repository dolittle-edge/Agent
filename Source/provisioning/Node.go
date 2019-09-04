/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package provisioning

// Node represents an edge node within a location
type Node struct {
	LocationId    string
	NodeId        string
	Configuration map[string]interface{}
	isValid       bool
}

// IsValid returns whether the node configuration is valid or not
func (n Node) IsValid() bool {
	return n.isValid
}
