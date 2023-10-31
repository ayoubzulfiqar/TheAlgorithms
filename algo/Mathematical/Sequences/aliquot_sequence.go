package sequence

import "fmt"

/*

Input:  n = 10
Output: 10 8 7 1 0
Sum of proper divisors of 10 is  5 + 2 + 1 = 8.
Sum of proper divisors of 8 is 4 + 2 + 1 = 7.
Sum of proper divisors of 7 is 1
Sum of proper divisors of 1 is 0
Note that there is no proper divisor of 1.

Input  : n = 6
Output : 6
         Repeats with 6

Input : n = 12
Output : 12 16 15 9 4 3 1 0

*/

func divisorsSum(number int) int {
	var sum int = 1 // Start with 1 to account for 1 being a proper divisor
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			sum += i
			if number/i != i {
				sum += number / i
			}
		}
	}
	fmt.Printf("Divisor Sum: %v\n", number)
	return sum
}

func Aliquot[T Integer](number T) []T {
	var sequence []T = []T{number}
	// Create a map to store the sum of proper divisors for each number in the sequence
	var divisorSumMap map[T]T = make(map[T]T)
	divisorSumMap[number] = T(divisorsSum(int(number)))

	for {
		var next T = divisorSumMap[sequence[len(sequence)-1]]
		if next == number || next < 0 {
			break
		}

		// Check if we've encountered this number before to avoid infinite loops
		if _, exists := divisorSumMap[next]; exists {
			break
		}

		sequence = append(sequence, next)
		divisorSumMap[next] = T(divisorsSum(int(next)))
	}
	fmt.Printf("Number: %v\nAliquot Sequence: %v\n", number, sequence)
	return sequence
}
