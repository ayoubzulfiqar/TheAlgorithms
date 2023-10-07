package sort

func CockTail[T Numeric](array []T) []T {
	// Initialize variables
	var swapped bool = true         // Flag to check if any swaps are made in a pass
	var start int = 0              // Start of the unsorted portion of the array
	var end int = len(array)       // End of the unsorted portion of the array

	// Continue sorting until no swaps are made in a pass
	for swapped {
		swapped = false // Reset the swapped flag at the beginning of each pass

		// Forward pass: Move the largest element to the end
		for i := start; i < end-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i] // Swap elements if they are out of order
				swapped = true // Mark as swapped
			}
		}

		// If no swaps were made in the forward pass, the array is sorted
		if !swapped {
			break
		}

		swapped = false // Reset the swapped flag for the backward pass
		end = end - 1   // Decrease the end pointer as the largest element is now at the end

		// Backward pass: Move the smallest element to the beginning
		for i := end - 1; i >= start; i-- {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i] // Swap elements if they are out of order
				swapped = true // Mark as swapped
			}
		}

		start = start + 1 // Increase the start pointer as the smallest element is now at the beginning
	}

	return array // Return the sorted array
}

