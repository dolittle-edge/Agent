/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package reporting

// ICanProvideTelemetryForNode defined a system that is capable of providing a TelemetrySample when asked
type ICanProvideTelemetryForNode interface {
	Provide() ([]NodeMetric, []NodeInfo)
	SetDebug(bool)
}
