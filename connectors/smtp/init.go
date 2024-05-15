package smtp

import (
	"github.com/cloudfoundry-community/gautocloud"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/smtp/client"
	"github.com/cloudfoundry-community/gautocloud/connectors/smtp/raw"
)

func init() {
	gautocloud.RegisterConnector(raw.NewSmtpRawConnector())
}
