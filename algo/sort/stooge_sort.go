package sort

func Stooge[T Numeric](array []T) []T {
	// Initialize the low and high indices for the entire array.
	var low int = 0
	var high int = len(array)

	// Check if the last element is smaller than the first element, swap them if needed.
	if array[high-1] < array[low] {
		temp := array[low]
		array[low] = array[high-1]
		array[high-1] = temp
	}

	// Calculate the length of the current sub-array.
	length := high - low

	// Check if there are more than 2 elements in the sub-array.
	if length > 2 {
		// Calculate one-third of the length.
		third := length / 3

		// Recursively sort the first 2/3 of the sub-array.
		Stooge(array[low : high-third])

		// Recursively sort the last 2/3 of the sub-array.
		Stooge(array[low+third : high])

		// Recursively sort the first 2/3 of the sub-array again.
		Stooge(array[low : high-third])
	}

	// Return the sorted sub-array.
	return array
}
