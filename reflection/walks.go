package reflection

import (
	"reflect"
)

func walks(x interface{}, fn func(arg string)) {
	value := reflect.ValueOf(x)
	field := value.Field(0)
	fn(field.String())
}
