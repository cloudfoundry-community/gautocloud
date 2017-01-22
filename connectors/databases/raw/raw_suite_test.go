package raw_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRaw(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Raw Suite")
}
