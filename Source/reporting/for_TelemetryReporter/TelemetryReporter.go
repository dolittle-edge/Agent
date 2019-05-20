/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TelemetryReporter Represents a system that can report telemetry back to the cloud
type TelemetryReporter struct {
	CurrentNode Node
	Providers   []ICanProvideTelemetryForNode
}

// New Creates a new instance of the TelemetryReporter
func (TelemetryReporter) New(currentNode Node, providers []ICanProvideTelemetryForNode) *TelemetryReporter {
	reporter := new(TelemetryReporter)
	reporter.CurrentNode = currentNode
	reporter.Providers = providers

	return reporter
}

// ReportCurrentStatus Report the current status of the current node in the current location
func (reporter *TelemetryReporter) ReportCurrentStatus() {
	fmt.Println("Reporting")
	node := reporter.CurrentNode

	for _, provider := range reporter.Providers {
		samples := provider.Provide()

		for _, sample := range samples {
			node.State[sample.Type] = sample.Value
		}
	}

	json, _ := json.Marshal(node)
	fmt.Println(string(json))

	body := bytes.NewBuffer(json)
	response, _ := http.Post("https://edge.dolittle.studio/api/Telemetry", "application/json", body)
	//response, _ := http.Post("http://localhost:5000/api/Telemetry", "application/json", body)
	result, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(result))
}
