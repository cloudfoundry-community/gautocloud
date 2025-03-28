package urfave_test

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors/generic"
	. "github.com/cloudfoundry-community/gautocloud/interceptor/cli/urfave"
	"github.com/urfave/cli"
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
	err := os.Setenv("DYNO", "true")
	if err != nil {
		panic(err)
	}
	// Here we set a value for Orig field from MyConfig schema
	err = os.Setenv("CONFIG_ORIG", "<injected by gautocloud>")
	if err != nil {
		panic(err)
	}
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
		fmt.Printf("%#v\n", config)
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
	err = app.Run([]string{"app", "--foo=bar", "doo", "--bar"})
	if err != nil {
		panic(err)
	}

	// Output: urfave_test.MyConfig{Foo:"bar", Bar:true, Orig:"<injected by gautocloud>"}
}
