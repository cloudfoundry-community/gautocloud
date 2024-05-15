package amqp

import (
	"github.com/cloudfoundry-community/gautocloud"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/amqp/client"
	"github.com/cloudfoundry-community/gautocloud/connectors/amqp/raw"
)

func init() {
	gautocloud.RegisterConnector(raw.NewAmqpRawConnector())
}
