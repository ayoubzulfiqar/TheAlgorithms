package sort

import "math/rand"

func ThreeWayQuick(array []int) []int {
	return array
}

func partition(array []int) (int, int) {
	return 0, 0
}

// It shuffles the array before sorting, which ensures that the algorithm performs well on already
// sorted or partially sorted input.
func shuffle(array []int) []int {
	rand.New(rand.NewSource(99))
	for i := len(array) - 1; i > 0; i-- {
		randomElement := rand.Intn(i + 1)
		array[i], array[randomElement] = array[randomElement], array[i]
	}
	return array
}
