package interceptor

import (
	"reflect"
)

func InterfaceAsPtr(i interface{}) interface{} {
	iType := reflect.TypeOf(i)
	if iType.Kind() == reflect.Pointer {
		return i
	}
	return reflect.New(iType).Interface()
}

func InterfaceAsPtrCopy(i interface{}) interface{} {
	iType := reflect.TypeOf(i)
	if iType.Kind() == reflect.Pointer {
		return i
	}
	iPtr := reflect.New(iType)
	iPtr.Elem().Set(reflect.ValueOf(i))
	return iPtr.Interface()
}
