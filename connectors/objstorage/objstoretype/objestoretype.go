package objstoretype

type S3 struct {
	Host            string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
	Port            int
	UseSsl          bool
}
