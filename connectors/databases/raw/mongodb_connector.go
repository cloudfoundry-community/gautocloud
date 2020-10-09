package raw

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	. "github.com/cloudfoundry-community/gautocloud/connectors/databases/schema"
)

type MongodbRawConnector struct{}

func NewMongodbRawConnector() connectors.Connector {
	return &MongodbRawConnector{}
}
func (c MongodbRawConnector) Id() string {
	return "raw:mongodb"
}
func (c MongodbRawConnector) Name() string {
	return ".*mongo.*"
}
func (c MongodbRawConnector) Tags() []string {
	return []string{"mongo.*"}
}
func (c MongodbRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(MongoDbSchema)
	if fSchema.Uri.Host == "" {
		return dbtype.MongodbDatabase{
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
	return dbtype.MongodbDatabase{
		User:     fSchema.Uri.Username,
		Password: fSchema.Uri.Password,
		Host:     fSchema.Uri.Host,
		Port:     port,
		Database: fSchema.Uri.Name,
		Options:  fSchema.Uri.RawQuery,
	}, nil
}
func (c MongodbRawConnector) Schema() interface{} {
	return MongoDbSchema{}
}
