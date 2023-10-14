package sort

func Bogo[T Numeric](array []T) []T {
	// while until not sorted we will keep sorting until we run run out of enough life on tis earth
	// to sustain humanity
	for !isBogoSorted(array) {
		// shuffle the array if not sorted
		shuffle(array)
	}
	return array
}

func isBogoSorted[T Numeric](array []T) bool {
	// looping through the elements of the array
	for i := 1; i < len(array); i++ {
		// if the element at index 0 is greater than the next one
		if array[i-1] > array[i] {
			// means not sorted
			return false
		}
	}
	// else sorted [almost]
	return true
}
