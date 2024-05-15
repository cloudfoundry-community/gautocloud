package mssql

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mssql"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func init() {
	gautocloud.RegisterConnector(NewGormMssqlConnector())
}

type GormMssqlConnector struct {
	wrapConn    connectors.Connector
	wrapRawConn connectors.Connector
}

func NewGormMssqlConnector() connectors.Connector {
	return &GormMssqlConnector{
		wrapConn:    mssql.NewMssqlConnector(),
		wrapRawConn: raw.NewMssqlRawConnector(),
	}
}
func (c GormMssqlConnector) Id() string {
	return "gorm:mssql"
}
func (c GormMssqlConnector) Name() string {
	return c.wrapConn.Name()
}
func (c GormMssqlConnector) Tags() []string {
	return c.wrapConn.Tags()
}
func (c GormMssqlConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.wrapRawConn.Load(schema)
	if err != nil {
		return &gorm.DB{}, err
	}
	fSchema := schema.(dbtype.MssqlDatabase)
	data, err := gorm.Open("mssql", c.wrapConn.(*mssql.MssqlConnector).GetConnString(fSchema))
	if err != nil {
		return &gorm.DB{}, nil
	}
	return data, nil
}
func (c GormMssqlConnector) Schema() interface{} {
	return c.wrapConn.Schema()
}
