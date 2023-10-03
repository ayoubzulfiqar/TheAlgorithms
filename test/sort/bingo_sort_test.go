package sort

import (
	"math/rand"
	"reflect"
	"testing"

	sorts "github.com/ayoubzulfiqar/TheAlgorithms/algo/sort"
)

func TestBingoSort(t *testing.T) {
	// Test case: Unsorted array
	unsortedArray := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	expectedSortedArray := []int{2, 3, 5, 6, 7, 9, 10, 11, 12, 14}

	sortedArray := sorts.Bingo(unsortedArray)

	// Check if the sorted array matches the expected result
	if !reflect.DeepEqual(sortedArray, expectedSortedArray) {
		t.Errorf("BingoSort(%v) = %v; want %v", unsortedArray, sortedArray, expectedSortedArray)
	}
}

// BenchmarkBingoSort benchmarks the BingoSort function.
func BenchmarkBingoSort(b *testing.B) {
	// Generate a random unsorted array for each benchmark iteration
	unsortedArray := generateRandomArray(100) // Change the array size as needed

	for i := 0; i < b.N; i++ {
		// Make a copy of the unsorted array for each iteration to ensure fairness
		arrayCopy := make([]int, len(unsortedArray))
		copy(arrayCopy, unsortedArray)

		// Run the sorting algorithm and measure the time it takes
		b.StartTimer()
		sorts.Bingo(arrayCopy)
		b.StopTimer()
	}
}

// Helper function to generate a random array of a given size
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(10000) // Adjust the upper bound as needed
	}
	return arr
}
