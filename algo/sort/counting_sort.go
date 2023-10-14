package sort

func Counting[T Numeric](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	// Find the minimum and maximum values in the input array
	min, max := findMinMax(arr)

	// Initialize the counting array and output array
	counts := make([]T, int(max-min+1))
	output := make([]T, len(arr))

	// Count the occurrences of each element
	for _, num := range arr {
		counts[int(num-min)]++
	}

	// Calculate the cumulative counts
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	// Build the sorted output array
	for i := len(arr) - 1; i >= 0; i-- {
		output[int(counts[int(arr[i]-min)])-1] = arr[i]
		counts[int(arr[i]-min)]--
	}

	return output
}
