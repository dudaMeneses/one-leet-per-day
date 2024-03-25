package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
}

func lengthOfLongestSubstring(s string) int {
	// initialize left and right vars with 0 and max with 1
	l, r, max := 0, 0, 0

	// initialize map
	hmap := map[byte]int{}
	for r < len(s) {
		// if right char exists, position left to the right-most index for that char
		if pos, rok := hmap[s[r]]; rok {
			l = Max(pos, l)
		}

		max = Max(max, r-l+1)

		hmap[s[r]] = r + 1

		r++
	}

	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
