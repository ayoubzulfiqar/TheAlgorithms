package main

import (
	"fmt"

	sorts "github.com/ayoubzulfiqar/TheAlgorithms/algo/sort"
)

func main() {

	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	bingedArray := sorts.Bingo(arr)

	fmt.Printf("Sorted Array: %v\n", bingedArray)

}
