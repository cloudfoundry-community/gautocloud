//go:build gautocloud_mock
// +build gautocloud_mock

package test_mock_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTestMock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestMock Suite")
}
