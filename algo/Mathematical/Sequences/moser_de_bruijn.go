package sequence

import "fmt"

/*

1. S(2 * n) = 4 * S(n)
2. S(2 * n + 1) = 4 * S(n) + 1
with S(0) = 0 and S(1) = 1



Input : 5
Output : 0 1 4 5 16

Input : 10
Output : 0 1 4 5 16 17 20 21 64 65
*/

func MoserDeBruijn[T Integer](number T) []T {
	var sequence []T = make([]T, number+1)
	sequence[0] = 0
	if number != 0 {
		sequence[1] = 1
	}
	for i := 2; i <= int(number); i++ {
		// S(2 * n) = 4 * S(n)
		if i%2 == 0 {
			sequence[i] = 4 * sequence[i/2]
		} else {
			// S(2 * n + 1) = 4 * S(n) + 1
			sequence[i] = 4*sequence[i/2] + 1

		}
	}
	fmt.Printf("Number: %v\nMoser Sequence: %v\n", number, sequence)
	return sequence
}
