package sequence

import "fmt"



func SylvesterSequence(number int) []int {

	var product int = 1
	var ans int = 2
	// N := 1000000007
	// max int Limit
	N := 9223372036854775807
	sequence := make([]int, number)

	for i := 0; i < number; i++ {
		sequence[i] = ans
		ans = ((product % N) * (ans % N)) % N
		product = ans
		ans = (ans + 1) % N
	}
	fmt.Printf("Number: %v\nSylvester Sequence: %v\nProduct: %v\n", number, sequence, product)
	return sequence
}
