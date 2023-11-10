package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type MaxRule struct {
	Max   string
	Value interface{}
}

func (r MaxRule) Passes() (bool, error) {
	length, err := strconv.Atoi(r.Max)
	if err != nil {
		return false, fmt.Errorf("invalid max value")
	}
	v := reflect.ValueOf(r.Value)
	switch v.Kind() {
	case reflect.Slice:
		return v.Len() <= length, nil
	case reflect.Int:
		return v.Int() <= int64(length), nil
	case reflect.String:
		return len([]rune(v.String())) <= length, nil
	case reflect.Pointer:
	case reflect.Float32, reflect.Float64:
	case reflect.Complex64, reflect.Complex128:
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	case reflect.Invalid, reflect.Bool, reflect.Array, reflect.Chan, reflect.Func:
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	case reflect.Interface, reflect.Map, reflect.Struct, reflect.UnsafePointer, reflect.Uintptr:
	}
	return false, fmt.Errorf("max:%s not support", v.Kind())
}

func (r MaxRule) Error() error {
	return fmt.Errorf("max %s", r.Max)
}

func NewMaxRule(v interface{}, m string) *MaxRule {
	return &MaxRule{Value: v, Max: m}
}
