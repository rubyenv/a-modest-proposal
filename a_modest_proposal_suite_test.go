package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAModestProposal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AModestProposal Suite")
}
