package cloudenv_test

import (
	. "github.com/cloudfoundry-community/gautocloud/cloudenv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("HerokuCloudenv", func() {
	var cloudEnv CloudEnv
	BeforeEach(func() {
		services := []string{
			"DYNO=id1",
			"GAUTOCLOUD_APP_NAME=myapp",
			"PORT=6001",
			"CLEARDB_DATABASE_URL=mysql://host.com/database",
			"CLEARDB_DATABASE_USER=myuser",
			"S3=s3://host.com/mys3",
			"MY_MYSQL_DATABASE_HOST=host.com",
			"MY_MYSQL_DATABASE_USER=user",
			"MY_MYSQL_DATABASE_PASSWORD=password",
			`SVC_CONFIG={"foo": "bar"}`,
		}
		cloudEnv = NewHerokuCloudEnvEnvironment(services)
	})
	Context("GetServicesFromTags", func() {
		It("should give correct service(s)", func() {
			services := cloudEnv.GetServicesFromTags([]string{"CLEARDB_DATABASE"})
			Expect(services).Should(HaveLen(1))
			Expect(services[0].Credentials).Should(HaveLen(2))
			Expect(services[0].Credentials["url"]).Should(Equal("mysql://host.com/database"))
			Expect(services[0].Credentials["user"]).Should(Equal("myuser"))
		})
		It("should give correct service(s) when multiple tags are given", func() {
			services := cloudEnv.GetServicesFromTags([]string{"CLEARDB_DATABASE", "MY_MYSQL_DATABASE"})
			Expect(services).Should(HaveLen(2))
			Expect(services[0].Credentials).Should(HaveLen(2))
			Expect(services[0].Credentials["url"]).Should(Equal("mysql://host.com/database"))
			Expect(services[0].Credentials["user"]).Should(Equal("myuser"))
			Expect(services[1].Credentials).Should(HaveLen(3))
			Expect(services[1].Credentials["host"]).Should(Equal("host.com"))
			Expect(services[1].Credentials["user"]).Should(Equal("user"))
			Expect(services[1].Credentials["password"]).Should(Equal("password"))
		})
		It("should give correct service(s) when giving first level tag", func() {
			services := cloudEnv.GetServicesFromTags([]string{"MY"})
			Expect(services).Should(HaveLen(1))
			Expect(services[0].Credentials).Should(HaveLen(3))
			Expect(services[0].Credentials["mysql_database_host"]).Should(Equal("host.com"))
			Expect(services[0].Credentials["mysql_database_user"]).Should(Equal("user"))
			Expect(services[0].Credentials["mysql_database_password"]).Should(Equal("password"))
		})
		It("should give correct service with uri when giving a tag which target an env var", func() {
			services := cloudEnv.GetServicesFromTags([]string{"S3"})
			Expect(services).Should(HaveLen(1))
			Expect(services[0].Credentials).Should(HaveLen(2))
			Expect(services[0].Credentials["s3"]).Should(Equal("s3://host.com/mys3"))
			Expect(services[0].Credentials["uri"]).Should(Equal("s3://host.com/mys3"))
		})
		Context("Value is json encoded", func() {
			It("should give correct service with json decoded when giving a tag which target an env var", func() {
				services := cloudEnv.GetServicesFromTags([]string{"SVC_CONFIG"})
				Expect(services).Should(HaveLen(1))
				Expect(services[0].Credentials).Should(HaveLen(1))
				Expect(services[0].Credentials["foo"]).Should(Equal("bar"))
			})
		})

	})
	Context("GetServicesFromName", func() {
		It("should give correct service(s) for non alone identifier", func() {
			servicesDb := cloudEnv.GetServicesFromName(".*database.*")
			Expect(servicesDb).Should(HaveLen(2))
			service := servicesDb[0]
			if len(service.Credentials) == 3 {
				service = servicesDb[1]
			}
			Expect(service.Credentials).Should(HaveLen(2))
			Expect(service.Credentials["url"]).Should(Equal("mysql://host.com/database"))
			Expect(service.Credentials["user"]).Should(Equal("myuser"))
			service = servicesDb[1]
			if len(service.Credentials) == 2 {
				service = servicesDb[0]
			}
			Expect(service.Credentials).Should(HaveLen(3))
			Expect(service.Credentials["host"]).Should(Equal("host.com"))
			Expect(service.Credentials["user"]).Should(Equal("user"))
			Expect(service.Credentials["password"]).Should(Equal("password"))
		})
		It("should give correct service(s) for alone identifier", func() {
			servicesS3 := cloudEnv.GetServicesFromName(".*s3.*")
			Expect(servicesS3).Should(HaveLen(1))
			Expect(servicesS3[0].Credentials).Should(HaveLen(2))
			Expect(servicesS3[0].Credentials["s3"]).Should(Equal("s3://host.com/mys3"))
			Expect(servicesS3[0].Credentials["uri"]).Should(Equal("s3://host.com/mys3"))
		})
		Context("Value is json encoded", func() {
			It("should give correct service with json decoded", func() {
				services := cloudEnv.GetServicesFromName(".*config")
				Expect(services).Should(HaveLen(1))
				Expect(services[0].Credentials).Should(HaveLen(1))
				Expect(services[0].Credentials["foo"]).Should(Equal("bar"))
			})
		})
	})
	Context("GetAppInfo", func() {
		It("should return informations about instance of the running application", func() {
			appInfo := cloudEnv.GetAppInfo()
			Expect(appInfo.Id).Should(Equal("id1"))
			Expect(appInfo.Name).Should(Equal("myapp"))
			Expect(appInfo.Port).Should(Equal(6001))
			Expect(appInfo.Properties["port"]).Should(Equal(6001))
		})
	})
	Context("IsInCloudEnv", func() {
		It("should return false when have DYNO env var exists", func() {
			os.Unsetenv("DYNO")
			Expect(cloudEnv.IsInCloudEnv()).Should(BeFalse())
		})
		It("should return true when have DYNO env var not exists", func() {
			err := os.Setenv("DYNO", "")
			Expect(err).NotTo(HaveOccurred())

			Expect(cloudEnv.IsInCloudEnv()).Should(BeTrue())
		})
	})
})
