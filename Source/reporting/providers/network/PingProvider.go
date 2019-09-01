/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package network

import (
	"agent/log"
	. "agent/reporting"
	"time"

	"github.com/sparrc/go-ping"
)

const (
	pingAddress = "8.8.8.8"
	pingPause   = 1 * time.Minute
)

// PingProvider provides metrics about network latency
type PingProvider struct {
	statistics ping.Statistics
	isRunning  bool
}

// NewPingProvider creates a new instance of PingProvider
func NewPingProvider() *PingProvider {
	provider := new(PingProvider)
	go provider.runPinger()
	return provider
}

func (provider *PingProvider) runPinger() {
	pinger, err := ping.NewPinger(pingAddress)
	if err != nil {
		log.Errorln("Could not initialize pinger:", err)
		return
	}
	pinger.Count = 5

	for {
		pinger.Run()
		provider.statistics = *pinger.Statistics()
		provider.isRunning = true
		time.Sleep(pingPause)
	}
}

// Provide the memory telemetry
func (provider *PingProvider) Provide() (metrics []NodeMetric, _ []NodeInfo) {
	if provider.isRunning {
		metrics = append(metrics,
			NodeMetric{
				Type:  "Network.Ping.Average",
				Value: float32(provider.statistics.AvgRtt / time.Millisecond),
			},
			NodeMetric{
				Type:  "Network.Ping.StdDev",
				Value: float32(provider.statistics.StdDevRtt / time.Millisecond),
			},
		)
	}
	return
}
