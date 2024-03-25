package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2})) // 2.0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge arrays with insertion sort
	arr := merge(nums1, nums2)

	// find median point
	avg := float64(len(arr) / 2)
	if len(arr)%2 == 0 {
		//  if even sum neighbour values and divide it by two
		return (float64(arr[int(avg)-1]) + float64(arr[int(avg)])) / 2
	} else {
		//  if odd then return number
		return float64(arr[int(math.Round(avg))])
	}

}

func merge(a []int, b []int) []int {
	i, j := 0, 0

	var final []int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
