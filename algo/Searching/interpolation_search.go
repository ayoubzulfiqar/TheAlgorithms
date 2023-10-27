package search

import (
	"runtime"
	"sync"
)

// Interpolation
func Interpolation[T Numeric](arr []T, key T) int {
	var low int = 0
	var high int = len(arr) - 1

	for low <= high && key >= arr[low] && key <= arr[high] {
		if low == high {
			// If the search range narrows down to a single element, check if it's the key.
			if arr[low] == key {
				return low
			}
			// If the single element is not the key, return -1 (key not found).
			return -1
		}

		// Calculate the position estimate using interpolation formula.
		var position int = low + ((int(key)-int(arr[low]))*(high-low))/(int(arr[high])-int(arr[low]))

		// Perform a bound check to ensure 'position' is within the array bounds.
		if position < low || position > high {
			return -1
		}

		if arr[position] == key {
			// If the key is found at the estimated position, return that position.
			return position
		}

		if arr[position] < key {
			// If the estimated value is less than the key, narrow the search range to the right.
			low = position + 1
		} else {
			// If the estimated value is greater than the key, narrow the search range to the left.
			high = position - 1
		}
	}

	// If the key is not found within the search range, return -1 (key not found).
	return -1
}

// Recursive Interpolation

func RecursiveInterpolation[T Numeric](arr []T, key T, high int, low int) int {
	if low <= high && key >= arr[low] && key <= arr[high] {
		if low == high {
			if arr[low] == key {
				return low
			}
			return -1
		}

		var position int = low + ((int(key)-int(arr[low]))*(high-low))/(int(arr[high])-int(arr[low]))

		if arr[position] == key {
			return position
		}

		if arr[position] < key {
			return RecursiveInterpolation(arr, key, position+1, high)
		}

		return RecursiveInterpolation(arr, key, low, position-1)
	}

	return -1
}

///

func ParallelInterpolation[T Numeric](arr []T, key T) int {
	// Determine the number of CPU cores available.
	numThreads := runtime.NumCPU()

	// Initialize variables for the result and a mutex for synchronization.
	var result int
	var resultMutex sync.Mutex
	var wg sync.WaitGroup

	// Function to perform interpolation search within a partition of the array.
	searchPartition := func(partitionIndex int) {
		defer wg.Done()

		// Calculate the search range for the current partition based on the number of threads.
		low := (len(arr) * partitionIndex) / numThreads
		high := (len(arr)*(partitionIndex+1))/numThreads - 1

		localResult := -1 // Initialize a local result for this partition.

		// Interpolation search within the partition.
		for low <= high && key >= arr[low] && key <= arr[high] {
			if low == high {
				if arr[low] == key {
					localResult = low // Set the local result if the key is found in this partition.
				}
				break
			}

			// Calculate the estimated position using interpolation formula.
			var position int = low + ((int(key)-int(arr[low]))*(high-low))/(int(arr[high])-int(arr[low]))

			if arr[position] == key {
				localResult = position // Set the local result if the key is found at the estimated position.
				break
			}

			if arr[position] < key {
				low = position + 1 // Narrow the search range to the right.
			} else {
				high = position - 1 // Narrow the search range to the left.
			}
		}

		// Update the overall result using a mutex to ensure that only the first found result is considered.
		resultMutex.Lock()
		if localResult != -1 {
			result = localResult
		}
		resultMutex.Unlock()
	}

	wg.Add(numThreads) // Increment the wait group counter to track goroutines.

	// Launch multiple goroutines to search different partitions in parallel.
	for i := 0; i < numThreads; i++ {
		go searchPartition(i)
	}

	wg.Wait() // Wait for all goroutines to finish.

	return result // Return the final result.
}
