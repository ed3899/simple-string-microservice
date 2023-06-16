package simplestringservice

import (
	"errors"
	"strings"
)

type StringServiceI interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type StringService struct {
}

func (StringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (StringService) Count(s string) int {
	return len(s)
}

var ErrEmpty = errors.New("empty string")