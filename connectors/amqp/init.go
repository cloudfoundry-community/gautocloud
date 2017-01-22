package amqp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/amqp/raw"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/amqp/client"
)

func init() {
	gautocloud.RegisterConnector(raw.NewAmqpRawConnector())
}
