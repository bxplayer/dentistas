package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func StructToKeyValueString(obj interface{}) (string, error) {
	var result []string
	val := reflect.ValueOf(obj)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		key := field.Tag.Get("json")
		value := val.Field(i).Interface()

		if value != "" {
			result = append(result, fmt.Sprintf("%s = '%v'", key, value))
		}
	}

	if len(result) == 0 {
		return "", errors.New("No hay elementos para actualizar")
	}

	return strings.Join(result, ", "), nil
}
