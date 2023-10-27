package search

import (
	"fmt"
)

func Ternary[T Numeric](arr []T, key T) int {
	// Initialize low and high as the first and last indices of the array.
	var low int = 0
	var high int = len(arr) - 1

	// Perform a loop until low is less than or equal to high.
	for low <= high {
		// Calculate the first and second midpoints.
		var firstMid int = low + (high-low)/3
		var secMid int = high - (high-low)/3

		// If the key is found at the first midpoint, return its index.
		if arr[firstMid] == key {
			return firstMid
		}

		// If the key is found at the second midpoint, return its index.
		if arr[secMid] == key {
			return secMid
		}

		// If the key is smaller than the element at the first midpoint, update 'high.'
		if key < arr[firstMid] {
			high = firstMid - 1
		} else if key > arr[secMid] {
			// If the key is larger than the element at the second midpoint, update 'low.'
			low = secMid + 1
		} else {
			// If the key is between the elements at firstMid and secMid, adjust 'low' and 'high.'
			low = firstMid + 1
			high = secMid - 1
		}
	}

	// If the key is not found in the array, return -1.
	return -1
}

// Recursive

func ternarySearch[T Numeric](arr []T, key T, low int, high int) int {

	if low > high {
		fmt.Printf("Key: %v Not found\n", key)
		return -1
	}

	var mid1 int = low + (high-low)/3
	var mid2 int = high - (high-low)/3

	if arr[mid1] == key {
		fmt.Printf("Key: %vFound at index %v\n", key, mid1)
		return mid1
	}

	if arr[mid2] == key {
		fmt.Printf("Key: %v Found at index %v\n", key, mid2)
		return mid2
	}

	if key < arr[mid1] {
		return ternarySearch(arr, key, low, mid1-1)
	} else if key > arr[mid2] {
		return ternarySearch(arr, key, mid2+1, high)
	} else {
		return ternarySearch(arr, key, mid1+1, mid2-1)
	}

}

// Helper function
func RecursiveTernary(arr []int, key int) int {
	return ternarySearch(arr, key, 0, len(arr)-1)

}

// MultiThreaded

func ParallelTernary[T Numeric](arr []T, key T) int {
	// Create a channel to communicate results between goroutines.
	var goroutines int = 4 // Default Adjust According to YOUR NEED
	resultChan := make(chan int, goroutines)
	var low int = 0
	var high int = len(arr) - 1

	// Function for a single search segment.
	searchSegment := func(start int, end int) {
		for i := start; i <= end; i++ {
			// Calculate the first and second midpoints for this segment.
			firstMid := i + (end-i)/3
			secMid := end - (end-i)/3

			// If the key is found at the first midpoint, send its index to the channel.
			if arr[firstMid] == key {
				resultChan <- firstMid
				return
			}

			// If the key is found at the second midpoint, send its index to the channel.
			if arr[secMid] == key {
				resultChan <- secMid
				return
			}

			// If the key is smaller than the element at the first midpoint, update 'end.'
			if key < arr[firstMid] {
				end = firstMid - 1
			} else if key > arr[secMid] {
				// If the key is larger than the element at the second midpoint, update 'start.'
				i = secMid + 1
			} else {
				// If the key is between the elements at firstMid and secMid, adjust 'start' and 'end.'
				i = firstMid + 1
				end = secMid - 1
			}
		}
		resultChan <- -1
	}

	// Divide the search space among goroutines.
	for i := 0; i < goroutines; i++ {
		segmentSize := (high - low + 1) / goroutines
		start := i * segmentSize
		end := start + segmentSize - 1

		// The last goroutine might cover extra elements if the division isn't exact.
		if i == goroutines-1 {
			end = high
		}

		// Start a goroutine to search its segment.
		go searchSegment(start, end)
	}

	// Wait for results from goroutines and close the result channel.
	for i := 0; i < goroutines; i++ {
		result := <-resultChan
		if result != -1 {
			return result
		}
	}
	close(resultChan)

	// If the key is not found in the array, return -1.
	return -1
}
