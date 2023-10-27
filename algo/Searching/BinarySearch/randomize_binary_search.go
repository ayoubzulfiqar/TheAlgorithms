package binary

import (
	"fmt"
	"math/rand"
)

func getRandom(x int, y int) int {
	return x + (rand.Intn(y - x + 1))
}

func RandomBinary[T Numeric](arr []T, key T) int {
	var left, right int = 0, len(arr) - 1

	for left <= right {

		// Check bounds
		if left >= len(arr) || right < 0 {
			return -1
		}

		var mid int = getRandom(left, right)

		// Handle mid out of bounds
		if mid >= len(arr) {
			mid = len(arr) - 1
		}

		if arr[mid] == key {
			fmt.Printf("Found key %v at index %v\n", key, mid)
			return mid
		}

		if arr[mid] < key {
			left = mid + 1
			// Check left in bounds
			if left < len(arr) && arr[left] == key {
				fmt.Printf("Found key %v at index %v\n", key, left)
				return left
			}
		} else {
			right = mid - 1
			// Check right in bounds
			if right >= 0 && arr[right] == key {
				fmt.Printf("Found key %v at index %v\n", key, right)
				return right
			}
		}
	}
	fmt.Printf("Key %v not found in the array\n", key)
	return -1
}
