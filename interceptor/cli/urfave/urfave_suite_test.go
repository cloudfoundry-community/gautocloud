package urfave_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUrfave(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Urfave Suite")
}
