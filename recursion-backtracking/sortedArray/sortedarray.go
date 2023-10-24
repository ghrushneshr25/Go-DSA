/*
	PROBLEM 2:
	Given an array, check whether the array is in sorted order with recursion.
*/

package sortedarray

func IsArraySorted(array []int, index int) bool {
	if len(array) == 0 || index == len(array)-1 {
		return true
	}
	if array[index] > array[index+1] {
		return false
	}
	return IsArraySorted(array, index+1)
}
