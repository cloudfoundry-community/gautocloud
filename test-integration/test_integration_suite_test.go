package test_integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"os"
	"fmt"
)

func TestTestIntegration(t *testing.T) {

	RegisterFailHandler(Fail)
	if os.Getenv("GAUTOCLOUD_HOST_SERVICES") == "" {
		fmt.Println("Integration tests skipped, you need to set env var " +
			"GAUTOCLOUD_HOST_SERVICES which target a docker host " +
			"(can be localhost or ip of a docker-machine)")
		t.SkipNow()
		return
	}

	RunSpecs(t, "TestIntegration Suite")
}
