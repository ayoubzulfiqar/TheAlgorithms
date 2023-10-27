package binary

import "fmt"

func Ubiquitous[T Numeric](arr []T, key T) int {

	var low int = 0
	var high int = len(arr) - 1
	if low >= len(arr) || high < 0 {
		return -1
	}
	for low <= high {
		var mid int = high - (low+high)/2
		if mid >= len(arr) {
			mid = len(arr) - 1
		}
		if arr[mid] < key {
			low = mid + 1
			if low < len(arr) && arr[low] == key {
				fmt.Printf("Found key: %v at Index %v\n", key, low)
				return low
			}
		} else if arr[mid] > key {
			high = mid - 1
			if high >= 0 && arr[high] == key {
				fmt.Printf("Found key %v at index %v\n", key, high)
				return high
			}
		} else {
			fmt.Printf("Found key %v at index %v\n", key, mid)
			return mid // value found
		}
	}
	fmt.Printf("Key %v not found in the array\n", key)

	return -1 // value not found
}
