package generic_test

import (
	. "github.com/cloudfoundry-community/gautocloud/connectors/generic"

	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/interceptor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type SchemaFake struct {
	Foo string
}
type SchemaWithIntercepter struct {
	Foo string
}

func (s *SchemaWithIntercepter) Intercept(found interface{}) error {
	s.Foo = "hijack"
	return nil
}

var _ = Describe("Generic", func() {
	Context("SchemaBased", func() {
		Context("as intercepter", func() {
			It("should return an intercept function if schema implement intercepter", func() {
				conn := NewSchemaBasedGenericConnector("id", "name", []string{"tag"}, SchemaWithIntercepter{})
				intercepter := conn.(connectors.ConnectorIntercepter)

				interceptor := intercepter.Intercepter()

				Expect(interceptor).ToNot(BeNil())
				finalStruct, err := interceptor.Intercept(SchemaWithIntercepter{}, SchemaWithIntercepter{})
				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaWithIntercepter).Foo).To(Equal("hijack"))
			})
			It("should return an intercept function and run one by one interceptors if user set interceptors to connector", func() {
				alterFunc := interceptor.IntercepterFunc(func(current, found interface{}) (interface{}, error) {
					return SchemaFake{
						Foo: "altered",
					}, nil
				})
				addFunc := interceptor.IntercepterFunc(func(current, found interface{}) (interface{}, error) {
					c := current.(SchemaFake)
					return SchemaFake{
						Foo: c.Foo + "added",
					}, nil
				})
				conn := NewSchemaBasedGenericConnector(
					"id",
					"name",
					[]string{"tag"},
					SchemaFake{},
					alterFunc,
					addFunc,
				)
				intercepter := conn.(connectors.ConnectorIntercepter)

				interceptor := intercepter.Intercepter()

				Expect(interceptor).ToNot(BeNil())
				finalStruct, err := interceptor.Intercept(SchemaFake{}, SchemaFake{})
				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaFake).Foo).To(Equal("alteredadded"))
			})
			It("should return nil if schema does not implement intercepter", func() {
				conn := NewSchemaBasedGenericConnector("id", "name", []string{"tag"}, SchemaFake{})
				interceptor := conn.(connectors.ConnectorIntercepter)

				interceptFunc := interceptor.Intercepter()

				Expect(interceptFunc).To(BeNil())
			})
		})

	})
	Context("Config", func() {
		Context("as intercepter", func() {
			It("should return an intercept function if schema implement intercepter", func() {
				conn := NewConfigGenericConnector(SchemaWithIntercepter{})
				interceptor := conn.(connectors.ConnectorIntercepter)

				intercept := interceptor.Intercepter()

				Expect(intercept).ToNot(BeNil())
				finalStruct, err := intercept.Intercept(SchemaWithIntercepter{}, SchemaWithIntercepter{})
				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaWithIntercepter).Foo).To(Equal("hijack"))
			})
			It("should return a default function to overwrite found interface with value from interface given by user", func() {
				conn := NewConfigGenericConnector(SchemaFake{})
				interceptor := conn.(connectors.ConnectorIntercepter)

				interceptFunc := interceptor.Intercepter()

				Expect(interceptFunc).ToNot(BeNil())
			})
		})
	})
})
