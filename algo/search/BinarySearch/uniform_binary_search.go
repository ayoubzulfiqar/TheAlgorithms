package binary

import (
	"fmt"
)

// Create a lookup table of the indices of the middle elements of the search space.
func lookupTable[T Numeric](array []T) []T {
	var lookupTable []T = make([]T, len(array))

	// Calculate the middle element index for each element in the array.
	for i := range array {
		// mid := int(math.Floor(float64(i+1) / 2.0))
		var mid T = T((i + 1) / 2.0)

		lookupTable[i] = mid
	}

	return lookupTable
}

// Perform a uniform binary search on the given array.
func uniformBinarySearch[T Numeric](array []T, target T, lookupTable []T) int {
	var low int = 0
	var high int = len(array) - 1

	// While the low pointer is less than or equal to the high pointer, continue the search.
	for low <= high {
		// Get the index of the middle element from the lookup table.
		var mid int = int(lookupTable[low])

		// If the target element is equal to the middle element, return the index of the middle element.
		if array[mid] == target {
			return mid
		}

		// If the target element is less than the middle element, set the high pointer to the middle element minus one.
		if target < array[mid] {
			high = mid - 1
		} else {
			// Otherwise, set the low pointer to the middle element plus one.
			low = mid + 1
		}
	}

	// If the target element was not found, return -1.
	return -1
}

func Uniform[T Numeric](arr []T, key T) int {
	var lookupTable []T = lookupTable(arr)
	var index int = uniformBinarySearch(arr, key, lookupTable)
	if index != -1 {
		fmt.Printf("Found Key: %v at Index: %v", key, index)
	} else {
		fmt.Printf("Key: %v not Found", key)
	}
	return index
}
