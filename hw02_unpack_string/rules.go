package hw02unpackstring

import (
	"errors"
	"strconv"
)

type inputRule struct {
	inputString     string
	currentPosition int
	currentValue    string
}

type rule interface {
	make(inputRule) error
}

func isInt(v string) bool {
	_, err := strconv.Atoi(v)
	return err == nil
}

type isValidFirstSymbolRule struct{}

func (r *isValidFirstSymbolRule) make(input inputRule) error {
	if input.currentPosition != 0 {
		return nil
	}
	if isInt(input.currentValue) {
		return errors.New("isNotNumberFirstSymbolRul")
	}
	return nil
}

type isValidNumberRepeaterRule struct{}

func (r *isValidNumberRepeaterRule) make(input inputRule) error {
	if !isInt(input.currentValue) || input.currentPosition < 2 || input.currentValue == "0" {
		return nil
	}
	if string(input.inputString[input.currentPosition-1]) == "\\" &&
		string(input.inputString[input.currentPosition-2]) == "\\" {
		return nil
	}
	if input.inputString[input.currentPosition-1] == input.inputString[input.currentPosition-2] {
		return errors.New("isValidNumberRepeaterRule")
	}
	return nil
}

type isValidEscapedSymbolRule struct{}

func (r *isValidEscapedSymbolRule) make(input inputRule) error {
	if input.currentValue != "\\" || len(input.inputString)-1 == input.currentPosition {
		return nil
	}
	nextValue := string(input.inputString[input.currentPosition+1])
	if nextValue == "\\" || isInt(nextValue) {
		return nil
	}
	return errors.New("isValidEscapedSymbolRule")
}
