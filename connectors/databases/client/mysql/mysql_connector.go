package mysql

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)

func init() {
	gautocloud.RegisterConnector(NewMysqlConnector())
}

type MysqlConnector struct {
	mysqlRawConn connectors.Connector
}

func NewMysqlConnector() connectors.Connector {
	return &MysqlConnector{
		mysqlRawConn: raw.NewMysqlRawConnector(),
	}
}
func (c MysqlConnector) Id() string {
	return "mysql"
}
func (c MysqlConnector) Name() string {
	return c.mysqlRawConn.Name()
}
func (c MysqlConnector) Tags() []string {
	return c.mysqlRawConn.Tags()
}
func (c MysqlConnector) GetConnString(schema dbtype.MysqlDatabase) string {
	connString := schema.User
	if schema.Password != "" {
		connString += ":" + schema.Password
	}
	connString += fmt.Sprintf("@tcp(%s:%d)/%s", schema.Host, schema.Port, schema.Database)
	// FIXME: error from mysql with param ?reconnect=true Error 1193: Unknown system variable 'reconnect'
	schema.Options = "parseTime=true"
	if schema.Options != "" {
		connString += "?" + schema.Options
	}
	return connString
}
func (c MysqlConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.mysqlRawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("mysql", c.GetConnString(schema.(dbtype.MysqlDatabase)))
	if err != nil {
		return db, err
	}
	return &dbtype.MysqlDB{db}, nil
}
func (c MysqlConnector) Schema() interface{} {
	return c.mysqlRawConn.Schema()
}