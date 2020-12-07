package miniotype

import "github.com/minio/minio-go/v7"

type MinioClient struct {
	Client *minio.Client
	Bucket string
}
