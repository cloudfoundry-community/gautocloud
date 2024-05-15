package mongodb

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
	"gopkg.in/mgo.v2"
)

func init() {
	gautocloud.RegisterConnector(NewMongodbConnector())
}

type MongodbConnector struct {
	rawConn connectors.Connector
}

func NewMongodbConnector() connectors.Connector {
	return &MongodbConnector{
		rawConn: raw.NewMongodbRawConnector(),
	}
}
func (c MongodbConnector) Id() string {
	return "mongodb"
}
func (c MongodbConnector) Name() string {
	return c.rawConn.Name()
}
func (c MongodbConnector) Tags() []string {
	return c.rawConn.Tags()
}
func (c MongodbConnector) GetConnString(schema dbtype.MongodbDatabase) string {
	connString := "mongodb://"
	if schema.User != "" {
		connString += schema.User
	}
	if schema.User != "" && schema.Password != "" {
		connString += ":" + schema.Password
	}
	if schema.User != "" {
		connString += "@"
	}
	connString += fmt.Sprintf("%s:%d/%s", schema.Host, schema.Port, schema.Database)
	if schema.Options != "" {
		connString += "?" + schema.Options
	}
	return connString
}
func (c MongodbConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.rawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	return mgo.Dial(c.GetConnString(schema.(dbtype.MongodbDatabase)))
}
func (c MongodbConnector) Schema() interface{} {
	return c.rawConn.Schema()
}
