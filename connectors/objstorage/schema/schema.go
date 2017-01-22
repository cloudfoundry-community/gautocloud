package schema

import "github.com/cloudfoundry-community/gautocloud/decoder"

type S3Schema struct {
	Uri             decoder.ServiceUri `cloud:"ur(i|l),regex"`
	Host            string `cloud:".*host.*,regex"`
	AccessKeyID     string `cloud:"(.*user.*|access.*),regex"`
	SecretAccessKey string `cloud:"(.*pass.*|secret.*),regex"`
	Bucket          string `cloud:".*(bucket|name).*,regex"`
	Port            int
}
