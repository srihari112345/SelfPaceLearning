/*
URL : https://www.geeksforgeeks.org/check-if-all-the-1s-in-a-binary-string-are-equidistant-or-not/

Given a binary string str, the task is to check if all the 1’s in the string are equidistant or not. The term equidistant means that the distance between every two adjacent 1’s is same. Note that the string contains at least two 1’s.

Examples:

Input: str = “00111000”
Output: Yes
The distance between all the 1’s is same and is equal to 1.

Input: str = “0101001”
Output: No
The distance between the 1st and the 2nd 1’s is 2
and the distance between the 2nd and the 3rd 1’s is 3.
*/

package main

import "fmt"

func main() {

	input := "001001001"
	//input := "01010101010000010001010101010101010010"
	//input := "01010101010"

	fmt.Println(IsEquistant(input))
}

// IsEquistant Returns bool value based on input string
func IsEquistant(input string) bool {
	index := 0
	prevIndex := 0
	prevDistance := 0
	currentDistance := 0
	isEquidistant := true

	//Could be refactored to use range
	for {
		if string(input[index]) == "1" {
			if prevIndex == 0 {
				prevIndex = index
			} else if prevDistance == 0 {
				prevDistance = index - prevIndex
				prevIndex = index
			} else if prevIndex != 0 && prevDistance != 0 {
				currentDistance = index - prevIndex
				prevIndex = index
			}
		}

		index++

		if currentDistance > prevDistance {
			isEquidistant = false
			break
		}

		if index >= len(input) {
			break
		}
	}

	return isEquidistant
}
