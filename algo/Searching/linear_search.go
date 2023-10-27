package search

import (
	"fmt"
	"sync"
)

func Linear[T Numeric](array []T, target T) T {
	for i, value := range array {
		if value == target {
			fmt.Printf("Found %v at index %v\n", target, i)
			return T(rune(i))
		}
	}
	fmt.Printf("Element %v not found in the list\n", target)
	return T(rune(0))
}

func RecursiveLinear[T Numeric](arr []T, target T) T {
	var search func(int) int
	search = func(index int) int {
		if index >= len(arr) {
			fmt.Printf("%v not found in the list\n", target)
			return -1
		}

		if arr[index] == target {
			fmt.Printf("Found %v at index %v\n", target, index)
			return index
		}

		return search(index + 1)
	}

	return T(rune(search(0)))
}

// Multi Threading Linear Search

func LinearSearch[T Numeric](arr []T, target T) T {

	numThreads := 4 // Number of goroutines (adjust as needed)

	arrLen := len(arr)
	wg := sync.WaitGroup{}
	resultChan := make(chan int)

	// Divide the array into segments for each goroutine
	segmentSize := arrLen / numThreads
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		startIdx := i * segmentSize
		endIdx := (i + 1) * segmentSize
		if i == numThreads-1 {
			endIdx = arrLen // Handle the last segment
		}
		go func(arr []T, target T, start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				if arr[i] == target {
					resultChan <- i
					return
				}
			}
		}(arr, target, startIdx, endIdx)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect the results from goroutines
	foundIndex := -1
	for idx := range resultChan {
		if idx != -1 {
			foundIndex = idx
			break
		}
	}

	if foundIndex != -1 {
		fmt.Printf("Element %v found at index %v\n", target, foundIndex)
	} else {
		fmt.Printf("Element %v not found in the list.\n", target)
	}
	return T(rune(0))
}
