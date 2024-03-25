package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}

func findDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	hmap := make(map[int]bool)

	for l <= r {
		left := nums[l]
		right := nums[r]

		if right == left {
			return right
		}

		if _, ok := hmap[left]; ok {
			return left
		}

		if _, ok := hmap[right]; ok {
			return right
		}

		hmap[left] = true
		hmap[right] = true

		l++
		r--
	}

	return -1
}
