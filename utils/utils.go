package util

import "math/rand"

func GenerateInt(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000000) // Adjust the range as needed
	}
	return data
}



func GenerateFloat(size int) []float64 {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Float64() * 34 * 21 // Adjust the range as needed
	}
	return data
}