package fake

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/interceptor"
)

type FakeInterceptor struct {
	schema      interface{}
	interceptor interceptor.Intercepter
}

func NewFakeInterceptor(schema interface{}, interceptFunc interceptor.Intercepter) connectors.Connector {
	return &FakeInterceptor{
		schema:      schema,
		interceptor: interceptFunc,
	}
}
func (f FakeInterceptor) Id() string {
	return "fake"
}
func (f FakeInterceptor) Name() string {
	return ".*fake.*"
}
func (f FakeInterceptor) Tags() []string {
	return []string{"service"}
}
func (f *FakeInterceptor) Load(schema interface{}) (interface{}, error) {
	return schema, nil
}
func (f FakeInterceptor) Schema() interface{} {
	return f.schema
}
func (f *FakeInterceptor) SetSchema(schema interface{}) {
	f.schema = schema
}
func (f *FakeInterceptor) SetInterceptFunc(interceptFunc interceptor.Intercepter) {
	f.interceptor = interceptFunc
}
func (f FakeInterceptor) Intercepter() interceptor.Intercepter {
	return f.interceptor
}
