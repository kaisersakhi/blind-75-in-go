package main

import (
	"testing"
)

func TestOnePass(t *testing.T) {
	testCases := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 4},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 1, 5, 3, 6, 4, 8}, 7},
		{[]int{}, 0},
	}

	for _, tc := range testCases {
		result := onePass(tc.prices)
		if result != tc.expected {
			t.Errorf("For prices %v, expected %d, but got %d", tc.prices, tc.expected, result)
		}
	}
}

func TestBruteForce(t *testing.T) {
	testCases := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 4},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 1, 5, 3, 6, 4, 8}, 7},
		{[]int{}, 0},
	}

	for _, tc := range testCases {
		result := bruteForce(tc.prices)
		if result != tc.expected {
			t.Errorf("For prices %v, expected %d, but got %d", tc.prices, tc.expected, result)
		}
	}
}
