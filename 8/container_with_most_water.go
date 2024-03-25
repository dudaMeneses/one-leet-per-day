package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	l, r, max := 0, len(height)-1, 0

	for l < r {
		if new := calcArea(l, r, height); new > max {
			max = new
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return max
}

func calcArea(l int, r int, arr []int) int {
	left, right := arr[l], arr[r]

	if left < right {
		return left * (r - l)
	} else {
		return right * (r - l)
	}
}
