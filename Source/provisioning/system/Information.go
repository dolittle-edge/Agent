/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package system

import (
	"fmt"
	"io"
)

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

// Print prints the system information in a pretty format to the provided writer
func (info Information) Print(writer io.Writer) {
	fmt.Fprintf(writer, "Manufacturer: %s\n", info.Manufacturer)
	fmt.Fprintf(writer, "Family:       %s\n", info.Family)
	fmt.Fprintf(writer, "ProductName:  %s\n", info.ProductName)
	fmt.Fprintf(writer, "Version:      %s\n", info.Version)
	fmt.Fprintf(writer, "SKUNumber:    %s\n", info.SKUNumber)
	fmt.Fprintf(writer, "SerialNumber: %s\n", info.SerialNumber)
	fmt.Fprintf(writer, "UUID:         %s\n", info.UUID)
}
