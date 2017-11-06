package utils

import "reflect"

func ConvertMapInterface(toConvert interface{}) interface{} {
	typeData := reflect.TypeOf(toConvert)
	if typeData != reflect.TypeOf(make(map[interface{}]interface{})) && typeData != reflect.TypeOf(make([]interface{}, 0)) {
		return reflect.ValueOf(toConvert).Interface()
	}
	if typeData == reflect.TypeOf(make([]interface{}, 0)) {
		dataSlice := toConvert.([]interface{})
		for i, data := range dataSlice {
			dataSlice[i] = ConvertMapInterface(data)
		}
		return dataSlice
	}
	converted := make(map[string]interface{})
	for key, value := range toConvert.(map[interface{}]interface{}) {
		converted[key.(string)] = ConvertMapInterface(value)
	}

	return converted
}
