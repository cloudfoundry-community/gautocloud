// This is an interceptor dedicated to push flags found to the final schema given by gautocloud.
// if flags is not a zero value it will override value from schema given by gautocloud.
// It use https://github.com/alexflint/go-arg to translate flags into a struct.
package arg

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/cloudfoundry-community/gautocloud/interceptor"
	"io"
	"os"
	"reflect"
)

type ArgInterceptor struct {
	config arg.Config
	args   []string
	writer io.Writer
	exit   bool
}

// Option to set arg.Config from https://github.com/alexflint/go-arg
// Default: arg.Config{}
func Config(config arg.Config) optSetter {
	return func(f *ArgInterceptor) {
		f.config = config
	}
}

// Option to set args to be parsed as flags
// Default: os.Args
func Args(args []string) optSetter {
	return func(f *ArgInterceptor) {
		f.args = args
	}
}

// Option to set writer for output
// Default: os.Stdout
func Writer(w io.Writer) optSetter {
	return func(f *ArgInterceptor) {
		f.writer = w
	}
}

// Option to exit program or not when --help or --version flags has been found
// Default: true
func Exit(exit bool) optSetter {
	return func(f *ArgInterceptor) {
		f.exit = exit
	}
}

type optSetter func(i *ArgInterceptor)

func NewArg(setters ...optSetter) *ArgInterceptor {
	i := &ArgInterceptor{
		config: arg.Config{},
		args:   os.Args,
		writer: os.Stdout,
		exit:   true,
	}
	for _, s := range setters {
		s(i)
	}
	return i
}

func (i ArgInterceptor) Intercept(current, found interface{}) (interface{}, error) {
	schema := current
	if schema == nil {
		schema = found
	}
	sType := reflect.TypeOf(schema)
	if sType.Kind() == reflect.Ptr {
		return i.parse(schema, found)
	}
	final, err := i.parse(reflect.New(sType).Interface(), interceptor.InterfaceAsPtrCopy(found))
	if err != nil {
		return nil, err
	}
	return reflect.ValueOf(final).Elem().Interface(), nil
}

func (i ArgInterceptor) parse(schema, found interface{}) (interface{}, error) {
	p, err := arg.NewParser(i.config, schema)
	if err != nil {
		return nil, err
	}
	err = p.Parse(i.flags())
	if err == arg.ErrHelp {
		p.WriteHelp(i.writer)
		if !i.exit {
			return schema, nil
		}
		os.Exit(0)
	}
	version := "dev"
	if dest, ok := schema.(arg.Versioned); ok {
		version = dest.Version()
	}
	if err == arg.ErrVersion {
		fmt.Fprintln(i.writer, version)
		if !i.exit {
			return schema, nil
		}
		os.Exit(0)
	}
	if err != nil {
		return nil, err
	}
	return interceptor.NewOverwrite().Intercept(schema, found)
}
func (i ArgInterceptor) flags() []string {
	if len(i.args) == 0 { // os.Args could be empty
		return nil
	}
	return i.args[1:]
}
