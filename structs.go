package Astruct

import (
	"fmt"
	"reflect"
)

// pbUIConfig = &lvideo_pb.UIConfig{
// ChannelBgColor:              modelUIConfig.ChannelBgColor,
// BlockBgColor:                modelUIConfig.BlockBgColor,
// BlockTitleColor:             modelUIConfig.BlockTitleColor,
// BlockToprightColor:          modelUIConfig.BlockToprightColor,
// BlockSeplineColor:           modelUIConfig.BlockSeplineColor,
// ...
// }

//if two struct have same field and the field type is same, auto assign the field value in output
func AssignSameFieldStruct(input interface{}, output interface{}) error {

	if !IsStruct(input) || !IsStruct(output) {
		return fmt.Errorf("input and output must be struct or pointer to struct")
	}

	if !IsPtr(output) {
		return fmt.Errorf("output must be a pointer to struct")
	}

	inputVal := GetInterfaceValue(input)
	outputVal := GetInterfaceValue(output)

	inputFields := getStructFields(inputVal)

	for _, field := range inputFields {
		inputFieldVal := inputVal.FieldByName(field.Name)
		zero := reflect.Zero(inputFieldVal.Type()).Interface()
		if reflect.DeepEqual(inputFieldVal.Interface(), zero) { // zero field value
			continue
		}
		outputFieldVal := outputVal.FieldByName(field.Name)

		if !outputFieldVal.IsValid() {
			continue
		}

		if inputFieldVal.Type() != outputFieldVal.Type() {
			continue
		}
		if !outputFieldVal.CanSet() {
			continue
		}

		outputVal.FieldByName(field.Name).Set(inputFieldVal)
	}
	return nil
}

//  struct or pointer to struct
func IsStruct(s interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Invalid {
		return false
	}

	return v.Kind() == reflect.Struct
}

func IsPtr(s interface{}) bool {
	v := reflect.ValueOf(s)
	return v.Kind() == reflect.Ptr
}

func GetInterfaceValue(s interface{}) reflect.Value {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func getStructFields(val reflect.Value) []reflect.StructField {
	t := val.Type()
	var f []reflect.StructField
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// unexported field
		if field.PkgPath != "" {
			continue
		}
		f = append(f, field)
	}
	return f
}
