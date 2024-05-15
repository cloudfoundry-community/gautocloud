package urfave_test

import (
	"fmt"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/generic"
	. "github.com/cloudfoundry-community/gautocloud/interceptor/cli/urfave"
	"github.com/urfave/cli"
	"os"
)

func Example() {
	cliInterceptor := NewCli()
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
	gautocloud.RegisterConnector(generic.NewConfigGenericConnector(MyConfig{}, cliInterceptor))
	gautocloud.ReloadConnectors()
	//////

	// This is our urfave cli command action
	action := func(c *cli.Context) error {
		config := MyConfig{}
		// We pass context to interceptor, this is mandatory
		cliInterceptor.SetContext(c)

		// We are asking to retrieve a config schema through injection
		err := gautocloud.Inject(&config)
		if err != nil {
			panic(err)
		}
		// We can see that we have our config altered by flags found by urfave/cli
		fmt.Println(fmt.Sprintf("%#v", config))
		return nil
	}

	// Here we simply create a fake urfave/cli app
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "foo, f",
			Value: "Foo",
			Usage: "foo",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:        "doo",
			Aliases:     []string{"do"},
			Category:    "motion",
			Usage:       "do the doo",
			UsageText:   "doo - does the dooing",
			Description: "no really, there is a lot of dooing to be done",
			ArgsUsage:   "[arrgh]",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "bar, b"},
			},
			Action: action,
		},
	}
	app.Run([]string{"app", "--foo=bar", "doo", "--bar"})

	// Output: urfave_test.MyConfig{Foo:"bar", Bar:true, Orig:"<injected by gautocloud>"}
}
