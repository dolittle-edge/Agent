/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package configuring

// ICanConfigureNode defined a system that is capable of configuring a node
type ICanConfigureNode interface {
	Name() string
	Configure(interface{}) error
	SetDebug(bool)
}
