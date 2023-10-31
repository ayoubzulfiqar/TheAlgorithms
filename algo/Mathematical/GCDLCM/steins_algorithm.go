package math

// Algorithm to find Greatest Common Divisor using Stein's algorithm
func Steins(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	// k is the greatest
	var k int
	for k = 0; ((a | b) & 1) == 0; k++ {
		a >>= 1
		b >>= 1
	}
	for b != 0 {
		for (b & 1) == 0 {
			b >>= 1
		}
		if a > b {
			a, b = b, a
		}
		b = (b - a)
	}
	return a << k
}
