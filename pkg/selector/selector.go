package selector

import (
	"reflect"
)

// This package is for selecting(or hiding) fields in struct dynamicly

func fieldSet(fields ...string) map[string]bool {
	set := make(map[string]bool, len(fields))
	for _, s := range fields {
		set[s] = true
	}
	return set
}

func SelectFields(sturct interface{}, fields ...string) interface{} {

	newStruct := reflect.New(reflect.TypeOf(sturct).Elem())
	fs := fieldSet(fields...)
	rt, rv := reflect.TypeOf(sturct).Elem(), reflect.ValueOf(sturct).Elem()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		jsonKey := field.Tag.Get("selector")
		
		if fs[jsonKey] {
			newStruct.Elem().FieldByName(field.Name).Set(rv.FieldByName(field.Name))
		}
	}
	return newStruct.Elem().Interface()
}