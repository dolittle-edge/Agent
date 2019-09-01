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
type TelemetryProvider struct{}

// NewTelemetryProvider creates a new instance of TelemetryProvider
func NewTelemetryProvider() *TelemetryProvider {
	return new(TelemetryProvider)
}

// Provide the memory telemetry
func (provider *TelemetryProvider) Provide() (samples []NodeMetric, _ []NodeInfo) {

	mem := sigar.Mem{}
	if err := mem.Get(); err != nil {
		log.Errorln("Failed to get memory information:", err)
	} else {
		samples = append(samples,
			NodeMetric{
				Type:  "Memory",
				Value: 100 - ((float32(mem.Free) / float32(mem.Total)) * 100),
			}, NodeMetric{
				Type:  "ActualMemory",
				Value: 100 - (float32(mem.ActualFree) / float32(mem.Total) * 100),
			})
	}

	swap := sigar.Swap{}
	if err := swap.Get(); err != nil {
		log.Errorln("Failed to get swap information:", err)
	} else {
		samples = append(samples,
			NodeMetric{
				Type:  "SwapMemory",
				Value: 100 - (float32(swap.Free) / float32(swap.Total) * 100),
			})
	}
	return
}
