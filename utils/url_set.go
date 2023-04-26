package utils

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/pkg/errors"
)

func BuildURLQueryParams(s interface{}) (string, error) {
	urlQuery := make(url.Values)

	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldValue := v.Field(i)
		zeroValue := reflect.Zero(fieldValue.Type()).Interface()

		// set if value is not zero
		if !reflect.DeepEqual(fieldValue.Interface(), zeroValue) {
			if inPathTagValue := field.Tag.Get("inpath"); inPathTagValue == "true" {
				urlQuery.Set(field.Tag.Get("json"), fmt.Sprintf("%v", fieldValue.Interface()))
			}
		} else if defaultTagValue := field.Tag.Get("default"); defaultTagValue != "" { // set if the value is zero, but we have default value
			urlQuery.Set(field.Tag.Get("json"), fmt.Sprintf("%v", defaultTagValue))
		} else if requireTagValue := field.Tag.Get("require"); requireTagValue == "true" { //
			return "", errors.New(fmt.Sprintf("required field %s is not set", field.Name))
		}
	}
	return urlQuery.Encode(), nil
}
