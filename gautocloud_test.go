package gautocloud_test

import (
	. "github.com/cloudfoundry-community/gautocloud"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cloudfoundry-community/gautocloud/cloudenv"
)

var _ = Describe("Gautocloud", func() {
	It("should contains default cloud env in right order", func() {
		Expect(CloudEnvs()).Should(HaveLen(4))
		Expect(CloudEnvs()[0].Name()).Should(Equal(cloudenv.CfCloudEnv{}.Name()))
		Expect(CloudEnvs()[1].Name()).Should(Equal(cloudenv.HerokuCloudEnv{}.Name()))
		Expect(CloudEnvs()[2].Name()).Should(Equal(cloudenv.KubernetesCloudEnv{}.Name()))
		Expect(CloudEnvs()[3].Name()).Should(Equal(cloudenv.LocalCloudEnv{}.Name()))
	})
})
