package redis

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
	"github.com/go-redis/redis/v7"
	"github.com/cloudfoundry-community/gautocloud"
	"strconv"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)

func init() {
	gautocloud.RegisterConnector(NewRedisConnector())
}

type RedisConnector struct {
	rawConn connectors.Connector
}

func NewRedisConnector() connectors.Connector {
	return &RedisConnector{
		rawConn: raw.NewRedisRawConnector(),
	}
}
func (c RedisConnector) Id() string {
	return "redis"
}
func (c RedisConnector) Name() string {
	return c.rawConn.Name()
}
func (c RedisConnector) Tags() []string {
	return c.rawConn.Tags()
}
func (c RedisConnector) GetConnString(schema dbtype.RedisDatabase) string {
	return schema.Host + ":" + strconv.Itoa(schema.Port)
}
func (c RedisConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.rawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(dbtype.RedisDatabase)
	return redis.NewClient(&redis.Options{
		Addr: c.GetConnString(fSchema),
		Password: fSchema.Password,
	}), nil
}
func (c RedisConnector) Schema() interface{} {
	return c.rawConn.Schema()
}


