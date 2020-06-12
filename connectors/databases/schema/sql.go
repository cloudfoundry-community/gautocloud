package schema

import "github.com/cloudfoundry-community/gautocloud/decoder"

type MysqlSchema struct {
	Uri      decoder.ServiceUri `cloud:"ur(i|l),regex"`
	Database string `cloud:"(database|db)(_name)?,regex"`
	Port     int `cloud-default:"3306"`
	Host     string `cloud:".*host.*,regex" cloud-default:"localhost"`
	User     string `cloud:".*user.*,regex" cloud-default:"root"`
	Password string `cloud:".*pass.*,regex"`
}

type PostgresqlSchema struct {
	Uri      decoder.ServiceUri `cloud:"ur(i|l),regex"`
	Database string `cloud:"(database|db)(_name)?,regex"`
	Port     int `cloud-default:"5432"`
	Host     string `cloud:".*host.*,regex" cloud-default:"localhost"`
	User     string `cloud:".*user.*,regex" cloud-default:"root"`
	Password string `cloud:".*pass.*,regex"`
}
type MssqlSchema struct {
	Uri      decoder.ServiceUri `cloud:"ur(i|l),regex"`
	Database string `cloud:"(database|db)(_name)?,regex"`
	Port     int `cloud-default:"1433"`
	Host     string `cloud:".*host.*,regex" cloud-default:"localhost"`
	User     string `cloud:".*user.*,regex" cloud-default:"root"`
	Password string `cloud:".*pass.*,regex"`
}
type OracleSchema struct {
	Uri      decoder.ServiceUri `cloud:"ur(i|l),regex"`
	Database string `cloud:"(database|db)(_name)?,regex"`
	Port     int `cloud-default:"1521"`
	Host     string `cloud:".*host.*,regex" cloud-default:"localhost"`
	User     string `cloud:".*user.*,regex" cloud-default:"root"`
	Password string `cloud:".*pass.*,regex"`
}