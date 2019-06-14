/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package reporting

import (
	"fmt"
	statfs "github.com/ringtail/go-statfs"
)

// DiskUsageTelemetryProvider provides disk usage telemetry
type DiskUsageTelemetryProvider struct{}

// Provide the disk usage telemetry
func (provider DiskUsageTelemetryProvider) Provide() []*TelemetrySample {
	samples := []*TelemetrySample{}

	diskUsage, err := statfs.GetDiskInfo("/")
	if err != nil {
		fmt.Printf("Failed to get disk info,because of %s", err.Error())
	}

	diskUsageSample := new(TelemetrySample)
	diskUsageSample.Type = "DiskUsage"
	diskUsageSample.Value = 100-((float32(diskUsage.Capacity - diskUsage.Usage) / float32(diskUsage.Capacity))*100)
	samples = append(samples, diskUsageSample)

	fileUsageSample := new(TelemetrySample)
	fileUsageSample.Type = "FileUsage"
	fileUsageSample.Value = 100-((float32(diskUsage.Inodes - diskUsage.InodesUsed) / float32(diskUsage.Inodes))*100)
	samples = append(samples, fileUsageSample)

	return samples
}
