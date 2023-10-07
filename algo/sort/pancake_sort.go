package sort

func PanCake[T Numeric](array []T) []T {
	var n int = len(array)
	// looping at the back of the array then we will decrement as we move towards the left
	for i := n; i > 1; i-- {
		// we will look for the max element inside the array from the backside
		var maxIndex int = findMax(array[:i])
		// if the last element of the array is not the max
		// i-1: represent the last element of the array
		if maxIndex != i-1 {
			// it reverses the order of elements from the beginning of the array up to the largest pancake.
			// The largest pancake will end up at the top of the stack (at index 0).
			array = flip(array, maxIndex)
			// It flips the portion of the array from index 0 to i-1.
			// This flip operation effectively moves the largest element to its correct position at
			// the bottom of the sorted portion.
			array = flip(array, i-1)
		}
	}
	return array
}

// Finding the Max Element inside the Array
func findMax[T Numeric](array []T) int {
	var maxIndex int = 0
	// loop through the element of the array
	for i := 1; i < len(array); i++ {
		// if the element is inside the array is larger than the our max Value
		if array[i] > array[maxIndex] {
			// we found new larger element
			maxIndex = i
		}
	}
	return maxIndex

}

// This function will reverse the elements in the slice from index 0 to k inclusive.
func flip[T Numeric](array []T, k int) []T {
	// start : starting index of the Array
	var start int = 0
	// end: till at the end index of the array
	var end int = k
	// loop through left to right, if element are smaller than right element
	for start < end {
		// we swap
		array[start], array[end] = array[end], array[start]
		// moving forward toward the right array
		start++
		// moving back to left side of the array
		end--
	}
	return array
}

// Recursive
func RecursivePanCake[T Numeric](array []T) []T {
	var n int = len(array)
	if n == 1 {
		return array
	}
	var maxIndex int = 0
	for i := 0; i < n; i++ {
		if array[i] > array[maxIndex] {
			maxIndex = i
		}
	}
	if maxIndex != 0 {
		array = recursiveFlip(array, maxIndex)
	}
	array = recursiveFlip(array, n-1)
	RecursivePanCake(array[:n-1])
	return array
}
func recursiveFlip[T Numeric](array []T, n int) []T {
	// start : starting index of the Array
	var start int = 0
	// end: till at the end index of the array
	var end int = n
	// loop through left to right, if element are smaller than right element
	for start < end {
		// we swap
		array[start], array[end] = array[end], array[start]
		// moving forward toward the right array
		start++
		// moving back to left side of the array
		end--
	}
	return array
}
