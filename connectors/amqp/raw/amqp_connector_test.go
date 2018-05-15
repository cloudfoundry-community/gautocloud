package raw_test

import (
	. "github.com/cloudfoundry-community/gautocloud/connectors/amqp/raw"

	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/amqp/amqptype"
	"github.com/cloudfoundry-community/gautocloud/connectors/amqp/schema"
	"github.com/cloudfoundry-community/gautocloud/decoder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AmqpConnector", func() {
	var connector connectors.Connector
	BeforeEach(func() {
		connector = NewAmqpRawConnector()
	})
	It("Should return a Amqp struct when passing a AmqpSchema without uri", func() {
		data, err := connector.Load(schema.AmqpSchema{
			Host:     "localhost",
			Password: "pass",
			User:     "user",
			Port:     5672,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(data).Should(BeEquivalentTo(
			amqptype.Amqp{
				Host:     "localhost",
				Password: "pass",
				User:     "user",
				Port:     5672,
			},
		))
	})
	It("Should return a Amqp struct when passing a AmqpSchema with an uri", func() {
		data, err := connector.Load(schema.AmqpSchema{
			Uri: decoder.ServiceUri{
				Host:     "localhost",
				Username: "user",
				Password: "pass",
				Port:     5672,
			},
			Vhost: "foo",
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(data).Should(BeEquivalentTo(
			amqptype.Amqp{
				Host:     "localhost",
				Password: "pass",
				User:     "user",
				Port:     5672,
				Vhost:    "foo",
			},
		))
	})

})
