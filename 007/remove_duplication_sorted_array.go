package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func removeDuplicates(nums []int) int {
	hmap := make(map[int]int)
	var k []int
	last := 0

	for i, num := range nums {
		if _, ok := hmap[num]; !ok {
			hmap[num] = 1
			k = append(k, num)

			if last != i {
				rep := nums[last]
				nums[last] = num
				nums[i] = rep
			}

			last++
		}
	}

	return len(k)
}
