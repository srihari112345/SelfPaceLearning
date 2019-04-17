package main

import "testing"

// TestEquidistantForValidInputEndsWithOne validates IsEquistant for valid string and expects true result.
func TestEquidistantForValidInputEndsWithOne(t *testing.T) {
	result := IsEquistant("001001001")

	if result != true {
		t.Errorf("IsEquistant result was incorrect, got: %t, want: %t.", result, true)
	}
}

// TestEquidistantForInvalidInput validates IsEquistant for invalid string and expects false result.
func TestEquidistantForInvalidInput(t *testing.T) {
	result := IsEquistant("01010101010000010001010101010101010010")

	if result != false {
		t.Errorf("IsEquistant result was incorrect, got: %t, want: %t.", result, false)
	}
}

// TestEquidistantForValidInputEndsWithZero validates IsEquistant for valid string and expects true result.
func TestEquidistantForValidInputEndsWithZero(t *testing.T) {
	result := IsEquistant("01010101010")

	if result != true {
		t.Errorf("IsEquistant result was incorrect, got: %t, want: %t.", result, true)
	}
}
