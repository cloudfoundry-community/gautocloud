package schema

import "github.com/cloudfoundry-community/gautocloud/decoder"

type MongoDbSchema struct {
	Uri      decoder.ServiceUri `cloud:"ur(i|l),regex"`
	Database string             `cloud:"(database|db),regex"`
	Port     int                `cloud-default:"27017"`
	Host     string             `cloud:".*host.*,regex" cloud-default:"localhost"`
	User     string             `cloud:".*user.*,regex" cloud-default:"root"`
	Password string             `cloud:".*pass.*,regex"`
	Options  string             `cloud:"(options|opts),regex"`
}
