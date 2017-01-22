package connectors_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestConnectors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Connectors Suite")
}
