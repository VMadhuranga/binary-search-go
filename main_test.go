package main

import (
	"testing"
)

func createSortedArr(to int) []int {
	sortedArray := []int{}
	for i := range to {
		sortedArray = append(sortedArray, i)
	}
	return sortedArray
}

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		arr      []int
		target   int
		expected bool
	}{
		{
			arr:      createSortedArr(200),
			target:   10,
			expected: true,
		},
		{
			arr:      createSortedArr(20000),
			target:   -1,
			expected: false,
		},
		{
			arr:      []int{},
			target:   15,
			expected: false,
		},
		{
			arr:      []int{0},
			target:   0,
			expected: true,
		},
		{
			arr:      []int{-2, -1},
			target:   -1,
			expected: true,
		},
		{
			arr:      createSortedArr(2000000),
			target:   105028,
			expected: true,
		},
		{
			arr:      createSortedArr(2000000),
			target:   2000001,
			expected: false,
		},
	}

	for _, c := range cases {
		got := binarySearch(c.arr, c.target)

		if got != c.expected {
			t.Errorf("got: %v != expected: %v", got, c.expected)
		}
	}
}
