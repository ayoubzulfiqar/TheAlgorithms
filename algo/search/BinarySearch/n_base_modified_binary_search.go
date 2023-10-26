package binary

import "fmt"

func NumericBaseBinarySearch[T Numeric](arr []T, key T, base int) int {
	var N = 3
	var PowerOf2 = 4
	var i, step1, step2, times int

	// Find the first power of N greater than the array size
	for step1 = 1; step1 < base; step1 *= N {
	}

	for i = 0; step1 > 0; step1 /= N {
		// Each time a power can be used, count how many times it can be used
		if i+step1 < base && arr[i+step1] <= key {
			for times = 1; step2 > 0; step2 >>= 1 {
				if i+(step1*(times+step2)) < base && arr[i+(step1*(times+step2))] <= key {
					times += step2
				}
			}
			step2 = PowerOf2

			// Add to the final result how many times the power can be used
			i += times * step1
		}
	}

	// Return the index if the element is present in the array, else return -1
	if arr[i] == key {
		fmt.Printf("Key: %v Found at Index: %v", key, i)
		return i
	}
	fmt.Printf("Key: %v Not Found", key)
	return -1
}
