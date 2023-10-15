package sort

func Redix[T Numeric](arr []T) []T {
	max := getMax(arr)

	// Iterate through each digit place, from least significant to most significant
	for exp := 1; max/T(exp) > 0; exp *= 10 {
		countingSort(arr, T(exp))
	}
	return arr
}

func getMax[T Numeric](arr []T) T {
	max := arr[0] // Initialize max with a large positive integer
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func countingSort[T Numeric](arr []T, exp T) []T {
	n := len(arr)
	output := make([]T, n)
	count := make([]T, 19) // 19 for [-9, 9]

	for i := 0; i < n; i++ {
		count[int(arr[i]/exp)%10+9]++
	}

	for i := 1; i < 19; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[int(count[int(arr[i]/exp)%10+9])-1] = arr[i]
		count[int(arr[i]/exp)%10+9]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
	return output
}
