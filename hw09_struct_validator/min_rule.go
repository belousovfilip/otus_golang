package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type MinRule struct {
	Min   string
	Value interface{}
}

func (r MinRule) Passes() (bool, error) {
	length, err := strconv.Atoi(r.Min)
	if err != nil {
		return false, fmt.Errorf("invalid min value")
	}
	v := reflect.ValueOf(r.Value)
	switch v.Kind() {
	case reflect.String:
		return len([]rune(v.String())) >= length, nil
	case reflect.Slice:
		return v.Len() >= length, nil
	case reflect.Int:
		return v.Int() >= int64(length), nil
	case reflect.Pointer:
	case reflect.Float32, reflect.Float64:
	case reflect.Complex64, reflect.Complex128:
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	case reflect.Invalid, reflect.Bool, reflect.Array, reflect.Chan, reflect.Func:
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	case reflect.Interface, reflect.Map, reflect.Struct, reflect.UnsafePointer, reflect.Uintptr:
	}
	return false, fmt.Errorf("min:%s not support", v.Kind())
}

func (r MinRule) Error() error {
	return fmt.Errorf("min %s", r.Min)
}

func NewMinRule(v interface{}, m string) *MinRule {
	return &MinRule{Value: v, Min: m}
}
