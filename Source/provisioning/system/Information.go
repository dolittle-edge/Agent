/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package system

// Information represents system information coded in the BIOS of the node
type Information struct {
	Manufacturer string
	Family       string
	ProductName  string
	Version      string
	SKUNumber    string
	SerialNumber string
	UUID         string
}
