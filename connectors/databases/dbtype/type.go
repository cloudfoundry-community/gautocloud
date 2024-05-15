package dbtype

import "database/sql"

type MssqlDB struct {
	*sql.DB
}
type MysqlDB struct {
	*sql.DB
}
type OracleDB struct {
	*sql.DB
}
type PostgresqlDB struct {
	*sql.DB
}

type MongodbDatabase struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Options  string
}
type MssqlDatabase struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Options  string
}
type MysqlDatabase struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Options  string
}
type OracleDatabase struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Options  string
}
type PostgresqlDatabase struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Options  string
}
type RedisDatabase struct {
	Password string
	Host     string
	Port     int
}
