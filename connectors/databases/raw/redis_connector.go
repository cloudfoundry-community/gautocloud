package raw

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	. "github.com/cloudfoundry-community/gautocloud/connectors/databases/schema"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)

type RedisRawConnector struct{}

func NewRedisRawConnector() connectors.Connector {
	return &RedisRawConnector{}
}
func (c RedisRawConnector) Id() string {
	return "raw:redis"
}
func (c RedisRawConnector) Name() string {
	return ".*redis.*"
}
func (c RedisRawConnector) Tags() []string {
	return []string{"redis"}
}
func (c RedisRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(RedisSchema)
	if fSchema.Uri.Host == "" {
		return dbtype.RedisDatabase{
			Password: fSchema.Password,
			Host: fSchema.Host,
			Port: fSchema.Port,
		}, nil
	}
	if fSchema.Uri.Username != "" {
		fSchema.Uri.Password = fSchema.Uri.Username
	}
	port := fSchema.Uri.Port
	if port == 0 {
		port = fSchema.Port
	}
	return dbtype.RedisDatabase{
		Password: fSchema.Uri.Password,
		Host: fSchema.Uri.Host,
		Port: port,
	}, nil
}
func (c RedisRawConnector) Schema() interface{} {
	return RedisSchema{}
}
