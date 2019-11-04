package common

import (
	"html"
	"reflect"
)

func FilterXSS(in interface{}) interface{} {
	value := reflect.ValueOf(in).Elem()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if field.Type() != reflect.TypeOf("") {
			continue
		}
		str := field.Interface().(string)
		field.SetString(html.EscapeString(str))
	}
	return value
}
