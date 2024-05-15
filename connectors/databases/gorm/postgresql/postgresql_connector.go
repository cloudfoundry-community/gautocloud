package postgresql

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/client/postgresql"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func init() {
	gautocloud.RegisterConnector(NewGormPostgresqlConnector())
}

type GormPostgresqlConnector struct {
	wrapConn    connectors.Connector
	wrapRawConn connectors.Connector
}

func NewGormPostgresqlConnector() connectors.Connector {
	return &GormPostgresqlConnector{
		wrapConn:    postgresql.NewPostgresqlConnector(),
		wrapRawConn: raw.NewPostgresqlRawConnector(),
	}
}
func (c GormPostgresqlConnector) Id() string {
	return "gorm:postgresql"
}
func (c GormPostgresqlConnector) Name() string {
	return c.wrapConn.Name()
}
func (c GormPostgresqlConnector) Tags() []string {
	return c.wrapConn.Tags()
}
func (c GormPostgresqlConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.wrapRawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(dbtype.PostgresqlDatabase)
	return gorm.Open("postgres", c.wrapConn.(*postgresql.PostgresqlConnector).GetConnString(fSchema))
}
func (c GormPostgresqlConnector) Schema() interface{} {
	return c.wrapConn.Schema()
}
