package main

type Rule interface {
	Passes() (bool, error)
	Error() error
}
