package sort

func Selection[T Numeric](array []T) []T {
	// looping through the whole array
	for i := 0; i < len(array)-1; i++ {
		// assuming the element in the array is min
		var minIndex int = i
		// looping though the array and checking step by step one at a time and compare it if it is minimum
		for j := i + 1; j < len(array); j++ {
			// element at position at j is less than like minimum to the previous value
			for array[j] < array[minIndex] {
				// we set it as our next min
				minIndex = j
			}
		}
		// if the are not same
		if minIndex != i {
			// than Swap if necessary
			array[i], array[minIndex] = array[minIndex], array[i]
		}
	}
	return array
}
