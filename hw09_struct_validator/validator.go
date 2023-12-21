package main

import (
	"fmt"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	r := ""
	for _, e := range v {
		r += e.Field + ": " + e.Err.Error() + "\n"
	}
	return r
}

func Validate(v interface{}) error {
	if reflect.ValueOf(v).Kind() != reflect.Struct {
		return ValidationErrors{}
	}
	rules, err := parseRules(v)
	if err != nil {
		return err
	}
	errors := ValidationErrors{}
	for filed, rule := range rules {
		passes, err := rule.Passes()
		if err != nil {
			return err
		}
		if !passes {
			e := ValidationError{Field: filed, Err: (rule).Error()}
			errors = append(errors, e)
		}
	}
	return errors
}

func parseRules(v interface{}) (map[string]Rule, error) {
	rules := map[string]Rule{}
	refValue := reflect.ValueOf(v)
	for i := 0; i < refValue.NumField(); i++ {
		field := refValue.Type().Field(i)
		tag := field.Tag.Get("validate")
		if tag == "" {
			continue
		}
		for _, r := range strings.Split(tag, "|") {
			rawRule := strings.Split(r, ":")
			name := rawRule[0]
			args := rawRule[1:][0]
			value := refValue.FieldByName(field.Name).Interface()
			switch name {
			default:
				return nil, fmt.Errorf("%s rule not found", name)
			case "len":
				rules[field.Name] = NewLenRule(value, args)
			case "max":
				rules[field.Name] = NewMaxRule(value, args)
			case "min":
				rules[field.Name] = NewMinRule(value, args)
			case "regexp":
				rules[field.Name] = NewRegexpRule(value, args)
			case "in":
				rules[field.Name] = NewInRule(value, strings.Split(args, ","))
			}
		}
	}
	return rules, nil
}
