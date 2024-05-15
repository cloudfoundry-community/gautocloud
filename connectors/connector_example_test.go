package connectors_test

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/decoder"
)

// an init function in the same package of your connector and register it automatically in gautocloud when importing your connector
func init() {
	gautocloud.RegisterConnector(NewExampleConnector())
}

type ExampleSchema struct {
	// ServiceUri is a special type. Decoder will expect an uri as a value and will give a ServiceUri
	Uri decoder.ServiceUri
	// note: by default if you don't provide a cloud tag the key will be the field name in snake_case
	Name string `cloud:"name"`
	Host string `cloud:"host"`
	// by passing `regex` in cloud tag it will say to decoder that the expected key must be match the regex
	User string `cloud:".*user.*,regex"`
	// by passing `default=avalue` decoder will understand that if the key is not found it must fill the field with this value
	Password string `cloud:".*user.*,regex,default=apassword"`
	// you can also pass a slice
	Aslice []string `cloud:"aslice,default=value1,value2"`
}
type ExampleTypeOutput struct {
	Host     string
	Name     string
	User     string
	Password string
}
type ExampleConnector struct{}

func NewExampleConnector() connectors.Connector {
	return &ExampleConnector{}
}

// IThis is the id of your connector, and it must be unique and not have the same id of another connector
// Note: if a connector id is already taken gautocloud will complain
func (c ExampleConnector) Id() string {
	return "example"
}

// Name is the name of a service to lookup in the cloud environment
// Note: a regex can be passed
func (c ExampleConnector) Name() string {
	return ".*example.*"
}

// This should return a list of tags which designed what kind of service you want
// example: mysql, database, rmdb ...
// Note: a regex can be passed on each tag
func (c ExampleConnector) Tags() []string {
	return []string{"example", "doc.*"}
}

// The parameter is a filled schema you gave in the function Schema
// The first value to return is what you want, and you have no obligation to give always the same type. gautocloud is interface agnostic
// You can give an error if an error occurred, this error will appear in logs
func (c ExampleConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(ExampleSchema)
	if fSchema.Uri.Host != "" {
		return ExampleTypeOutput{
			Host:     fSchema.Uri.Host,
			Name:     fSchema.Uri.Name,
			User:     fSchema.Uri.Username,
			Password: fSchema.Uri.Password,
		}, nil
	}
	return ExampleTypeOutput{
		Host:     fSchema.Host,
		Name:     fSchema.Name,
		User:     fSchema.User,
		Password: fSchema.Password,
	}, nil
}

// It must return a structure
// this structure will be used by the decoder to create a structure of the same type and filled
// with service's credentials found by a cloud environment
func (c ExampleConnector) Schema() interface{} {
	return ExampleSchema{}
}
