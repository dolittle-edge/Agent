/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package memory

import (
	"agent/log"
	. "agent/reporting"

	sigar "github.com/cloudfoundry/gosigar"
)

// TelemetryProvider provides memory telemetry
type TelemetryProvider struct {
	debug bool
}

// NewTelemetryProvider creates a new instance of TelemetryProvider
func NewTelemetryProvider() *TelemetryProvider {
	return new(TelemetryProvider)
}

// SetDebug sets the debugging output flag
func (provider *TelemetryProvider) SetDebug(debug bool) {
	provider.debug = debug
}

// Provide the memory telemetry
func (provider *TelemetryProvider) Provide() (samples []NodeMetric, _ []NodeInfo) {

	mem := sigar.Mem{}
	if err := mem.Get(); err != nil {
		log.Errorln("Failed to get memory information:", err)
	} else {
		if provider.debug {
			log.Debugf("Got memory information total:%v free:%v\n", mem.Total, mem.Free)
			log.Debugf("Got actual memory information total:%v free:%v\n", mem.Total, mem.ActualFree)
		}
		samples = append(samples,
			NodeMetric{
				Type:  "Memory",
				Value: 100 - ((float64(mem.Free) / float64(mem.Total)) * 100),
			}, NodeMetric{
				Type:  "ActualMemory",
				Value: 100 - (float64(mem.ActualFree) / float64(mem.Total) * 100),
			})
	}

	swap := sigar.Swap{}
	if err := swap.Get(); err != nil {
		log.Errorln("Failed to get swap information:", err)
	} else {
		if provider.debug {
			log.Debugf("Got swap information total:%v free:%v\n", swap.Total, swap.Free)
		}
		if swap.Total > 0 {
			samples = append(samples,
				NodeMetric{
					Type:  "SwapMemory",
					Value: 100 - (float64(swap.Free) / float64(swap.Total) * 100),
				})
		}
	}
	return
}
