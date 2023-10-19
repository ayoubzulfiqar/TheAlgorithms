package sort

// Generic Implementation is Slow
func Tim[T Numeric](array []T) []T {
	var n int = len(array)
	var minRun int = findMinRun(n)

	// Sort individual sub-arrays of size minRun
	for start := 0; start < n; start += minRun {
		end := start + minRun
		if end > n {
			end = n
		}
		// We have already Implemented Insertion Sort Long Time ago
		Insertion(array[start:end])
	}

	// Now merge the sorted runs
	var size int = minRun
	for size < n {
		// Choose merges sized sz, sz*2, sz*4, until sz >= n
		for start := 0; start < n; start += 2 * size {
			mid := start + size
			end := min(start+2*size, n)
			merging(array[start:mid], array[mid:end])
		}
		size *= 2
	}
	return array
}

func findMinRun(n int) int {
	const minRun int = 64
	r := 0
	for n >= minRun {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

func min[T Numeric](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

// func insertionSort(array []int) {
// 	var j int

// 	for i := range array {
// 		j = i
// 		for j > 0 && array[j-1] > array[j] {
// 			array[j-1], array[j] = array[j], array[j-1]
// 			j--
// 		}
// 	}
// }

func merging[T Numeric](left []T, right []T) []T {
	// Merge two sorted sub-arrays into one
	var result []T = make([]T, len(left)+len(right))
	leftIndex, rightIndex, resultIndex := 0, 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result[resultIndex] = left[leftIndex]
			leftIndex++
		} else {
			result[resultIndex] = right[rightIndex]
			rightIndex++
		}
		resultIndex++
	}

	// Copy remaining elements from either a or b
	for leftIndex < len(left) {
		result[resultIndex] = left[leftIndex]
		leftIndex++
		resultIndex++
	}
	for rightIndex < len(right) {
		result[resultIndex] = right[rightIndex]
		rightIndex++
		resultIndex++
	}

	// Copy tmp back to a and b
	copy(left, result)
	copy(right, result[len(left):])
	return result
}
