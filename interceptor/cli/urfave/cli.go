// Package urfave This is an interceptor dedicated to push flags from https://github.com/urfave/cli to the final schema given by gautocloud.
// if flags is not a zero value it will override value from schema given by gautocloud.
package urfave

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudfoundry-community/gautocloud/decoder"
	"github.com/cloudfoundry-community/gautocloud/interceptor"
	"github.com/urfave/cli"
)

type CliInterceptor struct {
	context *cli.Context
}

func NewCli() *CliInterceptor {
	return &CliInterceptor{}
}
func (i CliInterceptor) Intercept(current, found interface{}) (interface{}, error) {
	if i.context == nil {
		return nil, fmt.Errorf("context must be passed to CliInterceptor, please set it with SetContext")
	}
	schema := current
	if schema == nil {
		schema = found
	}
	mapFlags := make(map[string]interface{})
	flagNames := i.context.FlagNames()
	flagNames = append(flagNames, i.context.GlobalFlagNames()...)
	for _, name := range flagNames {
		key := strings.ReplaceAll(name, "-", "_")
		value := i.context.Generic(name)
		if value == nil {
			value = i.context.GlobalGeneric(name)
		}
		mapFlags[key] = fmt.Sprint(value)
	}
	schemaPtr := interceptor.InterfaceAsPtr(schema)

	err := decoder.UnmarshalNoDefault(mapFlags, schemaPtr)
	if err != nil {
		return nil, err
	}
	schemaKind := reflect.TypeOf(schema).Kind()
	if schemaKind == reflect.Ptr { //nolint:govet
		return interceptor.NewOverwrite().Intercept(schemaPtr, found)
	}
	return interceptor.NewOverwrite().Intercept(
		reflect.ValueOf(schemaPtr).Elem().Interface(),
		found,
	)
}

func (i *CliInterceptor) SetContext(context *cli.Context) {
	i.context = context
}
