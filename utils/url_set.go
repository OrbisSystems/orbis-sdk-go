package utils

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/pkg/errors"
)

const (
	trueStr = "true"

	inPathTagName  = "inpath"
	defaultTagName = "default"
	jsonTagName    = "json"
	requireTagName = "require"
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
			if inPathTagValue := field.Tag.Get(inPathTagName); inPathTagValue == trueStr {
				urlQuery.Set(field.Tag.Get(jsonTagName), fmt.Sprintf("%v", fieldValue.Interface()))
			}
			// set if the value is zero, but we have default value
		} else if defaultTagValue := field.Tag.Get(defaultTagName); defaultTagValue != "" {
			urlQuery.Set(field.Tag.Get(jsonTagName), defaultTagValue)
			// also check is there and required fields that not set
		} else if requireTagValue := field.Tag.Get(requireTagName); requireTagValue == trueStr {
			return "", errors.New(fmt.Sprintf("required field %s is not set", field.Name))
		}
	}
	return urlQuery.Encode(), nil
}

func ArrayToJoinedString(value []string) string {
	res := ""
	if len(value) == 0 {
		return res
	}

	for idx := range value {
		if idx == len(value)-1 {
			res += value[idx]
		} else {
			res += value[idx] + ","
		}
	}

	return res
}
