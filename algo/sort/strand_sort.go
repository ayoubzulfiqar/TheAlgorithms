package sort

// Strand is a sorting algorithm that takes an unsorted integer slice and returns a sorted slice.
func Strand[T Numeric](arr []T) []T {
	// If the input slice has 1 or fewer elements, it's already sorted, so return it as is.
	if len(arr) <= 1 {
		return arr
	}

	// Initialize the result slice with the first element of the input.
	result := []T{arr[0]}
	// Remove the first element from the input slice.
	arr = arr[1:]

	// Continue processing the input slice.
	for len(arr) > 0 {
		// Initialize a subList with the first element of the input.
		subList := []T{arr[0]}
		// Remove the first element from the input slice.
		arr = arr[1:]

		// Iterate through the input slice to find elements that can be added to the subList.
		i := 0
		for i < len(arr) {
			// If the current element is greater than or equal to the last element of the subList, add it to the subList.
			if arr[i] >= subList[len(subList)-1] {
				subList = append(subList, arr[i])
				// Remove the added element from the input slice.
				arr = append(arr[:i], arr[i+1:]...)
			} else {
				i++
			}
		}

		// Merge the subList into the result using the sMerge function.
		result = sMerge(result, subList)
	}

	return result
}

// sMerge is a helper function to merge two sorted slices 'a' and 'b' into a single sorted slice.
func sMerge[T Numeric](a []T, b []T) []T {
	// Initialize a result slice to hold the merged values.
	result := make([]T, 0, len(a)+len(b))
	i, j := 0, 0

	// Merge the two slices while maintaining the sorted order.
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	// Append any remaining elements from slices 'a' and 'b' to the result.
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}
