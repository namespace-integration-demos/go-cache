package main

import (
	"errors"
	"strings"
	"testing"
)

func TestGrpcError(t *testing.T) {
	inputs := map[error]error{
		errors.New("hello"): errors.New("hello"),
		errors.New("world"): errors.New("world"),
	}

	for s, expected := range inputs {
		if c := WrapError(s); !strings.Contains(c.Error(), expected.Error()) {
			t.Errorf("Expected %v error is %v.", c, expected)
		}
	}
}
