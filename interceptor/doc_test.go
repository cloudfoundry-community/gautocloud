package interceptor_test

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/generic"
)

func init() {
	os.Clearenv()
	err := os.Setenv("DYNO", "true")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("CONFIG_FOO", "gautocloud")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("CONFIG_BAR", "<injected by gautocloud>")
	if err != nil {
		panic(err)
	}
	gautocloud.RegisterConnector(generic.NewConfigGenericConnector(MyConfig{}))
	gautocloud.RegisterConnector(generic.NewConfigGenericConnector(MySchema{}))
	gautocloud.ReloadConnectors()
}

type MyConfig struct {
	Foo string
	Bar string
}

func ExampleNewOverwrite() {
	config := MyConfig{
		Foo: "my own data",
	}
	err := gautocloud.Inject(&config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", config)
	// Output: interceptor_test.MyConfig{Foo:"my own data", Bar:"<injected by gautocloud>"}
}

type MySchema struct {
	Foo string
	Bar string
}

func (s *MySchema) Intercept(found interface{}) error {
	f := found.(MySchema)
	s.Foo = "write"
	s.Bar = f.Bar
	return nil
}

func ExampleNewSchema() {
	var mySchema MySchema
	err := gautocloud.Inject(&mySchema)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", mySchema)
	// Output: interceptor_test.MySchema{Foo:"write", Bar:"<injected by gautocloud>"}
}
