package raw

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	. "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/schema"
	"github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype"
	"strings"
)

type S3RawConnector struct{}

func NewS3RawConnector() connectors.Connector {
	return &S3RawConnector{}
}
func (c S3RawConnector) Id() string {
	return "raw:s3"
}
func (c S3RawConnector) Name() string {
	return ".*s3.*"
}
func (c S3RawConnector) Tags() []string {
	return []string{"s3", "riak.*"}
}
func (c S3RawConnector) IsVirtualHostBucket(schema objstoretype.S3) bool {
	if schema.Bucket != "" {
		return false
	}
	if len(strings.Split(schema.Host, ".")) >= 3 {
		return true
	}
	return false
}
func (c S3RawConnector) GetBucketFromHost(host string) (endpoint string, bucket string) {
	splitHost := strings.Split(host, ".")
	bucket = splitHost[0]
	endpoint = strings.Join(splitHost[1:], ".")
	return
}
func (c S3RawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(S3Schema)
	var gSchema objstoretype.S3
	if fSchema.Uri.Host != "" {
		useSsl := false
		if fSchema.Uri.Scheme == "s3" || fSchema.Uri.Scheme == "https" {
			useSsl = true
		}
		return objstoretype.S3{
			Host: fSchema.Uri.Host,
			AccessKeyID: fSchema.Uri.Username,
			SecretAccessKey: fSchema.Uri.Password,
			Bucket: fSchema.Uri.Name,
			UseSsl: useSsl,
			Port: fSchema.Uri.Port,
		}, nil
	}
	gSchema = objstoretype.S3{
		Host: fSchema.Host,
		AccessKeyID: fSchema.AccessKeyID,
		SecretAccessKey: fSchema.SecretAccessKey,
		Bucket: fSchema.Bucket,
		Port: fSchema.Port,
		UseSsl: true,
	}
	if c.IsVirtualHostBucket(gSchema) {
		host, bucket := c.GetBucketFromHost(gSchema.Host)
		gSchema.Host = host
		gSchema.Bucket = bucket
	}
	return gSchema, nil
}
func (c S3RawConnector) Schema() interface{} {
	return S3Schema{}
}
