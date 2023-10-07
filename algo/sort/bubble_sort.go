package sort

// Iterative Bubble Sort
func Bubble[T Numeric](array []T) []T {
	var swapped bool = true
	// entire length of the array
	var n int = len(array)
	// continues Loop until all the elements inside the array is swapped or Sorted
	for swapped {
		// we assume that the array is not sorted
		swapped = false
		// loop thorough the array at position second
		for i := 1; i < n; i++ {
			// array[i-1]: the first element of the array
			// array[i]: second element of the array
			// Checking if first element of the array is greater than the second element than we need to swap
			if array[i-1] > array[i] {
				array[i-1], array[i] = array[i], array[i-1]
				// outer loop will continue till this condition reached
				swapped = true
			}
		}
		// After each pass, the largest element is at the end, so we reduce the range
		n--
	}
	return array
}

// Recursive Bubble Sort
func RecursiveBubble[T Numeric](array []T) []T {
	var n int = len(array)
	if n <= 1 {
		return array
	}
	for i := 1; i < n; i++ {
		if array[i-1] > array[i] {
			array[i-1], array[i] = array[i], array[i-1]
		}
	}
	RecursiveBubble(array[:n-1])
	return array
}
