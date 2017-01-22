package minio

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/objstorage/raw"
	"github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype"
	"github.com/minio/minio-go"
	"strconv"
)

func init() {
	gautocloud.RegisterConnector(NewMinioConnector())
}

type MinioConnector struct {
	wrapConn connectors.Connector
}

func NewMinioConnector() connectors.Connector {
	return &MinioConnector{
		wrapConn: raw.NewS3RawConnector(),
	}
}
func (c MinioConnector) Id() string {
	return "minio:s3"
}
func (c MinioConnector) Name() string {
	return c.wrapConn.Name()
}
func (c MinioConnector) Tags() []string {
	return c.wrapConn.Tags()
}
func (c MinioConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.wrapConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(objstoretype.S3)
	port := ""
	if fSchema.Port != 0 {
		port += ":" + strconv.Itoa(fSchema.Port)
	}
	minioClient, err := minio.New(fSchema.Host + port, fSchema.AccessKeyID, fSchema.SecretAccessKey, fSchema.UseSsl)
	if err != nil {
		return nil, err
	}

	return &objstoretype.MinioClient{
		Client: minioClient,
		Bucket: fSchema.Bucket,
	}, nil
}
func (c MinioConnector) Schema() interface{} {
	return c.wrapConn.Schema()
}
