package main

import (
	"fmt"

	"github.com/ayoubzulfiqar/TheAlgorithms/algo/sort"
)

type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("Unsorted array:", arr)

	array := sort.RecursiveInsertion(arr)

	fmt.Println("Sorted array:", array)
}
