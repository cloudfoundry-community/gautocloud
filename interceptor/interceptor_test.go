package interceptor_test

import (
	. "github.com/cloudfoundry-community/gautocloud/interceptor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type SchemaFake struct {
	Foo      string
	MyMap    map[string]interface{}
	MySlice  []string
	MyStruct MyStruct
	MyPtr    *string
}
type MyStruct struct {
	Bar string
}

type SchemaFakePtrIntercepter struct {
	Foo    string
	MyFunc func(*SchemaFakePtrIntercepter) error
}

func (s *SchemaFakePtrIntercepter) Intercept(found interface{}) error {
	s.Foo = "hijack"
	if s.MyFunc == nil {
		return nil
	}
	return s.MyFunc(s)
}

type SchemaFakeIntercepter struct {
	Foo    string
	MyFunc func(SchemaFakeIntercepter) error
}

func (s SchemaFakeIntercepter) Intercept(found interface{}) error {
	s.Foo = "hijack"
	if s.MyFunc == nil {
		return nil
	}
	return s.MyFunc(s)
}

var _ = Describe("Interceptor", func() {
	Context("Schema", func() {
		Context("Schema passed is not a pointer", func() {
			It("should modified schema when schema pointer implements SchemaIntercepter", func() {
				finalStruct, err := NewSchema().Intercept(SchemaFakePtrIntercepter{
					Foo: "data",
				}, SchemaFakePtrIntercepter{})

				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaFakePtrIntercepter).Foo).To(Equal("hijack"))
			})
			It("should run intercept function if schema implements SchemaIntercepter", func() {
				schema := SchemaFakeIntercepter{
					Foo: "data",
				}
				schema.MyFunc = func(s SchemaFakeIntercepter) error {
					Expect(s.Foo).To(Equal("hijack"))
					return nil
				}

				finalStruct, err := NewSchema().Intercept(schema, SchemaFakeIntercepter{})

				Expect(err).ToNot(HaveOccurred())
				// Intercept function does not ask for pointer, this is normal that Foo is not override
				Expect(finalStruct.(SchemaFakeIntercepter).Foo).To(Equal("data"))
			})
			It("should return an error if schema does not implement SchemaIntercepter", func() {
				_, err := NewSchema().Intercept(SchemaFake{
					Foo: "data",
				}, SchemaFake{})

				Expect(err).To(HaveOccurred())
			})
		})
		Context("Schema passed is a pointer", func() {
			It("should modified schema when schema pointer implements SchemaIntercepter", func() {
				finalStruct, err := NewSchema().Intercept(&SchemaFakePtrIntercepter{
					Foo: "data",
				}, &SchemaFakePtrIntercepter{})

				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(*SchemaFakePtrIntercepter).Foo).To(Equal("hijack"))
			})
			It("should run intercept function if schema implements SchemaIntercepter", func() {
				schema := &SchemaFakeIntercepter{
					Foo: "data",
				}
				schema.MyFunc = func(s SchemaFakeIntercepter) error {
					Expect(s.Foo).To(Equal("hijack"))
					return nil
				}

				finalStruct, err := NewSchema().Intercept(schema, &SchemaFakeIntercepter{})

				Expect(err).ToNot(HaveOccurred())
				// Intercept function does not ask for pointer, this is normal that Foo is not override
				Expect(finalStruct.(*SchemaFakeIntercepter).Foo).To(Equal("data"))
			})
			It("should return an error if schema does not implement SchemaIntercepter", func() {
				_, err := NewSchema().Intercept(&SchemaFake{
					Foo: "data",
				}, &SchemaFake{})

				Expect(err).To(HaveOccurred())
			})
		})
	})
	Context("Overwrite", func() {
		It("should return found value directly if current is nil", func() {
			finalStruct, err := NewOverwrite().Intercept(nil, SchemaFake{
				Foo: "decoded",
			})

			Expect(err).ToNot(HaveOccurred())
			Expect(finalStruct.(SchemaFake).Foo).To(Equal("decoded"))
		})
		It("should be agnostic to pointer if pointer to schema is given", func() {

			finalStruct, err := NewOverwrite().Intercept(&SchemaFake{
				Foo: "overwrite",
			}, &SchemaFake{})

			Expect(err).ToNot(HaveOccurred())
			Expect(finalStruct.(*SchemaFake).Foo).To(Equal("overwrite"))
		})
		Context("value is a primitive type", func() {
			It("should return a final interface with values from interface given by user", func() {

				finalStruct, err := NewOverwrite().Intercept(SchemaFake{
					Foo: "overwrite",
				}, SchemaFake{})

				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaFake).Foo).To(Equal("overwrite"))
			})
			It("should return a final interface with value from interface given by connector if fields are 0 values", func() {
				finalStruct, err := NewOverwrite().Intercept(SchemaFake{}, SchemaFake{
					Foo: "decoded",
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaFake).Foo).To(Equal("decoded"))
			})
		})
		Context("value is a slice", func() {
			It("should return a final interface with values from interface given by user", func() {
				finalStruct, err := NewOverwrite().Intercept(SchemaFake{
					MySlice: []string{"overwrite1", "overwrite2"},
				}, SchemaFake{})

				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaFake).MySlice[0]).To(Equal("overwrite1"))
				Expect(finalStruct.(SchemaFake).MySlice[1]).To(Equal("overwrite2"))
			})
			It("should return a final interface with value from interface given by connector if fields are 0 values", func() {
				finalStruct, err := NewOverwrite().Intercept(SchemaFake{}, SchemaFake{
					MySlice: []string{"decoded1", "decoded2"},
				})

				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaFake).MySlice[0]).To(Equal("decoded1"))
				Expect(finalStruct.(SchemaFake).MySlice[1]).To(Equal("decoded2"))
			})
		})

		Context("value is a map", func() {
			It("should return a final interface with values from interface given by user", func() {
				finalStruct, err := NewOverwrite().Intercept(SchemaFake{
					MyMap: map[string]interface{}{
						"overwrite1": "data",
						"overwrite2": "data",
					},
				}, SchemaFake{})

				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaFake).MyMap).To(HaveKey("overwrite1"))
				Expect(finalStruct.(SchemaFake).MyMap).To(HaveKey("overwrite2"))
			})
			It("should return a final interface with value from interface given by connector if fields are 0 values", func() {
				finalStruct, err := NewOverwrite().Intercept(SchemaFake{}, SchemaFake{
					MyMap: map[string]interface{}{
						"decoded1": "data",
						"decoded2": "data",
					},
				})

				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaFake).MyMap).To(HaveKey("decoded1"))
				Expect(finalStruct.(SchemaFake).MyMap).To(HaveKey("decoded2"))
			})
		})

		Context("value is a struct", func() {
			It("should return a final interface with values from interface given by user", func() {
				finalStruct, err := NewOverwrite().Intercept(SchemaFake{
					MyStruct: MyStruct{
						Bar: "overwrite",
					},
				}, SchemaFake{})

				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaFake).MyStruct.Bar).To(Equal("overwrite"))
			})
			It("should return a final interface with value from interface given by connector if fields are 0 values", func() {
				finalStruct, err := NewOverwrite().Intercept(SchemaFake{}, SchemaFake{
					MyStruct: MyStruct{
						Bar: "decoded",
					},
				})

				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaFake).MyStruct.Bar).To(Equal("decoded"))
			})
		})
		Context("value is a pointer", func() {
			It("should return a final interface with values from interface given by user", func() {
				value := "overwrite"
				finalStruct, err := NewOverwrite().Intercept(SchemaFake{
					MyPtr: &value,
				}, SchemaFake{})

				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaFake).MyPtr).ShouldNot(BeNil())
				Expect(*finalStruct.(SchemaFake).MyPtr).To(Equal("overwrite"))
			})
			It("should return a final interface with value from interface given by connector if fields are 0 values", func() {
				value := "decoded"
				finalStruct, err := NewOverwrite().Intercept(SchemaFake{}, SchemaFake{
					MyPtr: &value,
				})

				Expect(err).ToNot(HaveOccurred())
				Expect(finalStruct.(SchemaFake).MyPtr).ShouldNot(BeNil())
				Expect(*finalStruct.(SchemaFake).MyPtr).To(Equal("decoded"))
			})
		})
	})
})
