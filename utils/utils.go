package util

import "math/rand"

func GenerateRandomData(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000000) // Adjust the range as needed
	}
	return data
}
