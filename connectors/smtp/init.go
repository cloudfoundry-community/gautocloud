package smtp

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/smtp/raw"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/smtp/client"
)

func init() {
	gautocloud.RegisterConnector(raw.NewSmtpRawConnector())
}

