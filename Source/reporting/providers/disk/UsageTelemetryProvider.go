/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package disk

import (
	"agent/log"
	. "agent/reporting"

	"github.com/ringtail/go-statfs"
)

// UsageTelemetryProvider provides disk usage telemetry
type UsageTelemetryProvider struct{}

// NewUsageTelemetryProvider creates a new instance of UsageTelemetryProvider
func NewUsageTelemetryProvider() *UsageTelemetryProvider {
	return new(UsageTelemetryProvider)
}

// Provide the disk usage telemetry
func (provider *UsageTelemetryProvider) Provide() ([]NodeMetric, []NodeInfo) {
	diskUsage, err := statfs.GetDiskInfo("/")
	if err != nil {
		log.Errorln("Failed to get disk info: ", err)
		return nil, nil
	}

	diskUsageMetric := NodeMetric{
		Type:  "DiskUsage",
		Value: 100 - ((float32(diskUsage.Capacity-diskUsage.Usage) / float32(diskUsage.Capacity)) * 100),
	}
	fileUsageMetric := NodeMetric{
		Type:  "FileUsage",
		Value: 100 - ((float32(diskUsage.Inodes-diskUsage.InodesUsed) / float32(diskUsage.Inodes)) * 100),
	}

	return []NodeMetric{diskUsageMetric, fileUsageMetric}, nil
}
