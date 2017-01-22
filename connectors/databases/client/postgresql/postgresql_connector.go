package postgresql

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)

func init() {
	gautocloud.RegisterConnector(NewPostgresqlConnector())
}

type PostgresqlConnector struct {
	rawConn connectors.Connector
}

func NewPostgresqlConnector() connectors.Connector {
	return &PostgresqlConnector{
		rawConn: raw.NewPostgresqlRawConnector(),
	}
}
func (c PostgresqlConnector) Id() string {
	return "postgresql"
}
func (c PostgresqlConnector) Name() string {
	return c.rawConn.Name()
}
func (c PostgresqlConnector) Tags() []string {
	return c.rawConn.Tags()
}
func (c PostgresqlConnector) GetConnString(schema dbtype.PostgresqlDatabase) string {
	connString := "postgres://" + schema.User
	if schema.Password != "" {
		connString += ":" + schema.Password
	}
	connString += fmt.Sprintf("@%s:%d/%s", schema.Host, schema.Port, schema.Database)
	if schema.Options != "" {
		connString += "?" + schema.Options
	}
	return connString
}
func (c PostgresqlConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.rawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", c.GetConnString(schema.(dbtype.PostgresqlDatabase)))
	if err != nil {
		return db, err
	}
	return &dbtype.PostgresqlDB{db}, nil
}
func (c PostgresqlConnector) Schema() interface{} {
	return c.rawConn.Schema()
}
