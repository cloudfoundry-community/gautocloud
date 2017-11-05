package urfave_test

import (
	. "github.com/cloudfoundry-community/gautocloud/interceptor/cli/urfave"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/urfave/cli"
)

type SchemaFake struct {
	Foo    string
	FooBar bool
	Found  string
}

var _ = Describe("Urfave", func() {
	var app *cli.App
	BeforeEach(func() {
		app = cli.NewApp()
	})
	Context("schema is not a pointer", func() {
		It("should put flags into schema provided by user", func() {
			ran := false
			i := NewCli()
			action := func(c *cli.Context) error {
				i.SetContext(c)
				ran = true
				final, err := i.Intercept(SchemaFake{}, SchemaFake{
					Found: "found",
				})
				Expect(err).ToNot(HaveOccurred())
				schema := final.(SchemaFake)
				Expect(schema.Foo).Should(Equal("bar"))
				Expect(schema.FooBar).Should(BeTrue())
				Expect(schema.Found).Should(Equal("found"))
				return nil
			}

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
						cli.BoolFlag{Name: "foo-bar, fb"},
					},
					Action: action,
				},
			}

			app.Run([]string{"app", "--foo=bar", "doo", "--foo-bar"})

			Expect(ran).To(BeTrue())
		})
	})
	Context("schema is a pointer", func() {
		It("should put flags into schema provided by user", func() {
			ran := false
			i := NewCli()
			action := func(c *cli.Context) error {
				i.SetContext(c)
				ran = true
				final, err := i.Intercept(&SchemaFake{}, &SchemaFake{
					Found: "found",
				})
				Expect(err).ToNot(HaveOccurred())
				schema := final.(*SchemaFake)
				Expect(schema.Foo).Should(Equal("bar"))
				Expect(schema.FooBar).Should(BeTrue())
				Expect(schema.Found).Should(Equal("found"))
				return nil
			}

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
						cli.BoolFlag{Name: "foo-bar, fb"},
					},
					Action: action,
				},
			}

			app.Run([]string{"app", "--foo=bar", "doo", "--foo-bar"})

			Expect(ran).To(BeTrue())
		})
	})
})
