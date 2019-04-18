package minelementfrequency

// Note the use of << to create an untyped constant.
const bitsPerWord = 32 << uint(^uint(0)>>63)

// BitsPerWord Implementation-specific size of int and uint in bits.
const BitsPerWord = bitsPerWord // either 32 or 64

// Implementation-specific integer limit values.
const (
	MaxInt  = 1<<(BitsPerWord-1) - 1 // either 1<<31 - 1 or 1<<63 - 1
	MinInt  = -MaxInt - 1            // either -1 << 31 or -1 << 63
	MaxUint = 1<<BitsPerWord - 1     // either 1<<32 - 1 or 1<<64 - 1
)

// CalculateMinimumElementFrequency will return frequency of minimum element from input slice.
func CalculateMinimumElementFrequency(input []int) int {
	minElement := MaxInt
	frequency := 0
	for _, element := range input {
		if element < minElement {
			frequency = 1
			minElement = element
		} else if element == minElement {
			frequency++
		}
	}

	return frequency
}

func main() {
	input := []int{1, 2, 3, 4}
	CalculateMinimumElementFrequency(input)
}
