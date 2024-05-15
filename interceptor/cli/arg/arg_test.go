package arg_test

import (
	. "github.com/cloudfoundry-community/gautocloud/interceptor/cli/arg"

	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

type SchemaFake struct {
	Foo  string
	Foo2 string
}
type SchemaFakeWithVersion struct {
	Foo  string
	Foo2 string
}

func (SchemaFakeWithVersion) Version() string {
	return "1.0.0"
}

var _ = Describe("Arg", func() {
	var buf *bytes.Buffer
	BeforeEach(func() {
		buf = new(bytes.Buffer)

	})
	Context("Schema is not a pointer", func() {
		It("should create a new schema merging what found as flag and what found by gautocloud", func() {
			i := NewArg(Writer(buf), Args([]string{
				"app",
				`--foo=bar`,
			}))

			final, err := i.Intercept(SchemaFake{}, SchemaFake{
				Foo:  "orig",
				Foo2: "bar",
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(final.(SchemaFake).Foo).Should(Equal("bar"))
			Expect(final.(SchemaFake).Foo2).Should(Equal("bar"))
		})
	})
	Context("Schema is a pointer", func() {
		It("should create a new schema merging what found as flag and what found by gautocloud", func() {
			i := NewArg(Writer(buf), Args([]string{
				"app",
				`--foo=bar`,
			}))

			final, err := i.Intercept(&SchemaFake{}, &SchemaFake{
				Foo:  "orig",
				Foo2: "bar",
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(final.(*SchemaFake).Foo).Should(Equal("bar"))
			Expect(final.(*SchemaFake).Foo2).Should(Equal("bar"))
		})
	})
	Context("Schema is not given by user", func() {
		It("should create a new schema merging what found as flag and what found by gautocloud", func() {
			i := NewArg(Writer(buf), Args([]string{
				"app",
				`--foo=bar`,
			}))

			final, err := i.Intercept(nil, SchemaFake{
				Foo:  "orig",
				Foo2: "bar",
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(final.(SchemaFake).Foo).Should(Equal("bar"))
			Expect(final.(SchemaFake).Foo2).Should(Equal("bar"))
		})
	})
	It("should return help if flag --help is given", func() {
		i := NewArg(Writer(buf), Exit(false), Args([]string{
			"app",
			`--help`,
		}))

		_, err := i.Intercept(SchemaFake{}, SchemaFake{
			Foo:  "orig",
			Foo2: "bar",
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(buf.String()).Should(ContainSubstring("--foo2"))
		Expect(buf.String()).Should(ContainSubstring("--foo"))
	})
	It("should fail when flag --version is given and schema does not provide a Version", func() {
		i := NewArg(Writer(buf), Exit(false), Args([]string{
			"app",
			`--version`,
		}))

		_, err := i.Intercept(SchemaFake{}, SchemaFake{
			Foo:  "orig",
			Foo2: "bar",
		})
		Expect(err).To(HaveOccurred())
	})
	It("should return version given by schema if flag --version is given and schema implement Versionned", func() {
		i := NewArg(Writer(buf), Exit(false), Args([]string{
			"app",
			`--version`,
		}))

		_, err := i.Intercept(SchemaFakeWithVersion{}, SchemaFakeWithVersion{
			Foo:  "orig",
			Foo2: "bar",
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(strings.TrimSpace(buf.String())).Should(Equal("1.0.0"))
	})
})
