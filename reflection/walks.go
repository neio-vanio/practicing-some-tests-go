package reflection

import (
	"reflect"
)

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}

func walks(x interface{}, fn func(arg string)) {

	value := getValue(x)

	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			walks(value.Index(i).Interface(), fn)
		}
		return
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		switch field.Kind() {
		case reflect.Struct:
			walks(field.Interface(), fn)
		case reflect.String:
			fn(field.String())
		}
	}
}

// Person type user
type Person struct {
	Name    string
	Profile Profile
}

// Profile type of Person
type Profile struct {
	Age  int
	City string
}
