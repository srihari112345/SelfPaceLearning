package main

import "testing"

// TestEquidistanceForValidInputEndsWithOne validates IsEquistant for valid string and expects true result.
func TestEquidistanceForValidInputEndsWithOne(t *testing.T) {
	result := IsEquistant("001001001")

	if result != true {
		t.Errorf("IsEquistant result was incorrect, got: %t, want: %t.", result, true)
	}
}

// TestEquidistanceForInValidInput validates IsEquistant for invalid string and expects false result.
func TestEquidistanceForInValidInput(t *testing.T) {
	result := IsEquistant("01010101010000010001010101010101010010")

	if result != false {
		t.Errorf("IsEquistant result was incorrect, got: %t, want: %t.", result, false)
	}
}

// TestEquidistanceForValidInputEndsWithZero validates IsEquistant for valid string and expects true result.
func TestEquidistanceForValidInputEndsWithZero(t *testing.T) {
	result := IsEquistant("01010101010")

	if result != true {
		t.Errorf("IsEquistant result was incorrect, got: %t, want: %t.", result, true)
	}
}
