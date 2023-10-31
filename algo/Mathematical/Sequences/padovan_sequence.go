package sequence

/*

 P(n) = P(n-2) + P(n-3)
P(0) = P(1) = P(2) = 1

*/

func Padovan[T Integer](number T) T {
	var padovan []T = make([]T, number)

	padovan[0] = 1
	padovan[1] = 1
	padovan[2] = 1
	for i := 3; i < len(padovan); i++ {
		padovan[i] = padovan[number-2] + padovan[number-3]
	}

	return padovan[number-1]
}
