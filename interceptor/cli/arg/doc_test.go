package arg_test

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/generic"
	. "github.com/cloudfoundry-community/gautocloud/interceptor/cli/arg"
)

func Example() {
	// This is only to fake args from console
	// Normally you don't have to set this option
	argsOpt := Args([]string{
		"app",
		`--foo=bar`,
		"--bar",
	})
	argInterceptor := NewArg(argsOpt)
	type MyConfig struct {
		Foo  string
		Bar  bool
		Orig string
	}

	// Initialize a fake cloud env only for example, normally you should do this in init() function
	os.Clearenv()
	err := os.Setenv("DYNO", "true")
	if err != nil {
		panic(err)
	}
	// Here we set a value for Orig field from MyConfig schema
	err = os.Setenv("CONFIG_ORIG", "<injected by gautocloud>")
	if err != nil {
		panic(err)
	}
	gautocloud.RegisterConnector(generic.NewConfigGenericConnector(MyConfig{}, argInterceptor))
	gautocloud.ReloadConnectors()
	//////

	var config MyConfig

	err = gautocloud.Inject(&config)
	if err != nil {
		panic(err)
	}
	// We can see that we have our config altered by flags found
	fmt.Printf("%#v\n", config)

	// Output: arg_test.MyConfig{Foo:"bar", Bar:true, Orig:"<injected by gautocloud>"}
}
