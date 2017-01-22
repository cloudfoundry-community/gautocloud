package cloudenv_test

import (
	. "github.com/cloudfoundry-community/gautocloud/cloudenv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cloudfoundry-community/go-cfenv"
	"os"
)

var _ = Describe("CfCloudenv", func() {
	var app *cfenv.App
	var cloudEnv CloudEnv
	BeforeEach(func() {
		validEnv := []string{
			`VCAP_APPLICATION={"instance_id":"451f045fd16427bb99c895a2649b7b2a","instance_index":0,"host":"0.0.0.0","port":61857,"started_at":"2013-08-12 00:05:29 +0000","started_at_timestamp":1376265929,"start":"2013-08-12 00:05:29 +0000","state_timestamp":1376265929,"limits":{"mem":512,"disk":1024,"fds":16384},"application_version":"c1063c1c-40b9-434e-a797-db240b587d32","application_name":"styx-james","application_uris":["styx-james.a1-app.cf-app.com"],"version":"c1063c1c-40b9-434e-a797-db240b587d32","name":"styx-james","space_id":"3e0c28c5-6d9c-436b-b9ee-1f4326e54d05","space_name":"jdk","uris":["styx-james.a1-app.cf-app.com"],"users":null}`,
			`HOME=/home/vcap/app`,
			`MEMORY_LIMIT=512m`,
			`PWD=/home/vcap`,
			`TMPDIR=/home/vcap/tmp`,
			`USER=vcap`,
			`VCAP_SERVICES={"elephantsql-dev":[{"name":"myelephantsql-dev-c6c60","label":"elephantsql-dev","tags":["New Product","relational","service","Data Store","postgresql"],"plan":"turtle","credentials":{"uri":"postgres://seilbmbd:PHxTPJSbkcDakfK4cYwXHiIX9Q8p5Bxn@babar.elephantsql.com:5432/seilbmbd"}}],"sendgrid":[{"name":"mysendgrid","label":"sendgrid","tags":["smtp","Email", "service"],"plan":"free","credentials":{"hostname":"smtp.sendgrid.net","username":"QvsXMbJ3rK","password":"HCHMOYluTv"}}]}`,
		}
		app, _ = cfenv.New(cfenv.Env(validEnv))
		cloudEnv = NewCfCloudEnvWithAppEnv(app)
	})
	Context("GetServicesFromTags", func() {
		It("should give correct service(s)", func() {
			services := cloudEnv.GetServicesFromTags([]string{"service"})
			Expect(services).Should(HaveLen(2))
		})
		It("should give correct service(s) when tag have regex", func() {
			services := cloudEnv.GetServicesFromTags([]string{"postgres.*"})
			Expect(services).Should(HaveLen(1))
		})
		It("should give correct service(s) when have mulitple tag", func() {
			services := cloudEnv.GetServicesFromTags([]string{"postgresql", "smtp"})
			Expect(services).Should(HaveLen(2))
		})
	})
	Context("GetServicesFromName", func() {
		It("should give correct service(s)", func() {
			services := cloudEnv.GetServicesFromName(".*my.*")
			Expect(services).Should(HaveLen(2))

			services = cloudEnv.GetServicesFromName(".*sql.*")
			Expect(services).Should(HaveLen(1))
		})
	})
	Context("IsInCloudEnv", func() {
		It("should return false when have VCAP_APPLICATION env var empty", func() {
			err := os.Setenv("VCAP_APPLICATION", "")
			Expect(err).NotTo(HaveOccurred())

			Expect(cloudEnv.IsInCloudEnv()).Should(BeFalse())
		})
		It("should return true when have VCAP_APPLICATION env var not empty", func() {
			err := os.Setenv("VCAP_APPLICATION", "data")
			Expect(err).NotTo(HaveOccurred())

			Expect(cloudEnv.IsInCloudEnv()).Should(BeTrue())
		})
	})
	Context("GetAppInfo", func() {
		It("should return informations about instance of the running application", func() {
			appInfo := cloudEnv.GetAppInfo()
			Expect(appInfo.Id).Should(Equal("451f045fd16427bb99c895a2649b7b2a"))
			Expect(appInfo.Name).Should(Equal("styx-james"))
			Expect(appInfo.Properties).Should(BeEquivalentTo(map[string]interface{}{
				"uris": []string{"styx-james.a1-app.cf-app.com"},
				"host": "0.0.0.0",
				"home": "/home/vcap/app",
				"index": 0,
				"memory_limit": "512m",
				"port": 61857,
				"space_id": "3e0c28c5-6d9c-436b-b9ee-1f4326e54d05",
				"space_name": "jdk",
				"temp_dir": "/home/vcap/tmp",
				"user": "vcap",
				"version": "c1063c1c-40b9-434e-a797-db240b587d32",
				"working_dir": "/home/vcap",
			}))
		})
	})
})
