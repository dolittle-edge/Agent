/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package reporting

//NodeMetric represents a single sampling for telemetry purposes
type NodeMetric struct {
	Type  string
	Value float64
}
