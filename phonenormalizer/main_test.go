package main

import (
	"testing"
)

func TestCase(t *testing.T) {
	test := []struct {
		given    string
		expected string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"(123)456-7892", "1234567892"},
	}
	for _, val := range test {
		t.Run(val.given, func(t *testing.T) {
			if normalize(val.given) != val.expected {
				t.Errorf("%v given : %v expected", val.given, val.expected)
			}
		})
	}
}
