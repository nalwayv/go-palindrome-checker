package palen

import (
	"testing"
)

func TestStrip(t *testing.T) {
	result := strip("_He_Ll*O")
	expected := "hello"

	if result != expected {
		t.Error("failed to strip string")
	}
}

func TestPalendrome(t *testing.T) {
	msg := "A man, a plan, a canal. Panama"

	result := IsPalendrome(msg)

	if result != true {
		t.Error("palendrome failed")
	}
}
