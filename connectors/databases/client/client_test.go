package client_test

import (
	. "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mysql"
	. "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mongodb"
	. "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mssql"
	. "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/postgresql"
	. "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/redis"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)

var _ = Describe("Client", func() {
	Context("Mysql", func() {
		Context("GetConnString", func() {
			mysqlConnector := MysqlConnector{}
			Context("When there is no password given", func() {
				PIt("should return correct connection string", func() {
					connString := mysqlConnector.GetConnString(dbtype.MysqlDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Database: "db",
						Options: "options=1",
					})
					Expect(connString).Should(Equal("user@tcp(localhost:3306)/db?options=1"))
				})
			})
			Context("When there is no options given", func() {
				It("should return correct connection string", func() {
					connString := mysqlConnector.GetConnString(dbtype.MysqlDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Password: "pass",
						Database: "db",
					})
					Expect(connString).Should(Equal("user:pass@tcp(localhost:3306)/db"))
				})
			})
			Context("When there is no password and no options given", func() {
				It("should return correct connection string", func() {
					connString := mysqlConnector.GetConnString(dbtype.MysqlDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Database: "db",
					})
					Expect(connString).Should(Equal("user@tcp(localhost:3306)/db"))
				})
			})
			Context("When there is password and options given", func() {
				PIt("should return correct connection string", func() {
					connString := mysqlConnector.GetConnString(dbtype.MysqlDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Password: "pass",
						Options: "options=1",
						Database: "db",
					})
					Expect(connString).Should(Equal("user:pass@tcp(localhost:3306)/db?options=1"))
				})
			})
		})
	})

	Context("Mongodb", func() {
		Context("GetConnString", func() {
			mongodbConnector := MongodbConnector{}
			Context("When there is no password given", func() {
				It("should return correct connection string", func() {
					connString := mongodbConnector.GetConnString(dbtype.MongodbDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Database: "db",
						Options: "options=1",
					})
					Expect(connString).Should(Equal("mongodb://user@localhost:3306/db?options=1"))
				})
			})
			Context("When there is no options given", func() {
				It("should return correct connection string", func() {
					connString := mongodbConnector.GetConnString(dbtype.MongodbDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Password: "pass",
						Database: "db",
					})
					Expect(connString).Should(Equal("mongodb://user:pass@localhost:3306/db"))
				})
			})
			Context("When there is no password and no options given", func() {
				It("should return correct connection string", func() {
					connString := mongodbConnector.GetConnString(dbtype.MongodbDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Database: "db",
					})
					Expect(connString).Should(Equal("mongodb://user@localhost:3306/db"))
				})
			})
			Context("When there is password and options given", func() {
				It("should return correct connection string", func() {
					connString := mongodbConnector.GetConnString(dbtype.MongodbDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Password: "pass",
						Options: "options=1",
						Database: "db",
					})
					Expect(connString).Should(Equal("mongodb://user:pass@localhost:3306/db?options=1"))
				})
			})
		})
	})
	Context("Mssql", func() {
		Context("GetConnString", func() {
			mssqlConnector := MssqlConnector{}
			Context("When there is no password given", func() {
				It("should return correct connection string", func() {
					connString := mssqlConnector.GetConnString(dbtype.MssqlDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Database: "db",
						Options: "options=1",
					})
					Expect(connString).Should(Equal("sqlserver://user@localhost:3306?database=db&options=1"))
				})
			})
			Context("When there is no options given", func() {
				It("should return correct connection string", func() {
					connString := mssqlConnector.GetConnString(dbtype.MssqlDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Password: "pass",
						Database: "db",
					})
					Expect(connString).Should(Equal("sqlserver://user:pass@localhost:3306?database=db"))
				})
			})
			Context("When there is no password and no options given", func() {
				It("should return correct connection string", func() {
					connString := mssqlConnector.GetConnString(dbtype.MssqlDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Database: "db",
					})
					Expect(connString).Should(Equal("sqlserver://user@localhost:3306?database=db"))
				})
			})
			Context("When there is password and options given", func() {
				It("should return correct connection string", func() {
					connString := mssqlConnector.GetConnString(dbtype.MssqlDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Password: "pass",
						Options: "options=1",
						Database: "db",
					})
					Expect(connString).Should(Equal("sqlserver://user:pass@localhost:3306?database=db&options=1"))
				})
			})
		})
	})
	Context("Postgresql", func() {
		Context("GetConnString", func() {
			postgresConnector := PostgresqlConnector{}
			Context("When there is no password given", func() {
				It("should return correct connection string", func() {
					connString := postgresConnector.GetConnString(dbtype.PostgresqlDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Database: "db",
						Options: "options=1",
					})
					Expect(connString).Should(Equal("postgres://user@localhost:3306/db?options=1"))
				})
			})
			Context("When there is no options given", func() {
				It("should return correct connection string", func() {
					connString := postgresConnector.GetConnString(dbtype.PostgresqlDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Password: "pass",
						Database: "db",
					})
					Expect(connString).Should(Equal("postgres://user:pass@localhost:3306/db"))
				})
			})
			Context("When there is no password and no options given", func() {
				It("should return correct connection string", func() {
					connString := postgresConnector.GetConnString(dbtype.PostgresqlDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Database: "db",
					})
					Expect(connString).Should(Equal("postgres://user@localhost:3306/db"))
				})
			})
			Context("When there is password and options given", func() {
				It("should return correct connection string", func() {
					connString := postgresConnector.GetConnString(dbtype.PostgresqlDatabase{
						Host: "localhost",
						Port: 3306,
						User: "user",
						Password: "pass",
						Options: "options=1",
						Database: "db",
					})
					Expect(connString).Should(Equal("postgres://user:pass@localhost:3306/db?options=1"))
				})
			})
		})
	})
	Context("Redis", func() {
		Context("GetConnString", func() {
			redisConnector := RedisConnector{}
			It("should return correct connection string", func() {
				connString := redisConnector.GetConnString(dbtype.RedisDatabase{
					Host: "localhost",
					Port: 3306,
				})
				Expect(connString).Should(Equal("localhost:3306"))
			})
		})
	})
})
