package arg_test

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/generic"
	. "github.com/cloudfoundry-community/gautocloud/interceptor/cli/arg"
	"os"
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
	os.Setenv("DYNO", "true")
	// Here we set a value for Orig field from MyConfig schema
	os.Setenv("CONFIG_ORIG", "<injected by gautocloud>")
	gautocloud.RegisterConnector(generic.NewConfigGenericConnector(MyConfig{}, argInterceptor))
	gautocloud.ReloadConnectors()
	//////

	var config MyConfig

	err := gautocloud.Inject(&config)
	if err != nil {
		panic(err)
	}
	// We can see that we have our config altered by flags found
	fmt.Println(fmt.Sprintf("%#v", config))

	// Output: arg_test.MyConfig{Foo:"bar", Bar:true, Orig:"<injected by gautocloud>"}
}
