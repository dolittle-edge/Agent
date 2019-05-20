/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package reporting

import (
	"os"
	"syscall"
)

// DiskUsageTelemetryProvider provides disk usage telemetry
type DiskUsageTelemetryProvider struct{}

// Provide the disk usage telemetry
func (provider DiskUsageTelemetryProvider) Provide() []*TelemetrySample {
	samples := []*TelemetrySample{}

	var stat syscall.Statfs_t
	wd, _ := os.Getwd()

	syscall.Statfs(wd, &stat)

	totalBytes := stat.Blocks * uint64(stat.Bsize)

	bytes := stat.Bfree * uint64(stat.Bsize)

	diskUsageSample := new(TelemetrySample)
	diskUsageSample.Type = "DiskUsage"
	diskUsageSample.Value = 100 - ((float32(bytes) / float32(totalBytes)) * 100)
	samples = append(samples, diskUsageSample)

	return samples
}
