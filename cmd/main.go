package main

import (
	"fmt"

	"github.com/ayoubzulfiqar/TheAlgorithms/algo/sort"
)

type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func main() {
	// data := util.GenerateRandomData(10)
	data := []int{2, 4, 6, 8, 9, 0, 1, 7, 3, 5}

	arr := sort.Sleep(data)

	fmt.Printf("Sorted: %v", arr)

}
