package loader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLoader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Loader Suite")
}
