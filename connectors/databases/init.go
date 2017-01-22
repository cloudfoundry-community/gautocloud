package databases

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
)

func init() {
	gautocloud.RegisterConnector(raw.NewMongodbRawConnector())
	gautocloud.RegisterConnector(raw.NewMssqlRawConnector())
	gautocloud.RegisterConnector(raw.NewMysqlRawConnector())
	gautocloud.RegisterConnector(raw.NewOracleRawConnector())
	gautocloud.RegisterConnector(raw.NewPostgresqlRawConnector())
	gautocloud.RegisterConnector(raw.NewRedisRawConnector())
}
