/*
URL: https://www.geeksforgeeks.org/find-frequency-of-smallest-value-in-an-array/

Find frequency of smallest value in an array
Given an array A of N elements. Find the frequency of the smallest value in the array.

Examples:

Input : N = 5, arr[] = {3, 2, 3, 4, 4}
Output : 1
The smallest element in the array is 2
and it occurs only once.

Input : N = 6, arr[] = {4, 3, 5, 3, 3, 6}
Output : 3
The smallest element in the array is 3
and it occurs 3 times.

Approaches:

1. Traverse through input slice/array and build dictionary with Key as element
value and Value as frequency of the element.Then traverse dictionary to find the minimum key and return the value.

2. Sort the input slice/array and traverse untill different element found and return the frequency.

3. Traverse through input slice/array maintain mimimum element & frequency.

4. Optimize traversal by N/2 of the input slice/array where N is size of the input and maintain minimum element & frequency.

/*/

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
