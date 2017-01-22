package schema

import "github.com/cloudfoundry-community/gautocloud/decoder"

type RedisSchema struct {
	Uri      decoder.ServiceUri `cloud:"ur(i|l),regex"`
	Port     int `cloud:",default=27017"`
	Host     string `cloud:".*host.*,regex,default=localhost"`
	Password string `cloud:".*pass.*,regex"`
}


