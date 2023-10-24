package sortedarray_test

import (
	sortedarray "godsa/recursion-backtracking/sortedArray"
	"testing"
)

func TestIsArraySorted(t *testing.T) {
	t.Run("Empty Array", func(t *testing.T) {
		array := []int{}
		result := sortedarray.IsArraySorted(array, 0)
		expected := true
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Single Element Array", func(t *testing.T) {
		array := []int{42}
		result := sortedarray.IsArraySorted(array, 0)
		expected := true
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Sorted Array", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5}
		result := sortedarray.IsArraySorted(array, 0)
		expected := true
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Unsorted Array", func(t *testing.T) {
		array := []int{5, 4, 3, 2, 1}
		result := sortedarray.IsArraySorted(array, 0)
		expected := false
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Partially Sorted Array", func(t *testing.T) {
		array := []int{3, 2, 1, 4, 5}
		result := sortedarray.IsArraySorted(array, 0)
		expected := false
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
