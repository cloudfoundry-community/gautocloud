package client_test

import (
	. "github.com/cloudfoundry-community/gautocloud/connectors/amqp/client"

	"github.com/cloudfoundry-community/gautocloud/connectors/amqp/amqptype"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AmqpConnector", func() {
	Context("GetConnString", func() {
		amqpConnector := AmqpConnector{}
		Context("When there is no password given", func() {
			It("should return correct connection string", func() {
				connString := amqpConnector.GetConnString(amqptype.Amqp{
					Host: "localhost",
					Port: 5672,
					User: "user",
				})
				Expect(connString).Should(Equal("amqp://user@localhost:5672/"))
			})
		})
		Context("When there is password given", func() {
			It("should return correct connection string", func() {
				connString := amqpConnector.GetConnString(amqptype.Amqp{
					Host:     "localhost",
					Port:     5672,
					User:     "user",
					Password: "pass",
				})
				Expect(connString).Should(Equal("amqp://user:pass@localhost:5672/"))
			})
		})
		Context("When there is vhost given", func() {
			It("should return correct connection string", func() {
				connString := amqpConnector.GetConnString(amqptype.Amqp{
					Host:     "localhost",
					Port:     5672,
					User:     "user",
					Password: "pass",
					Vhost:    "foo",
				})
				Expect(connString).Should(Equal("amqp://user:pass@localhost:5672/foo"))
			})
		})

	})
})
