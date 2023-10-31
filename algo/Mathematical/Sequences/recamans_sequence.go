package sequence

import "fmt"

// Implement Using HashSet - Not available in go YET But SOON
func RecamanSequence(number int) []int {
	var sequence []int = make([]int, number)
	sequence[0] = 0
	for i := 1; i < number; i++ {
		var current int = sequence[i-1] - i
		for j := 0; j < i; j++ {
			if sequence[j] == current || current < 0 {
				current = sequence[i-1] + i
				break
			}
		}
		sequence[i] = current
	}
	fmt.Printf("Number: %v\nRecaman Sequence: %v\n", number, sequence)
	return sequence
}
