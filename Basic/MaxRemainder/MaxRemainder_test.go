package MaxRemainder

import "testing"

// TestCaclulateMaxRemainderForValidInputExpectsOne
func TestCaclulateMaxRemainderForValidInputExpectsOne(t *testing.T) {
	actual := CaclulateMaxRemainder([]int{1, 2, 3, 4})
	expected := 1

	if actual != expected {
		t.Errorf("AreEquistant result was incorrect, got: %d, want: %d.", actual, expected)
	}
}
