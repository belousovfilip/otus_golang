package hw02unpackstring

import (
	"errors"
	"strings"
)

var errInvalidString = errors.New("invalid string")

type symbol struct {
	value  string
	repeat int
}

func unpack(input string) (string, error) {
	var output string
	var symbols []*symbol
	rules := []rule{
		&isValidFirstSymbolRule{},
		&isValidNumberRepeaterRule{},
		&isValidEscapedSymbolRule{},
	}
	handlers := []handler{
		&escapedHandler{},
		&numberHandler{},
		&alphabeticHandler{},
	}
	for pos, v := range input {
		currentValue := string(v)

		inputRule := inputRule{
			inputString:     input,
			currentPosition: pos,
			currentValue:    currentValue,
		}
		for _, rule := range rules {
			if err := rule.make(inputRule); err != nil {
				return "", errInvalidString
			}
		}

		inputHandler := inputHandler{
			symbols:         &symbols,
			inputString:     input,
			currentPosition: pos,
			currentValue:    currentValue,
		}
		for _, handler := range handlers {
			if len(symbols) > 0 {
				inputHandler.lastSymbol = symbols[len(symbols)-1]
			} else {
				inputHandler.lastSymbol = nil
			}
			if handler.handle(inputHandler) {
				break
			}
		}
	}
	for _, s := range symbols {
		if s.value == "\\\\" {
			s.value = "\\"
		}
		output += strings.Repeat(s.value, s.repeat)
	}
	return output, nil
}
