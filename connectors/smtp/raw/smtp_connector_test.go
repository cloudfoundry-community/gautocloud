package raw_test

import (
	. "github.com/cloudfoundry-community/gautocloud/connectors/smtp/raw"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cloudfoundry-community/gautocloud/decoder"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/smtp/schema"
	"github.com/cloudfoundry-community/gautocloud/connectors/smtp/smtptype"
)

var _ = Describe("SmtpConnector", func() {
	var connector connectors.Connector
	BeforeEach(func() {
		connector = NewSmtpRawConnector()
	})
	It("Should return a Smtp struct when passing a SmtpSchema without uri", func() {
		data, err := connector.Load(schema.SmtpSchema{
			Host: "localhost",
			Password: "pass",
			User: "user",
			Port: 3306,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(data).Should(BeEquivalentTo(
			smtptype.Smtp{
				Host: "localhost",
				Password: "pass",
				User: "user",
				Port: 3306,
			},
		))
	})
	It("Should return a Smtp struct when passing a SmtpSchema with an uri", func() {
		data, err := connector.Load(schema.SmtpSchema{
			Uri: decoder.ServiceUri{
				Host: "localhost",
				Username: "user",
				Password: "pass",
				Port: 3306,
			},
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(data).Should(BeEquivalentTo(
			smtptype.Smtp{
				Host: "localhost",
				Password: "pass",
				User: "user",
				Port: 3306,
			},
		))
	})
})
