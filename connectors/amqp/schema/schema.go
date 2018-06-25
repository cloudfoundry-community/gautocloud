package schema

import "github.com/cloudfoundry-community/gautocloud/decoder"

type AmqpSchema struct {
	Uri      decoder.ServiceUri `cloud:"ur(i|l),regex"`
	Port     int                `cloud:"" cloud-default:"5672"`
	Host     string             `cloud:"hostname" cloud-default:"localhost"`
	User     string             `cloud:".*user.*,regex" cloud-default:"root"`
	Password string             `cloud:".*pass.*,regex"`
	Vhost    string             `cloud:"vhost.*,regex"`
}
