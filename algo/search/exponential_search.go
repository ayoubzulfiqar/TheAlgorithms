package search

func Exponential[T Numeric](arr []T, key T) int {
	// Get the length of the array.
	var n int = len(arr)

	// If the array is empty, return -1 (key not found).
	if n == 0 {
		return -1
	}

	// Initialize the 'index' variable to 1 to start exponential search.
	var index int = 1

	// Double the 'iter' value until it's within the array bounds and the element at 'iter' is less than 'key'.
	for index < n && arr[index] < key {
		index *= 2
	}

	// Define 'left' as half of the 'iter' and 'right' as the minimum of 'iter' and the last index in the array.
	var left int = index / 2
	var right int = min(index, n-1)

	// Perform binary search within the range defined by 'left' and 'right'.
	for left <= right {
		// Calculate the middle index.
		var mid int = (left + right) / 2

		// Check if the element at 'mid' is equal to 'key'.
		if arr[mid] == key {
			// If the element is found at 'mid', return its index.
			return mid
		} else if arr[mid] < key {
			// If the element at 'mid' is less than 'key', narrow the search range to the right.
			left = mid + 1
		} else {
			// If the element at 'mid' is greater than 'key', narrow the search range to the left.
			right = mid - 1
		}
	}

	// If the key is not found, return -1.
	return -1
}
