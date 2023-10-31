package sequence

import (
	"fmt"
	"math"
)

func Juggler[T Integer](number T) []T {
	var sequence []T
	sequence = append(sequence, number)

	for number != 1 {
		var next T
		if int(number)%2 == 0 {
			next = T(math.Sqrt(float64(number)))
		} else {
			next = T(math.Sqrt(float64(number)) * math.Sqrt(float64(number)) * math.Sqrt(float64(number)))
		}
		sequence = append(sequence, next)
		number = next
	}
	fmt.Printf("Juggler Sequence for %v: %v\n", number, sequence)
	return sequence
}
