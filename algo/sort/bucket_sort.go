package sort

func Bucket[T Numeric](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	// Find the minimum and maximum values in the array
	minValue, maxValue := findMinMax(arr)

	// Create buckets
	numBuckets := maxValue - minValue + 1
	buckets := make([][]T, int(numBuckets))

	// Distribute elements into buckets
	for _, val := range arr {
		buckets[int(val-minValue)] = append(buckets[int(val-minValue)], val)
	}

	// Sort each bucket (using any sorting algorithm)
	sorted := make([]T, 0)
	for i := 0; i < int(numBuckets); i++ {
		if len(buckets[i]) > 0 {
			buckets[i] = Insertion(buckets[i]) // Use any sorting method for the buckets
			sorted = append(sorted, buckets[i]...)
		}
	}

	return sorted
}
