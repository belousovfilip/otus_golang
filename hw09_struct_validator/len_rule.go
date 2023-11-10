package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type LenRule struct {
	Length string
	Value  interface{}
}

func (r LenRule) Passes() (bool, error) {
	length, err := strconv.Atoi(r.Length)
	if err != nil {
		return false, fmt.Errorf("invalid length value")
	}
	v := reflect.ValueOf(r.Value)
	switch v.Kind() {
	case reflect.String:
		return len([]rune(v.String())) == length, nil
	case reflect.Slice:
		return v.Len() == length, nil
	case reflect.Pointer:
	case reflect.Float32, reflect.Float64:
	case reflect.Complex64, reflect.Complex128:
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	case reflect.Invalid, reflect.Bool, reflect.Array, reflect.Chan, reflect.Func:
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	case reflect.Interface, reflect.Map, reflect.Struct, reflect.UnsafePointer, reflect.Uintptr:
	}
	return false, fmt.Errorf("len:%s not support", v.Kind())
}

func (r LenRule) Error() error {
	return fmt.Errorf("length: %s", r.Length)
}

func NewLenRule(v interface{}, l string) *LenRule {
	return &LenRule{Value: v, Length: l}
}
