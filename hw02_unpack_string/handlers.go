package hw02unpackstring

import "strconv"

type inputHandler struct {
	symbols         *[]*symbol
	lastSymbol      *symbol
	inputString     string
	currentPosition int
	currentValue    string
}

type handler interface {
	handle(input inputHandler) bool
}

type alphabeticHandler struct{}

func (h alphabeticHandler) handle(input inputHandler) bool {
	*input.symbols = append(*input.symbols, &symbol{
		value:  input.currentValue,
		repeat: 1,
	})
	return true
}

type escapedHandler struct{}

func (h *escapedHandler) handle(input inputHandler) bool {
	if input.currentPosition == 0 {
		return false
	}
	if input.currentValue != "\\" {
		return false
	}
	if input.lastSymbol.value == "\\" {
		input.lastSymbol.value = "\\\\"
		return true
	}
	return false
}

type numberHandler struct{}

func (h *numberHandler) handle(input inputHandler) bool {
	repeat, err := strconv.Atoi(input.currentValue)
	if err != nil {
		return false
	}
	if input.lastSymbol.value == "\\" {
		input.lastSymbol.value = input.currentValue
		return true
	}
	input.lastSymbol.repeat = repeat
	return true
}
