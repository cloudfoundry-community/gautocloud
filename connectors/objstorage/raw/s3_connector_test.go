package raw_test

import (
	. "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/raw"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/decoder"
	"github.com/cloudfoundry-community/gautocloud/connectors/objstorage/schema"
	"github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype"
)

var _ = Describe("S3Connector", func() {
	var connector connectors.Connector
	BeforeEach(func() {
		connector = NewS3RawConnector()
	})
	It("Should return a S3 struct when passing a S3Schema without uri", func() {
		data, err := connector.Load(schema.S3Schema{
			Host: "localhost",
			SecretAccessKey: "pass",
			AccessKeyID: "user",
			Bucket: "bucket",
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(data).Should(BeEquivalentTo(
			objstoretype.S3{
				Host: "localhost",
				SecretAccessKey: "pass",
				AccessKeyID: "user",
				Bucket: "bucket",
				UseSsl: true,
			},
		))
	})
	Context("without use ssl in uri", func() {
		It("Should return a S3 struct when passing a S3Schema", func() {
			data, err := connector.Load(schema.S3Schema{
				Uri: decoder.ServiceUri{
					Scheme: "http",
					Host: "localhost",
					Username: "user",
					Password: "pass",
					Name: "bucket",
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				objstoretype.S3{
					Host: "localhost",
					SecretAccessKey: "pass",
					AccessKeyID: "user",
					Bucket: "bucket",
				},
			))
		})
	})
	Context("with use ssl in uri", func() {
		It("Should return a S3 struct when passing a S3Schema with scheme https", func() {
			data, err := connector.Load(schema.S3Schema{
				Uri: decoder.ServiceUri{
					Scheme: "https",
					Host: "localhost",
					Username: "user",
					Password: "pass",
					Name: "bucket",
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				objstoretype.S3{
					Host: "localhost",
					SecretAccessKey: "pass",
					AccessKeyID: "user",
					Bucket: "bucket",
					UseSsl: true,
				},
			))
		})
		It("Should return a S3 struct when passing a S3Schema with scheme s3", func() {
			data, err := connector.Load(schema.S3Schema{
				Uri: decoder.ServiceUri{
					Scheme: "s3",
					Host: "localhost",
					Username: "user",
					Password: "pass",
					Name: "bucket",
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).Should(BeEquivalentTo(
				objstoretype.S3{
					Host: "localhost",
					SecretAccessKey: "pass",
					AccessKeyID: "user",
					Bucket: "bucket",
					UseSsl: true,
				},
			))
		})
	})

})
