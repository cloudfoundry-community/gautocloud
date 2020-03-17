package test_integration_test

import (
	"net/smtp"
	"os"

	"github.com/cloudfoundry-community/gautocloud"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/all"
	camqp "github.com/cloudfoundry-community/gautocloud/connectors/amqp/client"
	coauth2 "github.com/cloudfoundry-community/gautocloud/connectors/auth/config/oauth2"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mongodb"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mssql"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mysql"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/client/postgresql"
	credis "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/redis"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	gmssql "github.com/cloudfoundry-community/gautocloud/connectors/databases/gorm/mssql"
	gmysql "github.com/cloudfoundry-community/gautocloud/connectors/databases/gorm/mysql"
	gpostgres "github.com/cloudfoundry-community/gautocloud/connectors/databases/gorm/postgresql"
	cs3goamz "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/client/s3/goamz"
	cs3minio "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/client/s3/minio"
	"github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype/miniotype"
	csmtp "github.com/cloudfoundry-community/gautocloud/connectors/smtp/client"
	. "github.com/cloudfoundry-community/gautocloud/test-utils"
	"github.com/go-redis/redis/v7"
	"github.com/goamz/goamz/s3"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"golang.org/x/oauth2"
	"gopkg.in/mgo.v2"
)

var _ = Describe("Connectors integration", func() {
	if os.Getenv("GAUTOCLOUD_HOST_SERVICES") == "" {
		return
	}

	log.SetLevel(log.DebugLevel)
	os.Unsetenv("MAIL") // travis set this env var which make connector detect it
	os.Setenv("MYSQL_URL", CreateEnvValue(ServiceUrl{
		Type:     "mysql",
		User:     "user",
		Password: "password",
		Port:     3406,
		Target:   "mydb",
	}))
	os.Setenv("POSTGRES_URL", CreateEnvValue(ServiceUrl{
		Type:     "postgres",
		User:     "user",
		Password: "password",
		Port:     5532,
		Target:   "mydb",
		Options:  "sslmode=disable",
	}))
	os.Setenv("MSSQL_URL", CreateEnvValue(ServiceUrl{
		Type:     "sqlserver",
		User:     "sa",
		Password: "password",
		Port:     1433,
		Target:   "test",
	}))
	os.Setenv("MONGODB_URL", CreateEnvValue(ServiceUrl{
		Type:   "mongo",
		Port:   27017,
		Target: "test",
	}))
	os.Setenv("SSO_TOKEN_URI", "http://localhost/tokenUri")
	os.Setenv("SSO_AUTH_URI", "http://localhost/authUri")
	os.Setenv("SSO_USER_INFO_URI", "http://localhost/userInfo")
	os.Setenv("SSO_CLIENT_ID", "myId")
	os.Setenv("SSO_CLIENT_SECRET", "mySecret")
	os.Setenv("SSO_GRANT_TYPE", "grant1,grant2")
	os.Setenv("SSO_SCOPES", "scope1,scope2")
	os.Setenv("REDIS_URL", CreateEnvValue(ServiceUrl{
		Type:     "redis",
		User:     "redis",
		Password: "redis",
		Port:     6379,
	}))
	os.Setenv("AMQP_URL", CreateEnvValue(ServiceUrl{
		Type:     "amqp",
		User:     "user",
		Password: "password",
		Port:     5672,
	}))
	os.Setenv("SMTP_URL", CreateEnvValue(ServiceUrl{
		Type: "smtp",
		Port: 587,
	}))
	os.Setenv("S3_URL", CreateEnvValue(ServiceUrl{
		Type:     "http",
		User:     "accessKey1",
		Password: "verySecretKey1",
		Port:     8090,
		Target:   "bucket",
	}))
	gautocloud.ReloadConnectors()

	Context("Mysql", func() {
		Context("client", func() {
			Context("By injection", func() {
				It("should inject a MysqlDB when use Get", func() {
					var db *dbtype.MysqlDB
					err := gautocloud.Inject(&db)
					Expect(err).ToNot(HaveOccurred())
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
				It("should inject a slice of MysqlDB when use Get and slice", func() {
					var dbs []*dbtype.MysqlDB
					err := gautocloud.Inject(&dbs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(dbs)).Should(Equal(1))
					db := dbs[0]
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
				It("should inject a MysqlDB when use GetWithId", func() {
					var db *dbtype.MysqlDB
					err := gautocloud.InjectFromId(mysql.MysqlConnector{}.Id(), &db)
					Expect(err).ToNot(HaveOccurred())
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
				It("should inject a slice of MysqlDB when use GetWithId and slice", func() {
					var dbs []*dbtype.MysqlDB
					err := gautocloud.InjectFromId(mysql.MysqlConnector{}.Id(), &dbs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(dbs)).Should(Equal(1))
					db := dbs[0]
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
			})
			Context("By return", func() {
				It("should return a MysqlDB when use GetFirst", func() {
					var db *dbtype.MysqlDB
					data, err := gautocloud.GetFirst(mysql.MysqlConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					db = data.(*dbtype.MysqlDB)
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
				It("should return a slice of MysqlDB when use GetAll", func() {
					var db *dbtype.MysqlDB
					data, err := gautocloud.GetAll(mysql.MysqlConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					Expect(len(data)).Should(Equal(1))
					db = data[0].(*dbtype.MysqlDB)
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
			})

		})
		Context("gorm", func() {
			It("should inject a gorm DB when user request it", func() {
				var gormDb *gorm.DB
				err := gautocloud.InjectFromId(gmysql.GormMysqlConnector{}.Id(), &gormDb)
				Expect(err).ToNot(HaveOccurred())
				Expect(gormDb).ShouldNot(BeNil())
				Expect(gormDb.Dialect().GetName()).Should(Equal("mysql"))
			})
		})
	})
	Context("Postgresql", func() {
		Context("client", func() {
			Context("By injection", func() {
				It("should inject a PostgresqlDB when use Get", func() {
					var db *dbtype.PostgresqlDB
					err := gautocloud.Inject(&db)
					Expect(err).ToNot(HaveOccurred())
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
				It("should inject a slice of PostgresqlDB when use Get and slice", func() {
					var dbs []*dbtype.PostgresqlDB
					err := gautocloud.Inject(&dbs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(dbs)).Should(Equal(1))
					db := dbs[0]
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
				It("should inject a PostgresqlDB when use GetWithId", func() {
					var db *dbtype.PostgresqlDB
					err := gautocloud.InjectFromId(postgresql.PostgresqlConnector{}.Id(), &db)
					Expect(err).ToNot(HaveOccurred())
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
				It("should inject a slice of PostgresqlDB when use GetWithId and slice", func() {
					var dbs []*dbtype.PostgresqlDB
					err := gautocloud.InjectFromId(postgresql.PostgresqlConnector{}.Id(), &dbs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(dbs)).Should(Equal(1))
					db := dbs[0]
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
			})
			Context("By return", func() {
				It("should return a PostgresqlDB when use GetFirst", func() {
					var db *dbtype.PostgresqlDB
					data, err := gautocloud.GetFirst(postgresql.PostgresqlConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					db = data.(*dbtype.PostgresqlDB)
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
				It("should return a slice of PostgresqlDB when use GetAll", func() {
					var db *dbtype.PostgresqlDB
					data, err := gautocloud.GetAll(postgresql.PostgresqlConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					Expect(len(data)).Should(Equal(1))
					db = data[0].(*dbtype.PostgresqlDB)
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
			})

		})
		Context("gorm", func() {
			It("should inject a gorm DB when user request it", func() {
				var gormDb *gorm.DB
				err := gautocloud.InjectFromId(gpostgres.GormPostgresqlConnector{}.Id(), &gormDb)
				Expect(err).ToNot(HaveOccurred())
				Expect(gormDb).ShouldNot(BeNil())
				Expect(gormDb.Dialect().GetName()).Should(Equal("postgres"))
			})
		})
	})
	Context("Mssql", func() {
		Context("client", func() {
			Context("By injection", func() {
				It("should inject a MssqlDB when use Get", func() {
					var db *dbtype.MssqlDB
					err := gautocloud.Inject(&db)
					Expect(err).ToNot(HaveOccurred())
					Expect(db).ShouldNot(BeNil())
				})
				It("should inject a slice of MssqlDB when use Get and slice", func() {
					var dbs []*dbtype.MssqlDB
					err := gautocloud.Inject(&dbs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(dbs)).Should(Equal(1))
					db := dbs[0]
					Expect(db).ShouldNot(BeNil())
				})
				It("should inject a MssqlDB when use GetWithId", func() {
					var db *dbtype.MssqlDB
					err := gautocloud.InjectFromId(mssql.MssqlConnector{}.Id(), &db)
					Expect(err).ToNot(HaveOccurred())
					Expect(db).ShouldNot(BeNil())
				})
				It("should inject a slice of MssqlDB when use GetWithId and slice", func() {
					var dbs []*dbtype.MssqlDB
					err := gautocloud.InjectFromId(mssql.MssqlConnector{}.Id(), &dbs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(dbs)).Should(Equal(1))
					db := dbs[0]
					Expect(db).ShouldNot(BeNil())
				})
			})
			Context("By return", func() {
				It("should return a MssqlDB when use GetFirst", func() {
					var db *dbtype.MssqlDB
					data, err := gautocloud.GetFirst(mssql.MssqlConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					db = data.(*dbtype.MssqlDB)
					Expect(db).ShouldNot(BeNil())
				})
				It("should return a slice of MssqlDB when use GetAll", func() {
					var db *dbtype.MssqlDB
					data, err := gautocloud.GetAll(mssql.MssqlConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					Expect(len(data)).Should(Equal(1))
					db = data[0].(*dbtype.MssqlDB)
					Expect(db).ShouldNot(BeNil())
				})
			})

		})
		Context("gorm", func() {
			It("should inject a gorm DB when user request it", func() {
				var gormDb *gorm.DB
				err := gautocloud.InjectFromId(gmssql.GormMssqlConnector{}.Id(), &gormDb)
				Expect(err).ToNot(HaveOccurred())
				Expect(gormDb).ShouldNot(BeNil())
			})
		})
	})
	Context("Redis", func() {
		Context("client", func() {
			Context("By injection", func() {
				It("should inject a redis.Client when use Get", func() {
					var db *redis.Client
					err := gautocloud.Inject(&db)
					Expect(err).ToNot(HaveOccurred())
					Expect(db).ShouldNot(BeNil())
					resp := db.Ping()
					Expect(resp).ShouldNot(BeNil())
					Expect(resp.Val()).Should(Equal("PONG"))
				})
				It("should inject a slice of redis.Client when use Get and slice", func() {
					var dbs []*redis.Client
					err := gautocloud.Inject(&dbs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(dbs)).Should(Equal(1))
					db := dbs[0]
					resp := db.Ping()
					Expect(resp).ShouldNot(BeNil())
					Expect(resp.Val()).Should(Equal("PONG"))
				})
				It("should inject a redis.Client when use GetWithId", func() {
					var db *redis.Client
					err := gautocloud.InjectFromId(credis.RedisConnector{}.Id(), &db)
					Expect(err).ToNot(HaveOccurred())
					Expect(db).ShouldNot(BeNil())
					resp := db.Ping()
					Expect(resp).ShouldNot(BeNil())
					Expect(resp.Val()).Should(Equal("PONG"))
				})
				It("should inject a slice of redis.Client when use GetWithId and slice", func() {
					var dbs []*redis.Client
					err := gautocloud.InjectFromId(credis.RedisConnector{}.Id(), &dbs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(dbs)).Should(Equal(1))
					db := dbs[0]
					Expect(db).ShouldNot(BeNil())
					resp := db.Ping()
					Expect(resp).ShouldNot(BeNil())
					Expect(resp.Val()).Should(Equal("PONG"))
				})
			})
			Context("By return", func() {
				It("should return a redis.Client when use GetFirst", func() {
					var db *redis.Client
					data, err := gautocloud.GetFirst(credis.RedisConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					db = data.(*redis.Client)
					Expect(db).ShouldNot(BeNil())
					resp := db.Ping()
					Expect(resp).ShouldNot(BeNil())
					Expect(resp.Val()).Should(Equal("PONG"))
				})
				It("should return a slice of redis.Client when use GetAll", func() {
					var db *redis.Client
					data, err := gautocloud.GetAll(credis.RedisConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					Expect(len(data)).Should(Equal(1))
					db = data[0].(*redis.Client)
					Expect(db).ShouldNot(BeNil())
					resp := db.Ping()
					Expect(resp).ShouldNot(BeNil())
					Expect(resp.Val()).Should(Equal("PONG"))
				})
			})

		})
	})
	Context("Mongodb", func() {
		Context("client", func() {
			Context("By injection", func() {
				It("should inject a mgo.Session when use Get", func() {
					var db *mgo.Session
					err := gautocloud.Inject(&db)
					Expect(err).ToNot(HaveOccurred())
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
				It("should inject a slice of mgo.Session when use Get and slice", func() {
					var dbs []*mgo.Session
					err := gautocloud.Inject(&dbs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(dbs)).Should(Equal(1))
					db := dbs[0]
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
				It("should inject a mgo.Session when use GetWithId", func() {
					var db *mgo.Session
					err := gautocloud.InjectFromId(mongodb.MongodbConnector{}.Id(), &db)
					Expect(err).ToNot(HaveOccurred())
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
				It("should inject a slice of mgo.Session when use GetWithId and slice", func() {
					var dbs []*mgo.Session
					err := gautocloud.InjectFromId(mongodb.MongodbConnector{}.Id(), &dbs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(dbs)).Should(Equal(1))
					db := dbs[0]
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
			})
			Context("By return", func() {
				It("should return a mgo.Session when use GetFirst", func() {
					var db *mgo.Session
					data, err := gautocloud.GetFirst(mongodb.MongodbConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					db = data.(*mgo.Session)
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
				It("should return a slice of mgo.Session when use GetAll", func() {
					var db *mgo.Session
					data, err := gautocloud.GetAll(mongodb.MongodbConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					Expect(len(data)).Should(Equal(1))
					db = data[0].(*mgo.Session)
					Expect(db).ShouldNot(BeNil())
					Expect(db.Ping()).ToNot(HaveOccurred())
				})
			})

		})
	})
	Context("Amqp", func() {
		Context("client", func() {
			Context("By injection", func() {
				It("should inject a amqp.Connection when use Get", func() {
					var svc *amqp.Connection
					err := gautocloud.Inject(&svc)
					Expect(err).ToNot(HaveOccurred())
					Expect(svc).ShouldNot(BeNil())
					_, err = svc.Channel()
					Expect(err).ToNot(HaveOccurred())
				})
				It("should inject a slice of amqp.Connection when use Get and slice", func() {
					var svcs []*amqp.Connection
					err := gautocloud.Inject(&svcs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(svcs)).Should(Equal(1))
					svc := svcs[0]
					Expect(svc).ShouldNot(BeNil())
					_, err = svc.Channel()
					Expect(err).ToNot(HaveOccurred())
				})
				It("should inject a amqp.Connection when use GetWithId", func() {
					var svc *amqp.Connection
					err := gautocloud.InjectFromId(camqp.AmqpConnector{}.Id(), &svc)
					Expect(err).ToNot(HaveOccurred())
					Expect(svc).ShouldNot(BeNil())
					_, err = svc.Channel()
					Expect(err).ToNot(HaveOccurred())
				})
				It("should inject a slice of amqp.Connection when use GetWithId and slice", func() {
					var svcs []*amqp.Connection
					err := gautocloud.InjectFromId(camqp.AmqpConnector{}.Id(), &svcs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(svcs)).Should(Equal(1))
					svc := svcs[0]
					Expect(svc).ShouldNot(BeNil())
					_, err = svc.Channel()
					Expect(err).ToNot(HaveOccurred())
				})
			})
			Context("By return", func() {
				It("should return a amqp.Connection when use GetFirst", func() {
					var svc *amqp.Connection
					data, err := gautocloud.GetFirst(camqp.AmqpConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					svc = data.(*amqp.Connection)
					Expect(svc).ShouldNot(BeNil())
					_, err = svc.Channel()
					Expect(err).ToNot(HaveOccurred())
				})
				It("should return a slice of amqp.Connection when use GetAll", func() {
					var svc *amqp.Connection
					data, err := gautocloud.GetAll(camqp.AmqpConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					Expect(len(data)).Should(Equal(1))
					svc = data[0].(*amqp.Connection)
					Expect(svc).ShouldNot(BeNil())
					_, err = svc.Channel()
					Expect(err).ToNot(HaveOccurred())
				})
			})

		})
	})
	Context("Smtp", func() {
		Context("client", func() {
			Context("By injection", func() {
				It("should inject a smtp.Client when use Get", func() {
					var svc *smtp.Client
					err := gautocloud.Inject(&svc)
					Expect(err).ToNot(HaveOccurred())
					Expect(svc).ShouldNot(BeNil())
					Expect(svc.Reset()).ShouldNot(HaveOccurred())
				})
				It("should inject a slice of smtp.Client when use Get and slice", func() {
					var svcs []*smtp.Client
					err := gautocloud.Inject(&svcs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(svcs)).Should(Equal(1))
					svc := svcs[0]
					Expect(svc).ShouldNot(BeNil())
					Expect(svc.Reset()).ShouldNot(HaveOccurred())
				})
				It("should inject a smtp.Client when use GetWithId", func() {
					var svc *smtp.Client
					err := gautocloud.InjectFromId(csmtp.SmtpConnector{}.Id(), &svc)
					Expect(err).ToNot(HaveOccurred())
					Expect(svc).ShouldNot(BeNil())
					Expect(svc.Reset()).ShouldNot(HaveOccurred())
				})
				It("should inject a slice of smtp.Client when use GetWithId and slice", func() {
					var svcs []*smtp.Client
					err := gautocloud.InjectFromId(csmtp.SmtpConnector{}.Id(), &svcs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(svcs)).Should(Equal(1))
					svc := svcs[0]
					Expect(svc).ShouldNot(BeNil())
					Expect(svc.Reset()).ShouldNot(HaveOccurred())
				})
			})
			Context("By return", func() {
				It("should return a smtp.Client when use GetFirst", func() {
					var svc *smtp.Client
					data, err := gautocloud.GetFirst(csmtp.SmtpConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					svc = data.(*smtp.Client)
					Expect(svc).ShouldNot(BeNil())
					Expect(svc.Reset()).ShouldNot(HaveOccurred())
				})
				It("should return a slice of smtp.Client when use GetAll", func() {
					var svc *smtp.Client
					data, err := gautocloud.GetAll(csmtp.SmtpConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					Expect(len(data)).Should(Equal(1))
					svc = data[0].(*smtp.Client)
					Expect(svc).ShouldNot(BeNil())
					Expect(svc.Reset()).ShouldNot(HaveOccurred())
				})
			})

		})
	})
	Context("S3", func() {
		Context("minio", func() {
			Context("By injection", func() {
				It("should inject a MinioClient when use Get", func() {
					var svc *miniotype.MinioClient
					err := gautocloud.Inject(&svc)
					Expect(err).ToNot(HaveOccurred())
					Expect(svc).ShouldNot(BeNil())
					Expect(svc.Client.MakeBucket(svc.Bucket, "")).ToNot(HaveOccurred())
					Expect(svc.Client.RemoveBucket(svc.Bucket)).ToNot(HaveOccurred())
				})
				It("should inject a slice of MinioClient when use Get and slice", func() {
					var svcs []*miniotype.MinioClient
					err := gautocloud.Inject(&svcs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(svcs)).Should(Equal(1))
					svc := svcs[0]
					Expect(svc).ShouldNot(BeNil())
					Expect(svc.Client.MakeBucket(svc.Bucket, "")).ToNot(HaveOccurred())
					Expect(svc.Client.RemoveBucket(svc.Bucket)).ToNot(HaveOccurred())
				})
				It("should inject a MinioClient when use GetWithId", func() {
					var svc *miniotype.MinioClient
					err := gautocloud.InjectFromId(cs3minio.MinioConnector{}.Id(), &svc)
					Expect(err).ToNot(HaveOccurred())
					Expect(svc).ShouldNot(BeNil())
					Expect(svc.Client.MakeBucket(svc.Bucket, "")).ToNot(HaveOccurred())
					Expect(svc.Client.RemoveBucket(svc.Bucket)).ToNot(HaveOccurred())
				})
				It("should inject a slice of MinioClient when use GetWithId and slice", func() {
					var svcs []*miniotype.MinioClient
					err := gautocloud.InjectFromId(cs3minio.MinioConnector{}.Id(), &svcs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(svcs)).Should(Equal(1))
					svc := svcs[0]
					Expect(svc).ShouldNot(BeNil())
					Expect(svc.Client.MakeBucket(svc.Bucket, "")).ToNot(HaveOccurred())
					Expect(svc.Client.RemoveBucket(svc.Bucket)).ToNot(HaveOccurred())
				})
			})
			Context("By return", func() {
				It("should return a MinioClient when use GetFirst", func() {
					var svc *miniotype.MinioClient
					data, err := gautocloud.GetFirst(cs3minio.MinioConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					svc = data.(*miniotype.MinioClient)
					Expect(svc).ShouldNot(BeNil())
					Expect(svc.Client.MakeBucket(svc.Bucket, "")).ToNot(HaveOccurred())
					Expect(svc.Client.RemoveBucket(svc.Bucket)).ToNot(HaveOccurred())
				})
				It("should return a slice of MinioClient when use GetAll", func() {
					var svc *miniotype.MinioClient
					data, err := gautocloud.GetAll(cs3minio.MinioConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					Expect(len(data)).Should(Equal(1))
					svc = data[0].(*miniotype.MinioClient)
					Expect(svc).ShouldNot(BeNil())
					Expect(svc.Client.MakeBucket(svc.Bucket, "")).ToNot(HaveOccurred())
					Expect(svc.Client.RemoveBucket(svc.Bucket)).ToNot(HaveOccurred())
				})
			})

		})
		Context("goamz", func() {
			BeforeEach(func() {
				var s3svcminio *miniotype.MinioClient
				err := gautocloud.Inject(&s3svcminio)
				if err != nil {
					Fail(err.Error())
				}
				err = s3svcminio.Client.MakeBucket(s3svcminio.Bucket, "")
				if err != nil {
					Fail(err.Error())
				}
			})
			AfterEach(func() {
				var s3svcminio *miniotype.MinioClient
				err := gautocloud.Inject(&s3svcminio)
				if err != nil {
					Fail(err.Error())
				}
				err = s3svcminio.Client.RemoveBucket(s3svcminio.Bucket)
				if err != nil {
					Fail(err.Error())
				}
			})
			Context("By injection", func() {
				It("should inject a s3.Bucket when use Get", func() {
					var svc *s3.Bucket
					err := gautocloud.Inject(&svc)
					Expect(err).ToNot(HaveOccurred())
					Expect(svc).ShouldNot(BeNil())
					err = svc.Put("myfile.txt", []byte("data"), "", s3.PublicRead, s3.Options{})
					Expect(err).ToNot(HaveOccurred())
					b, err := svc.Get("myfile.txt")
					Expect(string(b)).Should(Equal("data"))
					err = svc.Del("myfile.txt")
					Expect(err).ToNot(HaveOccurred())
				})
				It("should inject a slice of s3.Bucket when use Get and slice", func() {
					var svcs []*s3.Bucket
					err := gautocloud.Inject(&svcs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(svcs)).Should(Equal(1))
					svc := svcs[0]
					Expect(svc).ShouldNot(BeNil())
					err = svc.Put("myfile.txt", []byte("data"), "", s3.PublicRead, s3.Options{})
					Expect(err).ToNot(HaveOccurred())
					b, err := svc.Get("myfile.txt")
					Expect(string(b)).Should(Equal("data"))
					err = svc.Del("myfile.txt")
					Expect(err).ToNot(HaveOccurred())
				})
				It("should inject a s3.Bucket when use GetWithId", func() {
					var svc *s3.Bucket
					err := gautocloud.InjectFromId(cs3goamz.AmzS3Connector{}.Id(), &svc)
					Expect(err).ToNot(HaveOccurred())
					Expect(svc).ShouldNot(BeNil())
					err = svc.Put("myfile.txt", []byte("data"), "", s3.PublicRead, s3.Options{})
					Expect(err).ToNot(HaveOccurred())
					b, err := svc.Get("myfile.txt")
					Expect(string(b)).Should(Equal("data"))
					err = svc.Del("myfile.txt")
					Expect(err).ToNot(HaveOccurred())
				})
				It("should inject a slice of s3.Bucket when use GetWithId and slice", func() {
					var svcs []*s3.Bucket
					err := gautocloud.InjectFromId(cs3goamz.AmzS3Connector{}.Id(), &svcs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(svcs)).Should(Equal(1))
					svc := svcs[0]
					Expect(svc).ShouldNot(BeNil())
					err = svc.Put("myfile.txt", []byte("data"), "", s3.PublicRead, s3.Options{})
					Expect(err).ToNot(HaveOccurred())
					b, err := svc.Get("myfile.txt")
					Expect(string(b)).Should(Equal("data"))
					err = svc.Del("myfile.txt")
					Expect(err).ToNot(HaveOccurred())
				})
			})
			Context("By return", func() {
				It("should return a s3.Bucket when use GetFirst", func() {
					var svc *s3.Bucket
					data, err := gautocloud.GetFirst(cs3goamz.AmzS3Connector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					svc = data.(*s3.Bucket)
					Expect(svc).ShouldNot(BeNil())
					err = svc.Put("myfile.txt", []byte("data"), "", s3.PublicRead, s3.Options{})
					Expect(err).ToNot(HaveOccurred())
					b, err := svc.Get("myfile.txt")
					Expect(string(b)).Should(Equal("data"))
					err = svc.Del("myfile.txt")
					Expect(err).ToNot(HaveOccurred())
				})
				It("should return a slice of s3.Bucket when use GetAll", func() {
					var svc *s3.Bucket
					data, err := gautocloud.GetAll(cs3goamz.AmzS3Connector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					Expect(len(data)).Should(Equal(1))
					svc = data[0].(*s3.Bucket)
					Expect(svc).ShouldNot(BeNil())
					err = svc.Put("myfile.txt", []byte("data"), "", s3.PublicRead, s3.Options{})
					Expect(err).ToNot(HaveOccurred())
					b, err := svc.Get("myfile.txt")
					Expect(string(b)).Should(Equal("data"))
					err = svc.Del("myfile.txt")
					Expect(err).ToNot(HaveOccurred())
				})
			})

		})
	})
	Context("Oauth2", func() {
		Context("config", func() {
			Context("By injection", func() {
				It("should inject a oauth2.Config when use Get", func() {
					var svc *oauth2.Config
					err := gautocloud.Inject(&svc)
					Expect(err).ToNot(HaveOccurred())
					Expect(svc).ShouldNot(BeNil())

					Expect(svc.Scopes).Should(BeEquivalentTo([]string{"scope1", "scope2"}))
					Expect(svc.ClientSecret).Should(Equal("mySecret"))
					Expect(svc.ClientID).Should(Equal("myId"))
					Expect(svc.Endpoint.AuthURL).Should(Equal("http://localhost/authUri"))
					Expect(svc.Endpoint.TokenURL).Should(Equal("http://localhost/tokenUri"))
				})
				It("should inject a slice of oauth2.Config when use Get and slice", func() {
					var svcs []*oauth2.Config
					err := gautocloud.Inject(&svcs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(svcs)).Should(Equal(1))
					svc := svcs[0]
					Expect(svc).ShouldNot(BeNil())

					Expect(svc.Scopes).Should(BeEquivalentTo([]string{"scope1", "scope2"}))
					Expect(svc.ClientSecret).Should(Equal("mySecret"))
					Expect(svc.ClientID).Should(Equal("myId"))
					Expect(svc.Endpoint.AuthURL).Should(Equal("http://localhost/authUri"))
					Expect(svc.Endpoint.TokenURL).Should(Equal("http://localhost/tokenUri"))
				})
				It("should inject a oauth2.Config when use GetWithId", func() {
					var svc *oauth2.Config
					err := gautocloud.InjectFromId(coauth2.Oauth2ConfigConnector{}.Id(), &svc)
					Expect(err).ToNot(HaveOccurred())
					Expect(svc).ShouldNot(BeNil())

					Expect(svc.Scopes).Should(BeEquivalentTo([]string{"scope1", "scope2"}))
					Expect(svc.ClientSecret).Should(Equal("mySecret"))
					Expect(svc.ClientID).Should(Equal("myId"))
					Expect(svc.Endpoint.AuthURL).Should(Equal("http://localhost/authUri"))
					Expect(svc.Endpoint.TokenURL).Should(Equal("http://localhost/tokenUri"))
				})
				It("should inject a slice of oauth2.Config when use GetWithId and slice", func() {
					var svcs []*oauth2.Config
					err := gautocloud.InjectFromId(coauth2.Oauth2ConfigConnector{}.Id(), &svcs)
					Expect(err).ToNot(HaveOccurred())
					Expect(len(svcs)).Should(Equal(1))
					svc := svcs[0]
					Expect(svc).ShouldNot(BeNil())

					Expect(svc.Scopes).Should(BeEquivalentTo([]string{"scope1", "scope2"}))
					Expect(svc.ClientSecret).Should(Equal("mySecret"))
					Expect(svc.ClientID).Should(Equal("myId"))
					Expect(svc.Endpoint.AuthURL).Should(Equal("http://localhost/authUri"))
					Expect(svc.Endpoint.TokenURL).Should(Equal("http://localhost/tokenUri"))
				})
			})
			Context("By return", func() {
				It("should return a oauth2.Config when use GetFirst", func() {
					var svc *oauth2.Config
					data, err := gautocloud.GetFirst(coauth2.Oauth2ConfigConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					svc = data.(*oauth2.Config)
					Expect(svc).ShouldNot(BeNil())

					Expect(svc.Scopes).Should(BeEquivalentTo([]string{"scope1", "scope2"}))
					Expect(svc.ClientSecret).Should(Equal("mySecret"))
					Expect(svc.ClientID).Should(Equal("myId"))
					Expect(svc.Endpoint.AuthURL).Should(Equal("http://localhost/authUri"))
					Expect(svc.Endpoint.TokenURL).Should(Equal("http://localhost/tokenUri"))
				})
				It("should return a slice of oauth2.Config when use GetAll", func() {
					var svc *oauth2.Config
					data, err := gautocloud.GetAll(coauth2.Oauth2ConfigConnector{}.Id())
					Expect(err).ToNot(HaveOccurred())
					Expect(len(data)).Should(Equal(1))
					svc = data[0].(*oauth2.Config)
					Expect(svc).ShouldNot(BeNil())

					Expect(svc.Scopes).Should(BeEquivalentTo([]string{"scope1", "scope2"}))
					Expect(svc.ClientSecret).Should(Equal("mySecret"))
					Expect(svc.ClientID).Should(Equal("myId"))
					Expect(svc.Endpoint.AuthURL).Should(Equal("http://localhost/authUri"))
					Expect(svc.Endpoint.TokenURL).Should(Equal("http://localhost/tokenUri"))
				})
			})

		})
	})
	Context("Retrieving all gorm object by injection", func() {
		It("should inject a slice of gorm DB when user request it", func() {
			var gormDbs []*gorm.DB
			err := gautocloud.Inject(&gormDbs)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(gormDbs)).Should(Equal(3))
		})
	})
})
