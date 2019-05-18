/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package main

import (
	sigar "github.com/cloudfoundry/gosigar"
)

// MemoryTelemetryProvider Provides memory telemetry
type MemoryTelemetryProvider struct{}

// Provide the memory telemetry
func (provider MemoryTelemetryProvider) Provide() []*TelemetrySample {
	samples := []*TelemetrySample{}

	mem := sigar.Mem{}
	swap := sigar.Swap{}

	mem.Get()
	swap.Get()

	memorySample := new(TelemetrySample)
	memorySample.Type = "Memory"
	memorySample.Value = 100 - ((float32(mem.Free) / float32(mem.Total)) * 100)
	samples = append(samples, memorySample)

	actualMemorySample := new(TelemetrySample)
	actualMemorySample.Type = "ActualMemory"
	actualMemorySample.Value = 100 - (float32(mem.ActualFree) / float32(mem.Total) * 100)
	samples = append(samples, actualMemorySample)

	swapSample := new(TelemetrySample)
	swapSample.Type = "SwapMemory"
	swapSample.Value = 100 - (float32(swap.Free) / float32(swap.Total) * 100)
	samples = append(samples, swapSample)

	return samples
}
