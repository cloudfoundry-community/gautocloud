package aws_sdk

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype"
	"github.com/cloudfoundry-community/gautocloud/connectors/objstorage/raw"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"strconv"
)

func init() {
	gautocloud.RegisterConnector(NewAmzS3Connector())
}

type AmzS3Connector struct {
	wrapConn connectors.Connector
}

func NewAmzS3Connector() connectors.Connector {
	return &AmzS3Connector{
		wrapConn: raw.NewS3RawConnector(),
	}
}
func (c AmzS3Connector) Id() string {
	return "s3"
}
func (c AmzS3Connector) Name() string {
	return c.wrapConn.Name()
}
func (c AmzS3Connector) Tags() []string {
	return c.wrapConn.Tags()
}
func (c AmzS3Connector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.wrapConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(objstoretype.S3)
	scheme := "http"
	if fSchema.UseSsl {
		scheme = "https"
	}
	port := ""
	if fSchema.Port != 0 {
		port += ":" + strconv.Itoa(fSchema.Port)
	}
	forcedRegion := aws.Region{
		S3Endpoint: scheme + "://" + fSchema.Host + port,
	}
	auth := aws.Auth{
		AccessKey: fSchema.AccessKeyID,
		SecretKey: fSchema.SecretAccessKey,
	}
	conn := s3.New(auth, forcedRegion)
	return conn.Bucket(fSchema.Bucket), nil
}
func (c AmzS3Connector) Schema() interface{} {
	return c.wrapConn.Schema()
}
