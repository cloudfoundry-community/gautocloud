package auth

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/auth/raw"
)

func init() {
	gautocloud.RegisterConnector(raw.NewOauth2RawConnector())
}