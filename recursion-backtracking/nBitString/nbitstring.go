/*
	Problem 3
	Generate all the strings of n bits. Assume A[0..n – 1] is an array of size n.
*/

package nbitstring

import (
	"strings"
)

type NBitString struct {
	Array []string
}

func (nbit *NBitString) Binary(size int) []string {
	if size < 1 {
		return []string{strings.Join(nbit.Array, "")}
	} else {
		nbit.Array[size-1] = "0"
		zeroStrings := nbit.Binary(size - 1)
		nbit.Array[size-1] = "1"
		oneStrings := nbit.Binary(size - 1)
		result := append(zeroStrings, oneStrings...)
		return result
	}
}
