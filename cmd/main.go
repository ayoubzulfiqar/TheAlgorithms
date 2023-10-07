package main

import (
	"fmt"

	sorts "github.com/ayoubzulfiqar/TheAlgorithms/algo/sort"
)

type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func main() {
	arr := []float64{1.4, 3.4, 2.53, 0.5, 6.854, 4.3434,545.7, 8.455, 10/6, 9.35}
	array := sorts.OddEven(arr)
	fmt.Printf("Sorted: %v", array)

}
