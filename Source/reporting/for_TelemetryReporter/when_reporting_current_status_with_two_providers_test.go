/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package reporting

import (
	. "agent/reporting"
	. "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)


type FakeProvider struct {
}

func (provider FakeProvider) Provide() []*TelemetrySample {
	samples := []*TelemetrySample{}
	
	return samples
}


/*
func Test_when_reporting_current_status_with_two_providers(T *testing.T) {
}
*/

var _ = Describe("when reporting current status with two providers", func() {
	Context("", func() {

	})
	
	BeforeEach(func() {
		fp := new(FakeProvider)
		fp.Provide()
	
	})

	It("should do stuff", func() {
		Expect(42).To(Equal(42));
	})
})