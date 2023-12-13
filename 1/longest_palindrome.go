package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var (
		start int
		end   int
	)

	for i := 0; i < len(s); i++ {
		oddPalindrome := expandAroundCenter(s, i, i)
		evenPalindrome := expandAroundCenter(s, i, i+1)

		maxLen := highest(oddPalindrome, evenPalindrome)

		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func highest(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
