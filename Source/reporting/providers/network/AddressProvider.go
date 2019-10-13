/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package network

import (
	"agent/log"
	. "agent/reporting"
	"fmt"
	"net"
	"strings"
)

// AddressProvider provides information about the nodes current IP addresses
type AddressProvider struct {
	debug bool
}

// NewAddressProvider creates a new instance of AddressProvider
func NewAddressProvider() *AddressProvider {
	return new(AddressProvider)
}

// SetDebug sets the debugging output flag
func (provider *AddressProvider) SetDebug(debug bool) {
	provider.debug = debug
}

// Provide the memory telemetry
func (provider *AddressProvider) Provide() (_ []NodeMetric, infos []NodeInfo) {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Errorln("Error retrieving network interfaces:", err)
		return
	}

	for _, iface := range interfaces {
		basetype := fmt.Sprintf("Network.Interface.%d.", iface.Index)
		infos = append(infos,
			NodeInfo{
				Type:  basetype + "Name",
				Value: iface.Name,
			},
		)

		if provider.debug {
			log.Debugf("Found network interface %s\n", iface.Name)
		}

		addrs, err := iface.Addrs()
		if err != nil {
			log.Warningf("Error retrieving addresses for network interface %s: %v", iface.Name, err)
		} else {
			addresses := []string{}
			for _, addr := range addrs {
				addresses = append(addresses, addr.String())

				if provider.debug {
					log.Debugf("Found address %s for interface %s\n", addr.String(), iface.Name)
				}
			}

			infos = append(infos,
				NodeInfo{
					Type:  basetype + "Addresses",
					Value: strings.Join(addresses, ","),
				},
			)
		}
	}

	return
}
