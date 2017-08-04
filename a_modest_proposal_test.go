package main_test

import (
	"a-modest-proposal/printer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AModestProposal", func() {
	It("should print out configuration stuff", func() {
		Expect(printer.PrintConfig()).To(Equal("Configuration Complete. Battlecruiser Operational."))
	})
})
