package main

import (
	"fmt"
	"reflect"
	"strings"
)

type InRule struct {
	list  []string
	Value interface{}
}

func (r InRule) Passes() (bool, error) {
	v := reflect.ValueOf(r.Value)
	switch v.Kind() {
	case reflect.String:
		for _, el := range r.list {
			if el == v.String() {
				return true, nil
			}
		}
		return false, nil
	case reflect.Int:
		sv := fmt.Sprintf("%d", r.Value)
		for _, el := range r.list {
			if el == sv {
				return true, nil
			}
		}
		return false, nil
	case reflect.Pointer:
	case reflect.Float32, reflect.Float64:
	case reflect.Complex64, reflect.Complex128:
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	case reflect.Invalid, reflect.Bool, reflect.Array, reflect.Chan, reflect.Func:
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	case reflect.Interface, reflect.Map, reflect.Struct, reflect.UnsafePointer, reflect.Uintptr, reflect.Slice:
	}
	return false, fmt.Errorf("in:%s not support", v.Kind())
}

func (r InRule) Error() error {
	return fmt.Errorf("in: %s", strings.Join(r.list, ", "))
}

func NewInRule(v interface{}, l []string) *InRule {
	return &InRule{Value: v, list: l}
}
