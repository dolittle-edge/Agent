/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package main

// Daemon considerations: https://fabianlee.org/2017/05/21/golang-running-a-go-binary-as-a-systemd-service-on-ubuntu-16-04/

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	. "agent/reporting"
)

func mainloop() {
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}

func format(val uint64) uint64 {
	return val / 1024
}

func protect(g func()) {
	defer func() {
		log.Println("done") // Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}

func main() {
	fmt.Println("Dolittle Edge Agent - (C) Dolittle")


	memoryProvider := new(MemoryTelemetryProvider)
	diskUsageProvider := new(DiskUsageTelemetryProvider)
	currentNode := ReadConfiguration()

	providers := []ICanProvideTelemetryForNode{memoryProvider, diskUsageProvider}
	reporter := TelemetryReporter{}.New(currentNode, providers)

	fmt.Println("Starting")
	reporter.ReportCurrentStatus()

	ticker := time.NewTicker(30 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				protect(reporter.ReportCurrentStatus)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	mainloop()
}
