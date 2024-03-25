package main

import "fmt"

func main() {
	fmt.Printf("Done in: %v\n", removePalindromeSub("aababbaba"))
}

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}

	for i, j := 0, len(s)-1; i < j; i++ {
		if s[i] != s[j] {
			return 2
		}
		j--
	}

	return 1
}
