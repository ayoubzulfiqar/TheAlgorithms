package main

import (
	"fmt"

	sorts "github.com/ayoubzulfiqar/TheAlgorithms/algo/sort"
)

type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func main() {
	arr := []int{1, 3, 2, 5, 6, 4,7, 8, 10, 9}
	array := sorts.CockTail(arr)
	fmt.Printf("Sorted: %v", array)

}
