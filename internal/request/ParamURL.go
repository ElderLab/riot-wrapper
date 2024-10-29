package request

import (
	"fmt"
	"reflect"
)

// ParamURL is a struct that contains the key and value of a URL parameter.
type ParamURL struct {
	Key   string
	Value string
}

// StructToListOfParamURL is a function that converts a struct to a list of ParamURL.
func StructToListOfParamURL(structure interface{}) []ParamURL {
	var result []ParamURL
	v := reflect.ValueOf(structure)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("query")
		value := v.Field(i).Interface()
		// check if the value is empty
		// if it is empty, continue to the next field
		if value == reflect.Zero(v.Field(i).Type()).Interface() {
			continue
		}
		result = append(result, ParamURL{Key: tag, Value: fmt.Sprint(value)})
	}
	return result
}
