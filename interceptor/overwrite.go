package interceptor

import (
	"reflect"
)

// This interceptor function let the user pre filled values inside is config structure to be used instead of use values
// found by gautocloud.
// This is useful for schema used in connector generic Config to let user write some values from config
// before fill the rest by gautocloud.
// This can only be used when user use inject functions from gautocloud.
func NewOverwrite() Intercepter {
	return IntercepterFunc(overwrite)
}

func overwrite(current, found interface{}) (interface{}, error) {
	if current == nil {
		return found, nil
	}
	cVal := reflect.ValueOf(current)
	fVal := reflect.ValueOf(found)
	toReturn := reflect.New(reflect.TypeOf(found)).Elem()
	for index := 0; index < cVal.NumField(); index++ {
		vField := cVal.Field(index)
		if isZero(vField.Interface()) {
			toReturn.Field(index).Set(fVal.Field(index))
		} else {
			toReturn.Field(index).Set(vField)
		}
	}
	return toReturn.Interface(), nil
}

func isZero(v interface{}) bool {
	t := reflect.TypeOf(v)
	zeroInt := reflect.Zero(t).Interface()
	if !t.Comparable() {
		return reflect.DeepEqual(v, zeroInt)
	}
	return v == zeroInt
}
