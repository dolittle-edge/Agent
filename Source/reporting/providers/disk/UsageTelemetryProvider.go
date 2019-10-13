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
type UsageTelemetryProvider struct {
	debug bool
}

// NewUsageTelemetryProvider creates a new instance of UsageTelemetryProvider
func NewUsageTelemetryProvider() *UsageTelemetryProvider {
	return new(UsageTelemetryProvider)
}

// SetDebug sets the debugging output flag
func (provider *UsageTelemetryProvider) SetDebug(debug bool) {
	provider.debug = debug
}

// Provide the disk usage telemetry
func (provider *UsageTelemetryProvider) Provide() ([]NodeMetric, []NodeInfo) {
	diskUsage, err := statfs.GetDiskInfo("/")
	if err != nil {
		log.Errorln("Failed to get disk info: ", err)
		return nil, nil
	}

	if provider.debug {
		log.Debugf("Got disk space information capacity:%v usage:%v\n", diskUsage.Capacity, diskUsage.Usage)
		log.Debugf("Got disk inode information capacity:%v usage:%v\n", diskUsage.Inodes, diskUsage.InodesUsed)
	}

	diskUsageMetric := NodeMetric{
		Type:  "DiskUsage",
		Value: 100 - ((float64(diskUsage.Capacity-diskUsage.Usage) / float64(diskUsage.Capacity)) * 100),
	}
	fileUsageMetric := NodeMetric{
		Type:  "FileUsage",
		Value: 100 - ((float64(diskUsage.Inodes-diskUsage.InodesUsed) / float64(diskUsage.Inodes)) * 100),
	}

	return []NodeMetric{diskUsageMetric, fileUsageMetric}, nil
}
