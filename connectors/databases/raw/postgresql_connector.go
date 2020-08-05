package raw

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	. "github.com/cloudfoundry-community/gautocloud/connectors/databases/schema"
)

type PostgresqlRawConnector struct{}

func NewPostgresqlRawConnector() connectors.Connector {
	return &PostgresqlRawConnector{}
}
func (c PostgresqlRawConnector) Id() string {
	return "raw:postgresql"
}
func (c PostgresqlRawConnector) Name() string {
	return ".*postgres.*"
}
func (c PostgresqlRawConnector) Tags() []string {
	return []string{"postgres.*"}
}
func (c PostgresqlRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(PostgresqlSchema)
	if fSchema.Uri.Host == "" {
		return dbtype.PostgresqlDatabase{
			User:     fSchema.User,
			Password: fSchema.Password,
			Host:     fSchema.Host,
			Port:     fSchema.Port,
			Database: fSchema.Database,
			Options:  fSchema.Options,
		}, nil
	}
	port := fSchema.Uri.Port
	if port == 0 {
		port = fSchema.Port
	}
	return dbtype.PostgresqlDatabase{
		User:     fSchema.Uri.Username,
		Password: fSchema.Uri.Password,
		Host:     fSchema.Uri.Host,
		Port:     port,
		Database: fSchema.Uri.Name,
		Options:  fSchema.Uri.RawQuery,
	}, nil
}
func (c PostgresqlRawConnector) Schema() interface{} {
	return PostgresqlSchema{}
}
