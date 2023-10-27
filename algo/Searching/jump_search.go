package search

import (
	"runtime"
	"sync"
)

func Jump[T Numeric](arr []T, key T) int {
	var n int = len(arr)   // Get the length of the array.
	var step int = sqrt(n) // Calculate the jump step size based on the square root of n.

	var prev int = 0 // Initialize the 'prev' variable to 0.

	// The first loop finds the block where the element is or should be.
	for minStep := min(step, n) - 1; arr[minStep] < key; minStep = min(step, n) - 1 {
		prev = step
		step += sqrt(n)
		if prev >= n {
			return -1 // If 'prev' exceeds the length of the array, the key is not present.
		}
	}

	// The second loop performs a linear search within the block.
	for arr[prev] < key {
		prev++

		if prev == min(step, n) {
			return -1 // If 'prev' reaches the end of the block, the key is not found.
		}
	}

	if arr[prev] == key {
		return prev // If the key is found, return its index.
	}

	return -1 // If the key is not found in the array, return -1.
}

func min[T Numeric](x T, y T) T {
	if x <= y {
		return x // Return the smaller of the two values.
	}
	return y
}

func sqrt[T Numeric](x T) T {
	// This function approximates the square root of x using the Newton-Raphson method.
	var z T = T(rune(1.0)) // Initial guess for the square root.

	for i := 1; i <= 10; i++ { // Perform 10 iterations for a better approximation.
		z -= (z*z - x) / (2 * z) // Update the guess using the Newton-Raphson formula.
	}

	return z // Return the approximated square root.
}

// MultiThreaded Jump
// JumpParallel performs the search for the key in a specific range of the array.
func jumpParallel[T Numeric](arr []T, key T, startIndex int, endIndex int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	var prev int = startIndex

	for ; prev < endIndex; prev++ {
		if arr[prev] == key {
			resultChan <- prev // Send the result to the channel if the key is found.
			return
		}
	}

	resultChan <- -1 // Send -1 to indicate that the key was not found in this range.
}

// JumpMultithreaded performs a multithreaded jump search on the array.
func ParallelJump[T Numeric](arr []T, key T) int {
	var n int = len(arr)
	numThreads := runtime.NumCPU() // Determine the number of CPU cores available.

	resultChan := make(chan int, numThreads) // Create a channel for collecting search results.
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		startIndex := (n * i) / numThreads     // Calculate the start index for the current partition.
		endIndex := (n * (i + 1)) / numThreads // Calculate the end index for the current partition.

		wg.Add(1)                                                        // Increment the WaitGroup counter for the new goroutine.
		go jumpParallel(arr, key, startIndex, endIndex, &wg, resultChan) // Launch a goroutine to search this partition.
	}

	wg.Wait()         // Wait for all goroutines to finish.
	close(resultChan) // Close the result channel to signal that no more results will be sent.

	// Collect results from the channel. The first non-negative result indicates the index of the key.
	for result := range resultChan {
		if result != -1 {
			return result
		}
	}

	return -1 // If the key is not found in any partition, return -1.
}
