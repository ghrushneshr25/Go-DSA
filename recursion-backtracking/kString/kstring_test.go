package kstring_test

import (
	kstring "godsa/recursion-backtracking/kString"
	"reflect"
	"sort"
	"testing"
)

// replace with the correct import path for your package

func TestKString(t *testing.T) {
	t.Run("Valid input with k=2 and size=2", func(t *testing.T) {
		kstr := kstring.KString{}
		kstr.Array = make([]string, 2)
		result := kstr.KStringRecursion(2, 2)
		expected := []string{"00", "01", "10", "11"}

		sort.Strings(expected)
		sort.Strings(result)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test failed. Expected %v, got %v", expected, result)
		}
	})

	t.Run("Valid input with k=3 and size=3", func(t *testing.T) {
		kstr := kstring.KString{}
		kstr.Array = make([]string, 3)
		result := kstr.KStringRecursion(3, 3)
		expected := []string{"000", "001", "002", "010", "011", "012", "020", "021", "022", "100", "101", "102", "110", "111", "112", "120", "121", "122", "200", "201", "202", "210", "211", "212", "220", "221", "222"}

		sort.Strings(expected)
		sort.Strings(result)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test failed. Expected %v, got %v", expected, result)
		}
	})

	t.Run("Valid input with k=1 and size=4", func(t *testing.T) {
		kstr := kstring.KString{}
		kstr.Array = make([]string, 4)
		result := kstr.KStringRecursion(4, 1)
		expected := []string{"0000"}

		sort.Strings(expected)
		sort.Strings(result)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test failed. Expected %v, got %v", expected, result)
		}
	})

	t.Run("Invalid input with size < 1", func(t *testing.T) {
		kstr := kstring.KString{}
		kstr.Array = make([]string, 0)
		result := kstr.KStringRecursion(0, 2)
		expected := []string{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test failed. Expected %v, got %v", expected, result)
		}
	})

	t.Run("Valid input with k=0 and size=0", func(t *testing.T) {
		kstr := kstring.KString{}
		kstr.Array = make([]string, 0)
		result := kstr.KStringRecursion(0, 0)
		expected := []string{}

		sort.Strings(expected)
		sort.Strings(result)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test failed. Expected %v, got %v", expected, result)
		}
	})

	t.Run("Valid input with k=1 and size=1", func(t *testing.T) {
		kstr := kstring.KString{}
		kstr.Array = make([]string, 1)
		result := kstr.KStringRecursion(1, 1)
		expected := []string{"0"}

		sort.Strings(expected)
		sort.Strings(result)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test failed. Expected %v, got %v", expected, result)
		}
	})

	t.Run("Valid input with k=0 and size=5", func(t *testing.T) {
		kstr := kstring.KString{}
		kstr.Array = make([]string, 5)
		result := kstr.KStringRecursion(5, 0)
		expected := []string{}

		sort.Strings(expected)
		sort.Strings(result)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test failed. Expected %v, got %v", expected, result)
		}
	})

	t.Run("Valid input with k=10 and size=1", func(t *testing.T) {
		kstr := kstring.KString{}
		kstr.Array = make([]string, 1)
		result := kstr.KStringRecursion(1, 10)
		expected := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

		sort.Strings(expected)
		sort.Strings(result)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test failed. Expected %v, got %v", expected, result)
		}
	})
}
