package sort

import "math/rand"

// Helper function to generate a random array of a given size
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(10000) // Adjust the upper bound as needed
	}
	return arr
}

func generateRandomFloatArray(size int) []float64 {
	arr := make([]float64, size)
	rangeSize := 50.0 - 10.0

	for i := 0; i < size; i++ {
		arr[i] = rangeSize + rand.Float64()*545*62
	}

	return arr
}
