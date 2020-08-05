package raw

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	. "github.com/cloudfoundry-community/gautocloud/connectors/databases/schema"
)

type OracleRawConnector struct{}

func NewOracleRawConnector() connectors.Connector {
	return &OracleRawConnector{}
}
func (c OracleRawConnector) Id() string {
	return "raw:oracle"
}
func (c OracleRawConnector) Name() string {
	return ".*oracle.*"
}
func (c OracleRawConnector) Tags() []string {
	return []string{"oracle", "oci.*"}
}
func (c OracleRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(OracleSchema)
	if fSchema.Uri.Host == "" {
		return dbtype.OracleDatabase{
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
	return dbtype.OracleDatabase{
		User:     fSchema.Uri.Username,
		Password: fSchema.Uri.Password,
		Host:     fSchema.Uri.Host,
		Port:     port,
		Database: fSchema.Uri.Name,
		Options:  fSchema.Uri.RawQuery,
	}, nil
}
func (c OracleRawConnector) Schema() interface{} {
	return OracleSchema{}
}
