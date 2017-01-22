package fake

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"reflect"
)

type FakePtrConnector struct {
	schema interface{}
}

func NewFakePtrConnector(schema interface{}) connectors.Connector {
	fake := &FakePtrConnector{
		schema: schema,
	}
	return fake
}
func (f FakePtrConnector) Id() string {
	return "fake"
}
func (f FakePtrConnector) Name() string {
	return "fake"
}
func (f FakePtrConnector) Tags() []string {
	return []string{"service"}
}
func (f *FakePtrConnector) Load(schema interface{}) (interface{}, error) {
	vp := reflect.New(reflect.TypeOf(schema))
	vp.Elem().Set(reflect.ValueOf(schema))
	return vp.Interface(), nil
}
func (f FakePtrConnector) Schema() interface{} {
	return f.schema
}
func (f *FakePtrConnector) SetSchema(schema interface{}) {
	f.schema = schema
}