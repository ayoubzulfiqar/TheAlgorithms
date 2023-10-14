package main

import (
	"fmt"
	"time"

	"github.com/ayoubzulfiqar/TheAlgorithms/algo/sort"
)

type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func main() {
	ms, us := MeasureExecutionTime()
	fmt.Printf("Execution time in milliseconds: %d ms\n", ms)
	fmt.Printf("Execution time in microseconds: %d µs\n", us)
}

// Function to measure execution time
func MeasureExecutionTime() (time.Duration, time.Duration) {
	startTime := time.Now()

	// Call the function you want to measure
	// arr := util.GenerateInt(1000000)
	arr := []int{1, 3, 5, 7, 8, 0, 2, 4, 6, 9}
	fmt.Printf("Unsorted %v\n", arr)

	array := sort.ThreeWayQuick(arr)
	fmt.Printf("Sorted %v\n", array)

	endTime := time.Now()

	// Calculate the duration in milliseconds, microseconds, and minutes
	duration := endTime.Sub(startTime)
	durationMilliseconds := duration / time.Millisecond
	durationMicroseconds := duration / time.Microsecond

	return durationMilliseconds, durationMicroseconds
}
