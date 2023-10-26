package main

import (
	"fmt"
	"time"

	"github.com/ayoubzulfiqar/TheAlgorithms/algo/search"
)

type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func main() {
	ms, us, s := MeasureExecutionTime()
	fmt.Printf("Execution time in milliseconds: %d ms\n", ms)
	fmt.Printf("Execution time in microseconds: %d µs\n", us)
	fmt.Printf("Execution time in seconds: %d s\n", s)

}

// Function to measure execution time
func MeasureExecutionTime() (time.Duration, time.Duration, time.Duration) {
	startTime := time.Now()

	// Call the function you want to measure
	// arr := util.GenerateInt(1000000)
	// arr := []int{1, 3, 5, 7, 8, 2, 4, 6, 9}
	// arr := []int{0, 1, 2, 3, 4, 5}
	// arr := []string{"A", "B", "C", "D"}
	// arr := []float64{9.0, 7.0, 5.0, 11.0, 12.0, 2.0, 14.0, 3.0, 10.0, 6.0}
	// arr := []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0}
	// arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	// key := 1
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	key := 0
	idx := search.Fibonacci(arr, key)
	if idx == -1 {
		fmt.Printf("NOt FOUND%v", key)

	} else {
		fmt.Printf("Found Key: %v at Index: %v", key, idx)
	}

	endTime := time.Now()

	// Calculate the duration in milliseconds, microseconds, and minutes
	duration := endTime.Sub(startTime)
	durationMilliseconds := duration / time.Millisecond
	durationMicroseconds := duration / time.Microsecond
	durationSeconds := duration / time.Second

	return durationMilliseconds, durationMicroseconds, durationSeconds
}

//

/*

func FastSqrt(x float64) float64 {
	if x == 0 {
		return 0
	}

	bits := math.Float64bits(x)
	bits -= 1 << 52 // Subtract 2^52 from the exponent.
	bits >>= 1      // Divide by 2.

	bits += 1 << 61 // Add 2^61 to the exponent.
	return math.Float64frombits(bits)
}

*/
