/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package main

// Daemon considerations: https://fabianlee.org/2017/05/21/golang-running-a-go-binary-as-a-systemd-service-on-ubuntu-16-04/

import (
	"agent/log"
	"agent/provisioning"
	"agent/reporting"
	"agent/reporting/providers/disk"
	"agent/reporting/providers/memory"
	"agent/reporting/providers/network"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Dolittle Edge Agent - (C) Dolittle")

	debug := flag.Bool("debug", false, "Enable debug output")
	help := flag.Bool("help", false, "Prints this message")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	provisoner := provisioning.NewProvider()
	provisoner.SetDebug(*debug)

	providers := []reporting.ICanProvideTelemetryForNode{
		disk.NewUsageTelemetryProvider(),
		memory.NewTelemetryProvider(),
		network.NewAddressProvider(),
		network.NewPingProvider(),
	}

	for _, provider := range providers {
		provider.SetDebug(*debug)
	}

	reporter := reporting.NewTelemetryReporter(provisoner, providers)
	reporter.SetDebug(*debug)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(2 * time.Second)

	log.Informationln("Starting the agent")

	for {
		select {
		case <-ticker.C:
			reporter.ReportCurrentStatus()
		case <-quit:
			log.Informationln("Stopping the agent")
			os.Exit(0)
		}
	}
}
