package interceptor_test

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/generic"
	"os"
)

func init() {
	os.Clearenv()
	os.Setenv("DYNO", "true")
	os.Setenv("CONFIG_FOO", "gautocloud")
	os.Setenv("CONFIG_BAR", "<injected by gautocloud>")
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
	fmt.Println(fmt.Sprintf("%#v", config))
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
	gautocloud.Inject(&mySchema)
	fmt.Println(fmt.Sprintf("%#v", mySchema))
	// Output: interceptor_test.MySchema{Foo:"write", Bar:"<injected by gautocloud>"}
}
