package main

import (
	"fmt"
	"reflect"
	"regexp"
)

type RegexpRule struct {
	Regexp string
	Value  interface{}
}

func (r RegexpRule) Passes() (bool, error) {
	v := reflect.ValueOf(r.Value)
	switch v.Kind() {
	case reflect.String:
		re := regexp.MustCompile(r.Regexp)
		return re.MatchString(v.String()), nil
	case reflect.Pointer:
	case reflect.Float32, reflect.Float64:
	case reflect.Complex64, reflect.Complex128:
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	case reflect.Invalid, reflect.Bool, reflect.Array, reflect.Chan, reflect.Func:
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	case reflect.Interface, reflect.Map, reflect.Struct, reflect.UnsafePointer, reflect.Uintptr, reflect.Slice:
	}
	return false, fmt.Errorf("regexp:%s not support", v.Kind())
}

func (r RegexpRule) Error() error {
	return fmt.Errorf("regexp invalid")
}

func NewRegexpRule(v interface{}, r string) *RegexpRule {
	return &RegexpRule{Value: v, Regexp: r}
}
