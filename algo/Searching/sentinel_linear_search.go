package search

import (
	"fmt"
	"sync"
)

func Sentinel[T Numeric](arr []T, key T) T {
	var n int = len(arr)
	var last T = arr[n-1]
	arr[n-1] = key
	var idx int = 0

	for arr[idx] != key {
		idx++
	}

	arr[n-1] = last

	if (idx < n-1) || (arr[n-1] == key) {
		fmt.Printf("Found: %v present at Index: %v\n", key, idx)
	} else {
		fmt.Println("Element not found")
	}
	return T(rune(0))
}

func RecursiveSentinel[T Numeric](arr []T, key T) T {
	var n int = len(arr)
	if n <= 0 {
		return T(rune(0))
	}

	if arr[n-1] == key {
		fmt.Printf("Found: %v present at Index:%v\n", key, n-1)
		return T(rune(n - 1))
	} else {
		fmt.Printf("Not Found: %v\n", key)

	}
	return RecursiveSentinel(arr[n:], key)
}

// Multi Threading Sentinel Linear Search

func SentinelLinearSearch[T Numeric](arr []T, key T) T {
	n := len(arr)
	numThreads := 4 // You can adjust the number of threads based on your requirements.
	step := n / numThreads
	resultChan := make(chan T)
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go SentinelSearch(arr, key, i*step, (i+1)*step, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		return result
	}
	fmt.Printf("Not found %v\n", key)
	return T(rune(0))

}

func SentinelSearch[T Numeric](arr []T, key T, start int, end int, wg *sync.WaitGroup, resultChan chan T) {
	defer wg.Done()
	for i := start; i < end; i++ {
		if arr[i] == key {
			resultChan <- T(rune(i))
			fmt.Printf("Found: %v present at Index:%v\n", key, i)
			return
		}
	}
}
