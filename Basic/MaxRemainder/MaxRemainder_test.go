package MaxRemainder

import "testing"

// TestCaclulateMaxRemainderForValidEvenInputExpectsThree
func TestCaclulateMaxRemainderForValidEvenInputExpectsThree(t *testing.T) {
	actual := CaclulateMaxRemainder([]int{4, 2, 1, 3})
	expected := 3

	if actual != expected {
		t.Errorf("CaclulateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

// TestCaclulateMaxRemainderForValidOddnputExpectsFour
func TestCaclulateMaxRemainderForValidOddInputExpectsFour(t *testing.T) {
	actual := CaclulateMaxRemainder([]int{2, 1, 5, 3, 4})
	expected := 2

	if actual != expected {
		t.Errorf("CaclulateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

// TestCaclulateMaxRemainderForValidSameElementInputExpectsZero
func TestCaclulateMaxRemainderForValidSameElementInputExpectsZero(t *testing.T) {
	actual := CaclulateMaxRemainder([]int{4, 4, 4, 4, 4})
	expected := 0

	if actual != expected {
		t.Errorf("CaclulateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}
