package raw_test

import (
	. "github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"

	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/schema"
	"github.com/cloudfoundry-community/gautocloud/decoder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Raw", func() {
	var connector connectors.Connector
	Context("Mongodb", func() {
		BeforeEach(func() {
			connector = NewMongodbRawConnector()
		})
		It("Should return a MongodbDatabase struct when passing a MongoDbSchema without uri", func() {
			data, err := connector.Load(schema.MongoDbSchema{
				Host:     "localhost",
				Password: "pass",
				User:     "user",
				Port:     3306,
				Database: "db",
				Options:  "tls=true",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				dbtype.MongodbDatabase{
					Database: "db",
					Host:     "localhost",
					Password: "pass",
					User:     "user",
					Port:     3306,
					Options:  "tls=true",
				},
			))
		})
		It("Should return a MongodbDatabase struct when passing a MongoDbSchema with an uri", func() {
			data, err := connector.Load(schema.MongoDbSchema{
				Uri: decoder.ServiceUri{
					Host:     "localhost",
					Name:     "db",
					Username: "user",
					Password: "pass",
					Port:     3306,
					RawQuery: "options=1",
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				dbtype.MongodbDatabase{
					Database: "db",
					Host:     "localhost",
					Password: "pass",
					User:     "user",
					Port:     3306,
					Options:  "options=1",
				},
			))
		})
	})
	Context("Mssql", func() {
		BeforeEach(func() {
			connector = NewMssqlRawConnector()
		})
		It("Should return a MssqlDatabase struct when passing a MssqlSchema without uri", func() {
			data, err := connector.Load(schema.MssqlSchema{
				Host:     "localhost",
				Password: "pass",
				User:     "user",
				Port:     3306,
				Database: "db",
				Options:  "tls=true",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				dbtype.MssqlDatabase{
					Database: "db",
					Host:     "localhost",
					Password: "pass",
					User:     "user",
					Port:     3306,
					Options:  "tls=true",
				},
			))
		})
		It("Should return a MssqlDatabase struct when passing a MssqlSchema with an uri", func() {
			data, err := connector.Load(schema.MssqlSchema{
				Uri: decoder.ServiceUri{
					Host:     "localhost",
					Name:     "db",
					Username: "user",
					Password: "pass",
					Port:     3306,
					RawQuery: "options=1",
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				dbtype.MssqlDatabase{
					Database: "db",
					Host:     "localhost",
					Password: "pass",
					User:     "user",
					Port:     3306,
					Options:  "options=1",
				},
			))
		})
	})
	Context("Oracle", func() {
		BeforeEach(func() {
			connector = NewOracleRawConnector()
		})
		It("Should return a OracleDatabase struct when passing a OracleSchema without uri", func() {
			data, err := connector.Load(schema.OracleSchema{
				Host:     "localhost",
				Password: "pass",
				User:     "user",
				Port:     3306,
				Database: "db",
				Options:  "tls=true",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				dbtype.OracleDatabase{
					Database: "db",
					Host:     "localhost",
					Password: "pass",
					User:     "user",
					Port:     3306,
					Options:  "tls=true",
				},
			))
		})
		It("Should return a OracleDatabase struct when passing a OracleSchema with an uri", func() {
			data, err := connector.Load(schema.OracleSchema{
				Uri: decoder.ServiceUri{
					Host:     "localhost",
					Name:     "db",
					Username: "user",
					Password: "pass",
					Port:     3306,
					RawQuery: "options=1",
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				dbtype.OracleDatabase{
					Database: "db",
					Host:     "localhost",
					Password: "pass",
					User:     "user",
					Port:     3306,
					Options:  "options=1",
				},
			))
		})
	})
	Context("Mysql", func() {
		BeforeEach(func() {
			connector = NewMysqlRawConnector()
		})
		It("Should return a MysqlDatabase struct when passing a MysqlSchema without uri", func() {
			data, err := connector.Load(schema.MysqlSchema{
				Host:     "localhost",
				Password: "pass",
				User:     "user",
				Port:     3306,
				Database: "db",
				Options:  "tls=true",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				dbtype.MysqlDatabase{
					Database: "db",
					Host:     "localhost",
					Password: "pass",
					User:     "user",
					Port:     3306,
					Options:  "tls=true",
				},
			))
		})
		It("Should return a MysqlDatabase struct when passing a MysqlSchema with an uri", func() {
			data, err := connector.Load(schema.MysqlSchema{
				Uri: decoder.ServiceUri{
					Host:     "localhost",
					Name:     "db",
					Username: "user",
					Password: "pass",
					Port:     3306,
					RawQuery: "options=1",
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				dbtype.MysqlDatabase{
					Database: "db",
					Host:     "localhost",
					Password: "pass",
					User:     "user",
					Port:     3306,
					Options:  "options=1",
				},
			))
		})
	})
	Context("Postgresql", func() {
		BeforeEach(func() {
			connector = NewPostgresqlRawConnector()
		})
		It("Should return a PostgresqlDatabase struct when passing a PostgresqlSchema without uri", func() {
			data, err := connector.Load(schema.PostgresqlSchema{
				Host:     "localhost",
				Password: "pass",
				User:     "user",
				Port:     3306,
				Database: "db",
				Options:  "tls=true",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				dbtype.PostgresqlDatabase{
					Database: "db",
					Host:     "localhost",
					Password: "pass",
					User:     "user",
					Port:     3306,
					Options:  "tls=true",
				},
			))
		})
		It("Should return a PostgresqlDatabase struct when passing a PostgresqlSchema with an uri", func() {
			data, err := connector.Load(schema.PostgresqlSchema{
				Uri: decoder.ServiceUri{
					Host:     "localhost",
					Name:     "db",
					Username: "user",
					Password: "pass",
					Port:     3306,
					RawQuery: "options=1",
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				dbtype.PostgresqlDatabase{
					Database: "db",
					Host:     "localhost",
					Password: "pass",
					User:     "user",
					Port:     3306,
					Options:  "options=1",
				},
			))
		})
	})
	Context("Redis", func() {
		BeforeEach(func() {
			connector = NewRedisRawConnector()
		})
		It("Should return a RedisDatabase struct when passing a RedisSchema without uri", func() {
			data, err := connector.Load(schema.RedisSchema{
				Host:     "localhost",
				Password: "pass",
				Port:     3306,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				dbtype.RedisDatabase{
					Host:     "localhost",
					Password: "pass",
					Port:     3306,
				},
			))
		})
		It("Should return a RedisDatabase struct when passing a RedisSchema with an uri", func() {
			data, err := connector.Load(schema.RedisSchema{
				Uri: decoder.ServiceUri{
					Host:     "localhost",
					Name:     "db",
					Username: "pass",
					Port:     3306,
					RawQuery: "options=1",
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				dbtype.RedisDatabase{
					Host:     "localhost",
					Password: "pass",
					Port:     3306,
				},
			))
		})
	})
})
