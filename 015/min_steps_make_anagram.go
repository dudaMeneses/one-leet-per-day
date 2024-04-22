package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSteps("bab", "aba"))
}

func minSteps(first string, second string) int {
	var arr = make([]int, 26)
	n := len(first)

	for i := 0; i < n; i++ {
		arr[first[i]-'a']++
		arr[second[i]-'a']--
	}

	result := 0

	for _, x := range arr {
		if x > 0 {
			result += x
		}
	}

	return result
}
