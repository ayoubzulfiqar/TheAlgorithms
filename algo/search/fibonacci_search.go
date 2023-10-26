package search

func Fibonacci[T Numeric](arr []T, x T) int {
	// Calculate the length of the input array.
	var n int = len(arr)

	// If the array is empty, return -1.
	if n == 0 {
		return -1
	}

	// Initialize Fibonacci variables.
	var fib1 int = 0       // First Fibonacci number
	var fib2 int = 1       // Second Fibonacci number
	var fib3 = fib1 + fib2 // Third Fibonacci number

	// Find the smallest Fibonacci number that is greater than or equal to n.
	for fib3 < n {
		fib1 = fib2
		fib2 = fib3
		fib3 = fib1 + fib2
	}

	// Initialize an offset variable to -1.
	var offset int = -1

	// Search for the element x in the array using Fibonacci search.
	for fib3 > 1 {
		// Calculate the index to be checked.
		var i int = min(offset+fib2, n-1)

		if arr[i] < x {
			// Update Fibonacci variables for a smaller segment.
			fib3 = fib2
			fib2 = fib1
			fib1 = fib3 - fib2
			offset = i
		} else if arr[i] > x {
			// Update Fibonacci variables for a larger segment.
			fib3 = fib1
			fib2 = fib2 - fib1
			fib1 = fib3 - fib2
		} else {
			// Element x found at index i.
			return i
		}
	}

	// Check if the last element in the array is equal to x.
	if fib2 == 1 && arr[offset+1] == x {
		return offset + 1
	} else {
		// Element x was not found in the array.
		return -1
	}
}
