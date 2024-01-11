package misc

import "reflect"

func AssignFields[T any](src, dst *T) {
	srcVal := reflect.ValueOf(src).Elem()
	if srcVal.Kind() != reflect.Struct {
		panic("arguments must be pointers to structs")
	}

	dstVal := reflect.ValueOf(dst).Elem()
	srcType := srcVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcType.Field(i)
		dstField := dstVal.FieldByName(srcField.Name)
		if dstField.IsValid() {
			if dstField.CanSet() {
				dstField.Set(srcVal.Field(i))
			}
		}
	}
}
