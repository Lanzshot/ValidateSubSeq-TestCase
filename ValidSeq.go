/*
2 sets of array where need to determine if the 2nd array is a subset of the 1st array.
*/

package main

import (
	"fmt"
)

func main() {

	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{25, 8}

	var result bool = IsValidSubseq(array, sequence)

	fmt.Printf("The subseq is :%v\n", result)

}

func IsValidSubseq(array []int, sequence []int) bool {

	indSeq := 0
	for _, eachArray := range array {
		if indSeq == len(sequence) {
			break
		}
		if sequence[indSeq] == eachArray {
			indSeq += 1
		}
	}

	return indSeq == len(sequence)
}
