package client_test

import (
	. "github.com/cloudfoundry-community/gautocloud/connectors/amqp/client"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cloudfoundry-community/gautocloud/connectors/amqp/amqptype"
)

var _ = Describe("AmqpConnector", func() {
	Context("GetConnString", func() {
		mysqlConnector := AmqpConnector{}
		Context("When there is no password given", func() {
			It("should return correct connection string", func() {
				connString := mysqlConnector.GetConnString(amqptype.Amqp{
					Host: "localhost",
					Port: 3306,
					User: "user",
				})
				Expect(connString).Should(Equal("amqp://user@localhost:3306/"))
			})
		})
		Context("When there is password given", func() {
			It("should return correct connection string", func() {
				connString := mysqlConnector.GetConnString(amqptype.Amqp{
					Host: "localhost",
					Port: 3306,
					User: "user",
					Password: "pass",
				})
				Expect(connString).Should(Equal("amqp://user:pass@localhost:3306/"))
			})
		})
	})
})
