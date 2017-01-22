package decoder_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDecoder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Decoder Suite")
}
