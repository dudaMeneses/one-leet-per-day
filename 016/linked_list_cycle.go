package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Example 1
	// 3 -> 2 -> 0 -> 4
	//      ^         |
	//      |_________|
	// Output: true
	// Explanation: There is a cycle in the linked list
	// 3 -> 2 -> 0 -> 4 -> 2 -> 0 -> 4 -> ...
	head := &ListNode{Val: 3}
	item2 := &ListNode{Val: 2}
	item3 := &ListNode{Val: 0}
	item4 := &ListNode{Val: 4}

	head.Next = item2
	item2.Next = item3
	item3.Next = item4
	item4.Next = item2 // cycle

	fmt.Println(hasCycle(head))
}

func hasCycle(head *ListNode) bool {
	hmap := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := hmap[head]; ok {
			return true
		}

		hmap[head] = true
		head = head.Next
	}

	return false
}
