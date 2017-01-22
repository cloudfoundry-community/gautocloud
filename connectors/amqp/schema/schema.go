package schema

import "github.com/cloudfoundry-community/gautocloud/decoder"

type AmqpSchema struct {
	Uri      decoder.ServiceUri `cloud:"ur(i|l),regex"`
	Port     int `cloud:",default=5672"`
	Host     string `cloud:".*host.*,regex,default=localhost"`
	User     string `cloud:".*user.*,regex,default=root"`
	Password string `cloud:".*pass.*,regex"`
}
