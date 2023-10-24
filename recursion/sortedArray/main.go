/*
PROBLEM 2:
Given an array, check whether the array is in sorted order with recursion.
*/

package main

import "fmt"

func isArraySorted(array []int, index int) bool {
	if index == len(array)-1 {
		return true
	}
	if array[index] > array[index+1] {
		return false
	}
	return isArraySorted(array, index+1)
}

func main() {
	sortedArray := []int{1, 2, 3, 4, 5}
	fmt.Println(isArraySorted(sortedArray, 0))

	notSortedArray := []int{1, 6, 3, 4, 5}
	fmt.Println(isArraySorted(notSortedArray, 0))
}
