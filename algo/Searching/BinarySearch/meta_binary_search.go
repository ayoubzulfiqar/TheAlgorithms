package binary

func Meta[T Numeric](arr []T, key T) int {
	var n int = len(arr)       // Get the length of the array
	var lg int = log2(n-1) + 1 // Calculate the number of bits needed to represent the largest array index

	var pos int = 0            // Initialize the position variable to start searching from the first element
	for i := lg; i >= 0; i-- { // Iterate from the most significant bit to the least significant bit
		if arr[pos] == key { // If the current position contains the key, return the index
			return pos
		}

		var newPos int = pos | (1 << i) // Construct a new position by setting the i-th bit

		if newPos < n && arr[newPos] <= key { // If the new position is within bounds and its value is less than or equal to the key
			pos = newPos // Update the current position to the new position
		}
	}

	if arr[pos] == key { // Check if the key is found at the final position
		return pos
	}

	return -1 // If the key is not found, return -1
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Calculation of log2
func log2[I Integer](x I) I {
	if x <= 0 {
		return 0 // Invalid input; you can choose any suitable error value
	}

	var result I = 0
	for x > 1 {
		x >>= 1 // Right shift x by 1, which is equivalent to dividing by 2
		// x = x/T(2)

		result++
	}

	return result
}
