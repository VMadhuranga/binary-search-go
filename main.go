package main

import (
	"math"
)

func binarySearch(arr []int, targetVal int) bool {
	lowIdx := 0
	highIdx := len(arr) - 1

	for lowIdx <= highIdx {
		midIdx := int(math.Floor((float64(lowIdx + highIdx)) / 2))
		midVal := arr[midIdx]

		if midVal == targetVal {
			return true
		}

		if midVal < targetVal {
			lowIdx = midIdx + 1
		} else {
			highIdx = midIdx - 1
		}
	}

	return false
}
