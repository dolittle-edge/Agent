/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package reporting

import "testing"
import "github.com/dolittle-edge/agent/reporting"

type FakeProvider struct {
}

func (provider FakeProvider) Provide() []*TelemetrySample {
	samples := []*TelemetrySample{}

	return samples
}

func when_reporting_current_status_with_two_providers(T *testing.T) {
	fp := new(FakeProvider)
	fp.Provide()
}
