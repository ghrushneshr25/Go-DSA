package kstring

import (
	"fmt"
	"strings"
)

type KString struct {
	Array []string
}

func (kstr *KString) KStringRecursion(size, k int) []string {
	if size < 1 {
		if len(kstr.Array) > 0 {
			return []string{strings.Join(kstr.Array, "")}
		}
		return []string{}
	}
	result := []string{}
	for i := 0; i < k; i++ {
		kstr.Array[size-1] = fmt.Sprintf("%v", i)
		current := kstr.KStringRecursion(size-1, k)
		result = append(result, current...)
	}
	return result
}
