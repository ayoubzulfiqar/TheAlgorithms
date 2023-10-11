package sort

func RecursiveMerge[T Numeric](array []T) []T {
	// Base case: If the array has 0 or 1 element, it is already sorted.
	if len(array) <= 1 {
		return array
	}

	// Split the array into two halves
	var middle int = len(array) / 2
	var left []T = array[:middle]
	var right []T = array[middle:]

	// Recursively sort both halves
	left = RecursiveMerge(left)
	right = RecursiveMerge(right)

	// Merge the two sorted halves
	return merger(left, right)
}

func merger[T Numeric](left, right []T) []T {
	var merged []T = make([]T, 0, len(left)+len(right))

	// Initialize pointers for both halves
	i, j := 0, 0

	// Compare and merge elements from both halves
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	// Append any remaining elements from both halves
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}

// Iterative

func Merge[T Numeric](array []T) []T {
	if array == nil || len(array) <= 1 {
		return array
	}

	if len(array) > 1 {
		var mid int = len(array) / 2

		// Split left part
		var left []T = make([]T, mid)
		for i := 0; i < mid; i++ {
			left[i] = array[i]
		}

		// Split right part
		var right []T = make([]T, len(array)-mid)
		for i := mid; i < len(array); i++ {
			right[i-mid] = array[i]
		}

		Merge(left)
		Merge(right)

		var i int = 0
		var j int = 0
		var k int = 0

		// Merge left and right arrays
		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				array[k] = left[i]
				i++
			} else {
				array[k] = right[j]
				j++
			}
			k++
		}

		// Collect remaining elements
		for i < len(left) {
			array[k] = left[i]
			i++
			k++
		}

		for j < len(right) {
			array[k] = right[j]
			j++
			k++
		}
	}
	return array
}
