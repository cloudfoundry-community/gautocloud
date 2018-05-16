package raw

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/amqp/amqptype"
	. "github.com/cloudfoundry-community/gautocloud/connectors/amqp/schema"
)

type AmqpRawConnector struct{}

func NewAmqpRawConnector() connectors.Connector {
	return &AmqpRawConnector{}
}
func (c AmqpRawConnector) Id() string {
	return "raw:amqp"
}
func (c AmqpRawConnector) Name() string {
	return ".*amqp.*"
}
func (c AmqpRawConnector) Tags() []string {
	return []string{"amqp", "rabbitmq"}
}
func (c AmqpRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(AmqpSchema)
	if fSchema.Uri.Host == "" {
		return amqptype.Amqp{
			User:     fSchema.User,
			Password: fSchema.Password,
			Host:     fSchema.Host,
			Port:     fSchema.Port,
			Vhost:    fSchema.Vhost,
		}, nil
	}
	port := fSchema.Uri.Port
	if port == 0 {
		port = fSchema.Port
	}
	return amqptype.Amqp{
		User:     fSchema.Uri.Username,
		Password: fSchema.Uri.Password,
		Host:     fSchema.Uri.Host,
		Vhost:    fSchema.Uri.Name,
		Port:     port,
	}, nil
}
func (c AmqpRawConnector) Schema() interface{} {
	return AmqpSchema{}
}
