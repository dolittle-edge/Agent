/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package reporting

// Node Represents a node within a location
type Node struct {
	LocationId string
	NodeId     string
	State      map[string]float32
}
