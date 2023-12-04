package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {
	if len(s) < 2 || numRows < 2 {
		return s
	}

	rows := make([]string, numRows)
	curr := 0
	dir := -1

	for i := 0; i < len(s); i++ {
		if curr == 0 || curr == numRows-1 {
			dir *= -1
		}

		rows[curr] += string(s[i])

		curr += dir
	}

	res := ""

	for _, l := range rows {
		res += l
	}

	return res
}
