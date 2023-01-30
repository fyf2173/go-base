package utils

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func ViperGetNode(key string, node interface{}) error {
	return viper.UnmarshalKey(key, node, func(m *mapstructure.DecoderConfig) {
		m.TagName = "yaml"
	})
}

// Validate 校验参数
func Validate(params interface{}) error {
	reflectVal := reflect.ValueOf(params)
	if reflectVal.Type().Kind() == reflect.Ptr {
		reflectVal = reflectVal.Elem()
	}

	if err := validator.New().Struct(params); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			structField, _ := reflectVal.Type().FieldByName(v.Field())
			msg := structField.Tag.Get(v.Tag())
			return fmt.Errorf(msg)
		}
	}
	return nil
}
