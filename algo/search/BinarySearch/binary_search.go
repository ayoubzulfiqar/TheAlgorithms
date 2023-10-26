package binary

import (
	"fmt"
	"sync"
)

func Binary[T Numeric](arr []T, key T) int {

	var left, right int = 0, len(arr) - 1

	for left <= right {

		// Check bounds
		if left >= len(arr) || right < 0 {
			return -1
		}

		var mid int = left + (right-left)/2

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

// Recursive
func RecursiveBinarySearch[T Numeric](arr []T, key T) int {
	return binarySearch(arr, key, 0, len(arr))
}

func binarySearch[T Numeric](arr []T, key T, low, high int) int {
	if low >= high {
		return -1 // Element not found
	}

	var mid int = low + (high-low)/2

	if arr[mid] == key {
		return mid // Element found
	} else if arr[mid] > key {
		return binarySearch(arr, key, low, mid)
	} else {
		return binarySearch(arr, key, mid+1, high)
	}
}

// MultiThreaded

func BinarySearch[T Numeric](arr []T, target T) int {
	var numThreads int = 4 // setting default number of thread
	var wg sync.WaitGroup
	resultChan := make(chan int)

	subArraySize := len(arr) / numThreads

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		start := i * subArraySize
		end := start + subArraySize

		if i == numThreads-1 {
			end = len(arr)
		}
		// Slandered MultiThreaded Binary Search
		go func(arr []T, target T, left, right int, wg *sync.WaitGroup, result chan int) {
			defer wg.Done()

			for left <= right {
				mid := left + (right-left)/2

				if arr[mid] == target {
					result <- mid
					return
				} else if arr[mid] < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}

			result <- -1 // Element not found
		}(arr, target, start, end-1, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for index := range resultChan {
		if index != -1 {
			return index
		}
	}

	return -1
}
