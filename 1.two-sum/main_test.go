package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 13, []int{0, 2}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3}, 7, []int{-1, -1}},
		{[]int{1, 4, 55, 6, 9, 2, 8, -10}, 12, []int{1, 6}},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		if !isEqual(result, test.expected) {
			t.Errorf("For nums %v and target %d, expected %v, but got %v", test.nums, test.target, test.expected, result)
		}
	}
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
