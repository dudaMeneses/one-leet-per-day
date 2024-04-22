package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	var result int
	local := 0
	for i, num := range nums {
		if i == 0 {
			local = num
			result = local
		} else {
			local = max(num, num+local)
			if local > result {
				result = local
			}
		}
	}

	return result
}
