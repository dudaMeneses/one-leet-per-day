package main

import "testing"

func TestDirectPath(t *testing.T) {
	value := LargestPathValue("abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}})

	if value != 3 {
		t.Errorf("Expected value 3 but it was %v\n", value)
	}
}

func TestCycle(t *testing.T) {
	value := LargestPathValue("a", [][]int{{0, 0}})

	if value != -1 {
		t.Errorf("Expected value -1 but it was %v\n", value)
	}
}
