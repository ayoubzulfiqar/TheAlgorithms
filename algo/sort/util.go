package sort

import "math/rand"

// shuffle function shuffles the array to mitigate worst-case scenarios.
func shuffle[T Numeric](arr []T) {
	// Initialize the random number generator with a seed for consistency.
	rand.New(rand.NewSource(99))

	// Iterate through the array and swap elements randomly to shuffle.
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
