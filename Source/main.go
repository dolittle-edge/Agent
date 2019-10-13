/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package main

// Daemon considerations: https://fabianlee.org/2017/05/21/golang-running-a-go-binary-as-a-systemd-service-on-ubuntu-16-04/

import (
	"agent/configuring"
	"agent/configuring/configurers/location"
	"agent/log"
	"agent/provisioning"
	"agent/provisioning/system"
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
	info := flag.Bool("system-information", false, "Prints the system information of this node")
	help := flag.Bool("help", false, "Prints this message")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *info {
		information, err := system.ReadSystemInformation()
		if err != nil {
			fmt.Println("Could not read system information:", err)
		} else {
			fmt.Println()
			fmt.Println("System information")
			fmt.Println()
			information.Print(os.Stdout)
		}
		os.Exit(0)
	}

	provisoner := provisioning.NewProvider()
	provisoner.SetDebug(*debug)

	configurers := []configuring.ICanConfigureNode{
		location.NewConfigurer(),
	}

	configurator, err := configuring.NewConfigurator(provisoner, configurers)
	if err != nil {
		log.Errorln("Could not initialize configurator", err)
		os.Exit(1)
	}
	configurator.SetDebug(*debug)

	providers := []reporting.ICanProvideTelemetryForNode{
		disk.NewUsageTelemetryProvider(),
		memory.NewTelemetryProvider(),
		network.NewAddressProvider(),
		network.NewPingProvider(),
	}

	reporter := reporting.NewTelemetryReporter(provisoner, providers)
	reporter.SetDebug(*debug)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(2 * time.Second)

	log.Informationln("Starting the agent")
	provisoner.Start()

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
