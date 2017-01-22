package gautocloud_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGautocf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gautocf Suite")
}
