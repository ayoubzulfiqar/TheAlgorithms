package sequence

import "fmt"

/*

P(n) = P(P(n - 1)) + P(n - P(n - 1))

*/

func NewmanConway(number int) int {
	var sequence []int = make([]int, number+1)
	sequence[0] = 0
	sequence[1] = 1
	sequence[2] = 1
	for i := 3; i <= number; i++ {
		sequence[i] = sequence[sequence[i-1]] + sequence[i-sequence[i-1]]
	}
	fmt.Printf("Number: %v\nNewmanConway Sequence: %v\nResult: %v\n", number, sequence, sequence[number])
	return sequence[number]
}
