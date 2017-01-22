package schema

import "github.com/cloudfoundry-community/gautocloud/decoder"

type SmtpSchema struct {
	Uri      decoder.ServiceUri `cloud:"ur(i|l),regex"`
	Port     int `cloud:",default=587"`
	Host     string `cloud:".*host.*,regex,default=localhost"`
	User     string `cloud:".*user.*,regex"`
	Password string `cloud:".*pass.*,regex"`
}
