package cracklib

import (
	"testing"
)

var testFascistCheck = []struct {
	pw       string
	expected bool
}{
	{"foo", false},
	{"", false},
	{"123", false},
	{"!!%6AsdfGhjLlojj", true},
}

func TestFascistCheck(t *testing.T) {
	for _, tt := range testFascistCheck {
		_, ok := FascistCheck(tt.pw)
		if ok != tt.expected {
			t.Errorf("FascistCheck(%s): expected %t, actual %t", tt.pw, tt.expected, ok)
		}
	}
}

func TestFascistCheckUser(t *testing.T) {
	for _, tt := range testFascistCheck {
		_, ok := FascistCheckUser(tt.pw, "foobar")
		if ok != tt.expected {
			t.Errorf("FascistCheck(%s): expected %t, actual %t", tt.pw, tt.expected, ok)
		}
	}
}
