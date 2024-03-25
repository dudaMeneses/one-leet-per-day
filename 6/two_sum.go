package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		if aux, ok := hmap[diff]; ok {
			return []int{aux, i}
		}

		hmap[num] = i
	}

	return []int{}
}
