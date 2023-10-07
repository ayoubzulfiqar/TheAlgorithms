package sort

import (
	"reflect"
	"testing"

	sorts "github.com/ayoubzulfiqar/TheAlgorithms/algo/sort"
)

func TestBubbleIntSort(t *testing.T) {
	// Test case: Unsorted array
	unsortedArray := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	expectedSortedArray := []int{2, 3, 5, 6, 7, 9, 10, 11, 12, 14}

	sortedArray := sorts.Bubble(unsortedArray)

	// Check if the sorted array matches the expected result
	if !reflect.DeepEqual(sortedArray, expectedSortedArray) {
		t.Errorf("PigeonSort(%v) = %v; want %v", unsortedArray, sortedArray, expectedSortedArray)
	}
}

func TestBubbleFloatSort(t *testing.T) {
	// Test case: Unsorted array
	unsortedArray := []float64{9.0, 7.0, 5.0, 11.0, 12.0, 2.0, 14.0, 3.0, 10.0, 6.0}
	expectedSortedArray := []float64{2.0, 3.0, 5.0, 6.0, 7.0, 9.0, 10.0, 11.0, 12.0, 14.0}

	sortedArray := sorts.Bubble(unsortedArray)

	// Check if the sorted array matches the expected result
	if !reflect.DeepEqual(sortedArray, expectedSortedArray) {
		t.Errorf("BingoSort(%v) = %v; want %v", unsortedArray, sortedArray, expectedSortedArray)
	}
}

// BenchMark

func BenchmarkINTBubbleSort(b *testing.B) {
	// Generate a random unsorted array for each benchmark iteration
	unsortedArray := generateRandomArray(100) // Change the array size as needed

	for i := 0; i < b.N; i++ {
		// Make a copy of the unsorted array for each iteration to ensure fairness
		arrayCopy := make([]int, len(unsortedArray))
		copy(arrayCopy, unsortedArray)

		// Run the sorting algorithm and measure the time it takes
		b.StartTimer()
		sorts.Bubble(arrayCopy)

		b.StopTimer()
	}
}

func BenchmarkFLOATBubbleSort(b *testing.B) {
	// Generate a random unsorted array for each benchmark iteration
	unsortedArray := generateRandomFloatArray(100) // Change the array size as needed

	for i := 0; i < b.N; i++ {
		// Make a copy of the unsorted array for each iteration to ensure fairness
		arrayCopy := make([]float64, len(unsortedArray))
		copy(arrayCopy, unsortedArray)

		// Run the sorting algorithm and measure the time it takes
		b.StartTimer()
		sorts.Bubble(arrayCopy)

		b.StopTimer()
	}
}

// Recursive

func TestBubbleRecursiveIntSort(t *testing.T) {
	// Test case: Unsorted array
	unsortedArray := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	expectedSortedArray := []int{2, 3, 5, 6, 7, 9, 10, 11, 12, 14}

	sortedArray := sorts.RecursiveBubble(unsortedArray)

	// Check if the sorted array matches the expected result
	if !reflect.DeepEqual(sortedArray, expectedSortedArray) {
		t.Errorf("PigeonSort(%v) = %v; want %v", unsortedArray, sortedArray, expectedSortedArray)
	}
}

func TestRecursiveBubbleFloatSort(t *testing.T) {
	// Test case: Unsorted array
	unsortedArray := []float64{9.0, 7.0, 5.0, 11.0, 12.0, 2.0, 14.0, 3.0, 10.0, 6.0}
	expectedSortedArray := []float64{2.0, 3.0, 5.0, 6.0, 7.0, 9.0, 10.0, 11.0, 12.0, 14.0}

	sortedArray := sorts.RecursiveBubble(unsortedArray)

	// Check if the sorted array matches the expected result
	if !reflect.DeepEqual(sortedArray, expectedSortedArray) {
		t.Errorf("BingoSort(%v) = %v; want %v", unsortedArray, sortedArray, expectedSortedArray)
	}
}

// BenchMark

func BenchmarkINTRecursiveBubbleSort(b *testing.B) {
	// Generate a random unsorted array for each benchmark iteration
	unsortedArray := generateRandomArray(100) // Change the array size as needed

	for i := 0; i < b.N; i++ {
		// Make a copy of the unsorted array for each iteration to ensure fairness
		arrayCopy := make([]int, len(unsortedArray))
		copy(arrayCopy, unsortedArray)

		// Run the sorting algorithm and measure the time it takes
		b.StartTimer()
		sorts.RecursiveBubble(arrayCopy)

		b.StopTimer()
	}
}

func BenchmarkFLOATRecursiveBubbleSort(b *testing.B) {
	// Generate a random unsorted array for each benchmark iteration
	unsortedArray := generateRandomFloatArray(100) // Change the array size as needed

	for i := 0; i < b.N; i++ {
		// Make a copy of the unsorted array for each iteration to ensure fairness
		arrayCopy := make([]float64, len(unsortedArray))
		copy(arrayCopy, unsortedArray)

		// Run the sorting algorithm and measure the time it takes
		b.StartTimer()
		sorts.RecursiveBubble(arrayCopy)

		b.StopTimer()
	}
}
