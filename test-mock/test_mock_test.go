//go:build gautocloud_mock
// +build gautocloud_mock

package test_mock_test

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/loader/fake"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type FakeService struct {
	Data string
}
type SecondFakeService struct {
	Data string
}

var _ = Describe("TestMock", func() {
	mockedLoader := gautocloud.Loader().(*fake.MockLoader)
	Context("When injecting a config", func() {
		var fakeService FakeService
		BeforeEach(func() {
			mockedLoader.EXPECT().Inject(gomock.Any()).Do(func(fakeService *FakeService) {
				fakeService.Data = "data"
			}).Return(nil)
		})
		It("should inject what user want", func() {
			err := gautocloud.Inject(&fakeService)
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeService.Data).To(Equal("data"))
		})
	})
	Context("When injecting a second config", func() {
		var fakeService SecondFakeService
		BeforeEach(func() {
			mockedLoader.EXPECT().Inject(gomock.Any()).Do(func(fakeService *SecondFakeService) {
				fakeService.Data = "data"
			}).Return(nil)
		})
		It("should inject what user want", func() {
			err := gautocloud.Inject(&fakeService)
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeService.Data).To(Equal("data"))
		})
	})
})
