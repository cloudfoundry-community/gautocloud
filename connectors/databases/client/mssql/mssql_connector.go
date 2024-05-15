package mssql

import (
	"database/sql"
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
	_ "github.com/denisenkom/go-mssqldb"
)

func init() {
	gautocloud.RegisterConnector(NewMssqlConnector())
}

type MssqlConnector struct {
	rawConn connectors.Connector
}

func NewMssqlConnector() connectors.Connector {
	return &MssqlConnector{
		rawConn: raw.NewMssqlRawConnector(),
	}
}
func (c MssqlConnector) Id() string {
	return "mssql"
}
func (c MssqlConnector) Name() string {
	return c.rawConn.Name()
}
func (c MssqlConnector) Tags() []string {
	return c.rawConn.Tags()
}
func (c MssqlConnector) GetConnString(schema dbtype.MssqlDatabase) string {
	connString := "sqlserver://" + schema.User
	if schema.Password != "" {
		connString += ":" + schema.Password
	}
	connString += fmt.Sprintf("@%s:%d?database=%s", schema.Host, schema.Port, schema.Database)
	if schema.Options != "" {
		connString += "&" + schema.Options
	}
	return connString
}
func (c MssqlConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.rawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("mssql", c.GetConnString(schema.(dbtype.MssqlDatabase)))
	if err != nil {
		return db, err
	}
	return &dbtype.MssqlDB{db}, nil
}
func (c MssqlConnector) Schema() interface{} {
	return c.rawConn.Schema()
}
