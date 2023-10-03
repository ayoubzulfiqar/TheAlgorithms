package sort

type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~string
}

func Bingo[T Numeric](arr []T) []T {
	n := len(arr)

	for i := 0; i < n; i++ {
		// Find index of minimum unsorted element
		minIndex := findMinUnsorted(arr, i, n)

		// Swap minimum with current element
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}

func findMinUnsorted[T Numeric](arr []T, start, end int) int {
	minIndex := start

	for i := start + 1; i < end; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}

	return minIndex
}
