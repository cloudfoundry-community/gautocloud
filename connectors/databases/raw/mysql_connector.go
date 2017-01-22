package raw

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	. "github.com/cloudfoundry-community/gautocloud/connectors/databases/schema"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)

type MysqlRawConnector struct{}

func NewMysqlRawConnector() connectors.Connector {
	return &MysqlRawConnector{}
}
func (c MysqlRawConnector) Id() string {
	return "raw:mysql"
}
func (c MysqlRawConnector) Name() string {
	return ".*(mysql|maria).*"
}
func (c MysqlRawConnector) Tags() []string {
	return []string{"mysql", "maria.*"}
}
func (c MysqlRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(MysqlSchema)
	if fSchema.Uri.Host == "" {
		return dbtype.MysqlDatabase{
			User: fSchema.User,
			Password: fSchema.Password,
			Host: fSchema.Host,
			Port: fSchema.Port,
			Database: fSchema.Database,
		}, nil
	}
	port := fSchema.Uri.Port
	if port == 0 {
		port = fSchema.Port
	}
	return dbtype.MysqlDatabase{
		User: fSchema.Uri.Username,
		Password: fSchema.Uri.Password,
		Host: fSchema.Uri.Host,
		Port: port,
		Database: fSchema.Uri.Name,
		Options: fSchema.Uri.RawQuery,
	}, nil
}
func (c MysqlRawConnector) Schema() interface{} {
	return MysqlSchema{}
}