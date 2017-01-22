package objstoretype

import "github.com/minio/minio-go"

type S3 struct {
	Host            string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
	Port            int
	UseSsl          bool
}

type MinioClient struct {
	Client *minio.Client
	Bucket string
}