package fake

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
)

type FakeConnector struct {
	schema interface{}
}

func NewFakeConnector(schema interface{}) connectors.Connector {
	return &FakeConnector{
		schema: schema,
	}
}
func (f FakeConnector) Id() string {
	return "fake"
}
func (f FakeConnector) Name() string {
	return ".*fake.*"
}
func (f FakeConnector) Tags() []string {
	return []string{"service"}
}
func (f *FakeConnector) Load(schema interface{}) (interface{}, error) {
	return schema, nil
}
func (f FakeConnector) Schema() interface{} {
	return f.schema
}
func (f *FakeConnector) SetSchema(schema interface{}) {
	f.schema = schema
}
