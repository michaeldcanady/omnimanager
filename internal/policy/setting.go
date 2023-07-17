package policy

import (
	"fmt"
	"reflect"
)

type Setting interface{}

func NewSetting(DisplayName string, Uri string, Value any) (Setting, error) {
	valueType := reflect.TypeOf(Value)
	switch valueType.Kind() {
	case reflect.String:
		return newSettingString(DisplayName, Uri, Value.(string)), nil
	default:
		return nil, fmt.Errorf("unable supported setting type: %t", Value)
	}
}
