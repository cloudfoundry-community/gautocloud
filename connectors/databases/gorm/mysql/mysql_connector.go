package mysql

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mysql"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
)

func init() {
	gautocloud.RegisterConnector(NewGormMysqlConnector())
}

type GormMysqlConnector struct {
	wrapConn    connectors.Connector
	wrapRawConn connectors.Connector
}

func NewGormMysqlConnector() connectors.Connector {
	return &GormMysqlConnector{
		wrapConn: mysql.NewMysqlConnector(),
		wrapRawConn: raw.NewMysqlRawConnector(),
	}
}
func (c GormMysqlConnector) Id() string {
	return "gorm:mysql"
}
func (c GormMysqlConnector) Name() string {
	return c.wrapConn.Name()
}
func (c GormMysqlConnector) Tags() []string {
	return c.wrapConn.Tags()
}
func (c GormMysqlConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.wrapRawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(dbtype.MysqlDatabase)
	return gorm.Open("mysql", c.wrapConn.(*mysql.MysqlConnector).GetConnString(fSchema))
}
func (c GormMysqlConnector) Schema() interface{} {
	return c.wrapConn.Schema()
}
