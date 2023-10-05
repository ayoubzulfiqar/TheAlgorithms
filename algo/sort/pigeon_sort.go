package sort

import "errors"

func Pigeon[T Numeric](array []T) ([]T, error) {
	if len(array) == 0 {
		return nil, errors.New("The Provided Array is Empty")
	}
	// Finding min and max with the help of helper function
	minValue, maxValue, minMaxErr := findMinMax(array)
	if minMaxErr != nil {
		return nil, minMaxErr
	}
	// Range of possible value - because This determine the Amount of bucket needed for the Collection
	var rangeOfValue T = maxValue - minValue + 1
	// we will create pigeon Holes based on range of value for possible collection of the Values
	pigeonHolesBucket := make([]T, int(rangeOfValue))
	// Iterate through the values of the original array
	for _, num := range array {
		// adding those values into the relative position
		// num-minValue := calculates the index of the pigeonhole where the current element should be placed.
		// and by incrementing ++ we will effectively place the each and every value inside the bucket
		pigeonHolesBucket[int(num-minValue)]++
	}

	/// This part is collecting the element from the Bucket and sorting
	//  it and placing it back to original array as Sorted

	// index is to keep track index at value of the Original Array
	var index int = 0
	// this loop iterate through the range of Values inside the Pigeon Hole in Ascending Order (small to Large)
	for i := 0; i < int(rangeOfValue); i++ {
		// It collect the smallest element from the bucket till pigeon hole is empty
		for pigeonHolesBucket[i] > 0 {
			// array[index] := is the starting position of the Original array
			// i := is the index of the position of the current pigeon Hole Bucket
			// minValue := minimum value in the Original array
			// By adding (i + minValue) minVal to i, we "de-normalize" the value, converting
			// it back to its original value within the range.

			array[index] = minValue + T(i)
			// Increment the Index for preparing the next element
			index++
			// Decrement the values from the bucket to original array till it's empty
			pigeonHolesBucket[i]--
		}
	}

	return array, nil
}

// Finding the min and max value inside the Unsorted Array
func findMinMax[T Numeric](array []T) (T, T, error) {
	// base case - if our array has nothing that It will return Nothing
	if len(array) == 0 {
		return T(0), T(0), errors.New("The Array is Empty: Nothing Found")
	}

	// Lets Assume that that the First element in the array is the min and also max

	var maxValue T = array[0]
	var minValue T = array[0]
	// We iterate though each and every value inside the array and compare it and assign the correct
	// and corresponding value to min and max
	for _, num := range array {
		if num < minValue {
			minValue = num
		}
		if num > maxValue {
			maxValue = num
		}
	}

	return minValue, maxValue, nil
}
