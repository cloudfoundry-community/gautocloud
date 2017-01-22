package objstorage

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/objstorage/raw"
)

func init() {
	gautocloud.RegisterConnector(raw.NewS3RawConnector())
}
