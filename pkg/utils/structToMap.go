package utils

import (
	"reflect"
	"strings"
)

func StructToMap(obj interface{}) map[string]interface{} {

	result := make(map[string]interface{})
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()

	}
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue

		}
		key := jsonTag
		idx := strings.Index(key, ",")
		if idx != -1 {
			key = key[:idx]

		}
		if !v.Field(i).IsZero() {
			result[key] = value

		}
	}
	return result

}
