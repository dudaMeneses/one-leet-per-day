package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDuplicates(nums []int) []int {
	hmap := make(map[int]bool)

	var duplicates []int
	for _, num := range nums {
		if _, ok := hmap[num]; ok {
			duplicates = append(duplicates, num)
		} else {
			hmap[num] = true
		}
	}

	return duplicates
}
