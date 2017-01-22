package raw

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	. "github.com/cloudfoundry-community/gautocloud/connectors/databases/schema"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)

type MssqlRawConnector struct{}

func NewMssqlRawConnector() connectors.Connector {
	return &MssqlRawConnector{}
}
func (c MssqlRawConnector) Id() string {
	return "raw:mssql"
}
func (c MssqlRawConnector) Name() string {
	return ".*mssql.*"
}
func (c MssqlRawConnector) Tags() []string {
	return []string{"mssql.*", "sqlserver"}
}
func (c MssqlRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(MssqlSchema)
	if fSchema.Uri.Host == "" {
		return dbtype.MssqlDatabase{
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
	return dbtype.MssqlDatabase{
		User: fSchema.Uri.Username,
		Password: fSchema.Uri.Password,
		Host: fSchema.Uri.Host,
		Port: port,
		Database: fSchema.Uri.Name,
		Options: fSchema.Uri.RawQuery,
	}, nil
}
func (c MssqlRawConnector) Schema() interface{} {
	return MssqlSchema{}
}
