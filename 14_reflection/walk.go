package walk

import (
	"reflect"
)

// Goal - passing any type x and a callback in walk func
// We expect to call callback fn() for each x or x file of type string - with x as an arg
func walk(x any, fn func(string)) {

	val := getValue(x)

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}

	case reflect.Struct:
		// for i := 0; i < val.NumField(); i++ {
		// 	walk(val.Field(i).Interface(), fn)
		// }
		// cleaner new syntax
		for _, fieldVal := range val.Fields() {
			walk(fieldVal.Interface(), fn)
		}

	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}

	case reflect.Chan:
		for {
			valOfChan, ok := val.Recv()

			if ok {
				walk(valOfChan.Interface(), fn)
			} else {
				break
			}
		}

	case reflect.Func:
		funcResult := val.Call(nil)
		for _, v := range funcResult {
			walk(v.Interface(), fn)
		}

	case reflect.String:
		fn(val.String())
	}

}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
