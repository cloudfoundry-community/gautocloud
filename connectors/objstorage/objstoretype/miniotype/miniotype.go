package minio

import "github.com/minio/minio-go"

type MinioClient struct {
	Client *minio.Client
	Bucket string
}
