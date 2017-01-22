package cloudenv_test

import (
	. "github.com/cloudfoundry-community/gautocloud/cloudenv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"bytes"
	"path"
)

var yamlServices = []byte(`
app_name: "myapp"
services:
- name: myelephantsql
  tags: [postgresql, service]
  credentials:
    uri: postgres://seilbmbd:PHxTPJSbkcDakfK4cYwXHiIX9Q8p5Bxn@babar.elephantsql.com:5432/seilbmbd
- name: mysendgrid
  tags: [smtp, service]
  credentials:
    hostname: smtp.sendgrid.net
    username: QvsXMbJ3rK
    password: HCHMOYluTv
`)
var jsonServices = []byte(`
{
"app_name": "myapp",
"services": [
  {
  "name": "myelephantsql",
  "tags": ["postgresql", "service"],
  "credentials":
    {
      "uri": "postgres://seilbmbd:PHxTPJSbkcDakfK4cYwXHiIX9Q8p5Bxn@babar.elephantsql.com:5432/seilbmbd"
    }
  },
  {
  "name": "mysendgrid",
  "tags": ["smtp", "service"],
  "credentials":
    {
      "hostname": "smtp.sendgrid.net",
      "username": "QvsXMbJ3rK",
      "password": "HCHMOYluTv"
    }
  }
 ]
}
`)
var tomlServices = []byte(`

app_name = "myapp"

[[services]]
name = "myelephantsql"
tags = [ "postgresql", "service" ]
	[services.credentials]
	uri = "postgres://seilbmbd:PHxTPJSbkcDakfK4cYwXHiIX9Q8p5Bxn@babar.elephantsql.com:5432/seilbmbd"
[[services]]
name = "mysendgrid"
tags = [ "smtp", "service" ]
	[services.credentials]
	hostname = "smtp.sendgrid.net"
	username = "QvsXMbJ3rK"
	password = "HCHMOYluTv"
`)
var hclServices = []byte(`
services {
  name = "myelephantsql"
  tags = ["postgresql", "service"]
  credentials {
    uri = "postgres://seilbmbd:PHxTPJSbkcDakfK4cYwXHiIX9Q8p5Bxn@babar.elephantsql.com:5432/seilbmbd"
  }
}
services {
  name = "mysendgrid"
  tags = ["smtp", "service"]
  credentials {
    hostname = "smtp.sendgrid.net"
    username = "QvsXMbJ3rK"
    password = "HCHMOYluTv"
  }
}
`)

type FormatService struct {
	Type    string
	Content []byte
}

func defaultTest(cloudEnv CloudEnv) {

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
	Context("GetAppInfo", func() {
		It("should return informations about instance of the running application", func() {
			appInfo := cloudEnv.GetAppInfo()
			Expect(appInfo.Id).ShouldNot(BeEmpty())
			Expect(appInfo.Name).Should(Or(Equal("myapp"), Equal("<unknown>")))
			Expect(appInfo.Properties).Should(HaveLen(0))
		})
	})
}

var _ = Describe("LocalCloudenv", func() {
	AfterEach(func() {
		os.Unsetenv(LOCAL_ENV_KEY)
	})
	formatServices := []FormatService{
		{
			Type: "yaml",
			Content: yamlServices,
		},
		{
			Type: "json",
			Content: jsonServices,
		},
		{
			Type: "toml",
			Content: tomlServices,
		},
		{
			Type: "hcl",
			Content: hclServices,
		},
	}
	for _, formatService := range formatServices {
		Describe("When config file is a " + formatService.Type + " file", func() {
			var cloudEnv CloudEnv
			cloudEnv = NewLocalCloudEnvFromReader(bytes.NewBuffer(formatService.Content), formatService.Type)
			defaultTest(cloudEnv)
		})
	}
	Describe("When config file path is provided", func() {
		dir, err := os.Getwd()
		if err != nil {
			Fail(err.Error())
		}
		fixturePath := path.Join(dir, "test-fixtures")
		configPath := path.Join(fixturePath, "services.json")
		err = os.Setenv(LOCAL_ENV_KEY, configPath)
		if err != nil {
			Fail(err.Error())
		}
		cloudEnv := NewLocalCloudEnv()
		cloudEnv.Load()
		defaultTest(cloudEnv)

	})

	Context("IsInCloudEnv", func() {
		var cloudEnv CloudEnv
		BeforeEach(func() {
			cloudEnv = NewLocalCloudEnvFromReader(bytes.NewBuffer(formatServices[0].Content), formatServices[0].Type)
		})
		It("should return false when have " + LOCAL_ENV_KEY + " env var empty", func() {
			err := os.Setenv(LOCAL_ENV_KEY, "")
			Expect(err).NotTo(HaveOccurred())

			Expect(cloudEnv.IsInCloudEnv()).Should(BeFalse())
		})
		It("should return true when have " + LOCAL_ENV_KEY + " env var not empty", func() {
			err := os.Setenv(LOCAL_ENV_KEY, "data")
			Expect(err).NotTo(HaveOccurred())

			Expect(cloudEnv.IsInCloudEnv()).Should(BeTrue())
		})
	})

})
