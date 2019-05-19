/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package main

// https://stackoverflow.com/questions/10067295/how-to-start-a-go-program-as-a-daemon-in-ubuntu
// http://www.ryanday.net/2012/09/04/the-problem-with-a-golang-daemon/
// https://www.captaincodeman.com/2015/03/05/dependency-injection-in-go-golang

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
