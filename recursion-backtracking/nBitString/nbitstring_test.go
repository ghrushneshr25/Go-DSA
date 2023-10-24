package nbitstring_test

import (
	nbitstring "godsa/recursion-backtracking/nBitString"
	"reflect"
	"sort"
	"testing"
)

func TestNBitString_Binary(t *testing.T) {
	t.Run("Binary with size 1", func(t *testing.T) {
		nbit := nbitstring.NBitString{}
		nbit.Array = make([]string, 1)
		result := nbit.Binary(1)

		expected := []string{"0", "1"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected result to be %v, got %v", expected, result)
		}
	})

	t.Run("Binary with size 3", func(t *testing.T) {
		nbit := nbitstring.NBitString{}
		nbit.Array = make([]string, 3)
		result := nbit.Binary(3)

		expected := []string{"000", "001", "010", "011", "100", "101", "110", "111"}
		sort.Strings(expected)
		sort.Strings(result)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected result to be %v, got %v", expected, result)
		}
	})
}
