package client

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/amqp/amqptype"
	"github.com/cloudfoundry-community/gautocloud/connectors/amqp/raw"
	"github.com/streadway/amqp"
)

func init() {
	gautocloud.RegisterConnector(NewAmqpConnector())
}

type AmqpConnector struct {
	rawConn connectors.Connector
}

func NewAmqpConnector() connectors.Connector {
	return &AmqpConnector{
		rawConn: raw.NewAmqpRawConnector(),
	}
}
func (c AmqpConnector) Id() string {
	return "amqp"
}
func (c AmqpConnector) Name() string {
	return c.rawConn.Name()
}
func (c AmqpConnector) Tags() []string {
	return c.rawConn.Tags()
}
func (c AmqpConnector) GetConnString(schema amqptype.Amqp) string {
	connString := "amqp://" + schema.User
	if schema.Password != "" {
		connString += ":" + schema.Password
	}
	connString += fmt.Sprintf("@%s:%d/%s", schema.Host, schema.Port, schema.Vhost)
	return connString
}
func (c AmqpConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.rawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	return amqp.Dial(c.GetConnString(schema.(amqptype.Amqp)))
}
func (c AmqpConnector) Schema() interface{} {
	return c.rawConn.Schema()
}
