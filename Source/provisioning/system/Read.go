/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package system

import (
	"errors"
	"fmt"

	"github.com/digitalocean/go-smbios/smbios"
)

// ReadSystemInformation reads system information from the BIOS if available
func ReadSystemInformation() (Information, error) {
	reader, _, err := smbios.Stream()
	if err != nil {
		return Information{}, err
	}
	defer reader.Close()

	structures, err := smbios.NewDecoder(reader).Decode()
	if err != nil {
		return Information{}, err
	}

	for _, structure := range structures {
		if structure.Header.Type == 1 {
			return decodeSystemInformation(structure.Formatted, structure.Strings), nil
		}
	}

	return Information{}, errors.New("Did not find SystemInformation structure")
}

func decodeSystemInformation(data []byte, strings []string) Information {
	uuid := ""
	if len(data) > 8 {
		uuid = convertUUID(data[4:20])
	}
	return Information{
		Manufacturer: getStringFromPointer(0, data, strings),
		Family:       getStringFromPointer(22, data, strings),
		ProductName:  getStringFromPointer(1, data, strings),
		Version:      getStringFromPointer(2, data, strings),
		SKUNumber:    getStringFromPointer(21, data, strings),
		SerialNumber: getStringFromPointer(3, data, strings),
		UUID:         uuid,
	}
}

func getStringFromPointer(location int, data []byte, strings []string) string {
	if location >= len(data) {
		return ""
	}
	pointer := data[location]
	if pointer < 1 || int(pointer) > len(strings) {
		return ""
	}
	return strings[pointer-1]
}

func convertUUID(data []byte) string {
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%x-%x", data[3], data[2], data[1], data[0], data[5], data[4], data[7], data[6], data[8:10], data[10:16])
}
