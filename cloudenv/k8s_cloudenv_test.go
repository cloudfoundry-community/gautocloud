package cloudenv_test

import (
	. "github.com/cloudfoundry-community/gautocloud/cloudenv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("KubernetesCloudenv", func() {
	var cloudEnv CloudEnv
	BeforeEach(func() {
		services := []string{
			"KUBERNETES_SERVICE_PORT=443",
			"KUBERNETES_PORT=tcp://10.100.200.1:443",
			"HEAPSTER_PORT_80_TCP_PROTO=tcp",
			"KUBERNETES_DASHBOARD_SERVICE_PORT=80",
			"KUBE_DNS_SERVICE_PORT_DNS_TCP=53",
			"HOSTNAME=kube-dns-3917200835-0m5h5",
			"MONITORING_INFLUXDB_SERVICE_PORT=80",
			"MONITORING_INFLUXDB_SERVICE_PORT_API=8086",
			"KUBE_DNS_SERVICE_HOST=10.100.200.10",
			"KUBE_DNS_SERVICE_PORT=53",
			"HEAPSTER_SERVICE_HOST=10.100.200.71",
			"KUBE_DNS_PORT_53_TCP_PORT=53",
			"KUBE_DNS_SERVICE_PORT_DNS=53",
			"KUBE_DNS_PORT_53_UDP_PROTO=udp",
			"MONITORING_INFLUXDB_SERVICE_PORT_HTTP=80",
			"KUBERNETES_SERVICE_PORT_HTTPS=443",
			"KUBERNETES_SERVICE_HOST=10.100.200.1",
			"MONITORING_INFLUXDB_SERVICE_HOST=10.100.200.49",
		}
		cloudEnv = NewKubernetesCloudEnvEnvironment(KubernetesCloudEnv{}.SanitizeEnvVars(services))

	})
	Context("GetServicesFromTags", func() {
		It("should give correct service(s)", func() {
			services := cloudEnv.GetServicesFromTags([]string{"KUBE_DNS"})
			Expect(services).Should(HaveLen(1))
			Expect(services[0].Credentials).Should(HaveLen(6))
			Expect(services[0].Credentials["host"]).Should(Equal("10.100.200.10"))
			Expect(services[0].Credentials["port"]).Should(Equal("53"))
		})
		It("should give correct service(s) when multiple tags are given", func() {
			services := cloudEnv.GetServicesFromTags([]string{"KUBE_DNS", "MONITORING_INFLUXDB"})
			Expect(services).Should(HaveLen(2))
			Expect(services[0].Credentials).Should(HaveLen(6))
			Expect(services[0].Credentials["host"]).Should(Equal("10.100.200.10"))
			Expect(services[0].Credentials["port"]).Should(Equal("53"))

			Expect(services[1].Credentials).Should(HaveLen(4))
			Expect(services[1].Credentials["port"]).Should(Equal("80"))
			Expect(services[1].Credentials["port_http"]).Should(Equal("80"))
			Expect(services[1].Credentials["port_api"]).Should(Equal("8086"))
			Expect(services[1].Credentials["host"]).Should(Equal("10.100.200.49"))
		})
		It("should give correct service(s) when giving first level tag", func() {
			services := cloudEnv.GetServicesFromTags([]string{"KUBE"})
			Expect(services).Should(HaveLen(1))
			Expect(services[0].Credentials).Should(HaveLen(6))
			Expect(services[0].Credentials["dns_host"]).Should(Equal("10.100.200.10"))
			Expect(services[0].Credentials["dns_port"]).Should(Equal("53"))
		})
	})
	Context("GetServicesFromName", func() {
		It("should give correct service(s) for non alone identifier", func() {
			services := cloudEnv.GetServicesFromName(".*kube.*")
			Expect(services).Should(HaveLen(2))
			service := services[0]
			if len(service.Credentials) == 6 {
				service = services[1]
			}
			Expect(service.Credentials).Should(HaveLen(5))
			Expect(service.Credentials["host"]).Should(Equal("10.100.200.1"))
			Expect(service.Credentials["port_https"]).Should(Equal("443"))

			service = services[1]
			if len(service.Credentials) == 5 {
				service = services[0]
			}
			Expect(service.Credentials).Should(HaveLen(6))
			Expect(service.Credentials["dns_host"]).Should(Equal("10.100.200.10"))
			Expect(service.Credentials["dns_port"]).Should(Equal("53"))
		})
		It("should give correct service(s) for alone identifier", func() {
			services := cloudEnv.GetServicesFromName(".*kubernetes.*")
			Expect(services).Should(HaveLen(1))
			Expect(services[0].Credentials).Should(HaveLen(5))
			Expect(services[0].Credentials["host"]).Should(Equal("10.100.200.1"))
			Expect(services[0].Credentials["port_https"]).Should(Equal("443"))
		})
	})
	Context("GetAppInfo", func() {
		It("should return informations about instance of the running application", func() {
			appInfo := cloudEnv.GetAppInfo()
			Expect(appInfo.Id).Should(Equal("kube-dns-3917200835-0m5h5"))
			Expect(appInfo.Name).Should(Equal("kube-dns-3917200835-0m5h5"))
			Expect(appInfo.Properties["host"]).Should(Equal("10.100.200.1"))
			Expect(appInfo.Properties["port"]).Should(Equal(443))
		})
	})
	Context("IsInCloudEnv", func() {
		It("should return false when have KUBERNETES_PORT env var exists", func() {
			os.Unsetenv("KUBERNETES_PORT")
			Expect(cloudEnv.IsInCloudEnv()).Should(BeFalse())
		})
		It("should return true when have KUBERNETES_PORT env var not exists", func() {
			err := os.Setenv("KUBERNETES_PORT", "")
			Expect(err).NotTo(HaveOccurred())

			Expect(cloudEnv.IsInCloudEnv()).Should(BeTrue())
		})
	})
})
