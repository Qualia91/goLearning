package main

import (
	"errors"
	"strings"
)

//Model service as an interface
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

var ErrEmpty = errors.New("Empty String")

type stringService struct {
}

// Implements StringService
func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

// Implements StringService
func (stringService) Count(s string) int {
	return len(s)
}

// Primary message patter is remote procedure call (RPC) so every method in our interface will be modelled as such.
// Therefore, for each method, we define a request and response.
type uppercaseRequest struct {
	S string `json:"s"`
}
