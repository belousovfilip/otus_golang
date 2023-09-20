package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var errInvalidString = errors.New("invalid string")

var list []symbol

type symbol struct {
	value  string
	repeat int
}

func unpack(input string) (output string, err error) {
	if len(input) == 0 {
		for _, symbol := range list {
			output += strings.Repeat(symbol.value, symbol.repeat)
		}
		list = []symbol{}
		return output, nil
	}
	if len(list) == 0 {
		if _, err := strconv.Atoi(string(input[0])); err == nil {
			return "", errInvalidString
		}
	}
	if string(input[0]) == "\\" && len(input) > 1 {
		_, err := strconv.Atoi(string(input[1]))
		if err == nil || string(input[1]) == "\\" {
			list = append(list, symbol{string(input[1]), 1})
			return unpack(input[2:])
		}
		list = []symbol{}
		return "", errInvalidString
	}
	if repeat, err := strconv.Atoi(string(input[0])); err == nil {
		if repeat != 0 &&
			len(list) > 1 &&
			list[len(list)-1].value == list[len(list)-2].value &&
			list[len(list)-2].value != "\\" {
			list = []symbol{}
			return "", errInvalidString
		}
		list[len(list)-1].repeat = repeat
		return unpack(input[1:])
	}
	list = append(list, symbol{string(input[0]), 1})
	return unpack(input[1:])
}
