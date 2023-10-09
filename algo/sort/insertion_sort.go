package sort

func Insertion[T Numeric](array []T) []T {
	var n int = len(array)
	// Loop through the elements of the array, starting from index 1,
	// because at position 0, we assume the first element is already in the sorted portion.
	for currentIndex := 1; currentIndex < n; currentIndex++ {
		// key: Holds the current element of the array that we're comparing and potentially inserting.
		var key T = array[currentIndex]
		// sortedIndex: Represents the index of the last element in the sorted portion of the array,
		// which will be compared with the key.
		var sortedIndex int = currentIndex - 1

		// Compare the key with elements in the sorted portion and shift them to the right
		// until we find the correct position for the key.
		for sortedIndex >= 0 && array[sortedIndex] > key {
			array[sortedIndex+1] = array[sortedIndex]
			sortedIndex = sortedIndex - 1
		}
		// Insert the key into its correct position within the sorted portion.
		array[sortedIndex+1] = key
	}
	return array
}

func RecursiveInsertion[T Numeric](array []T) []T {
	var n int = len(array)
	if n <= 1 {
		return array
	}
	RecursiveInsertion(array[:n-1])
	var key T = array[n-1]
	var j int = n - 2
	for j >= 0 && array[j] > key {
		array[j+1] = array[j]
		j--
	}
	array[j+1] = key
	return array
}
