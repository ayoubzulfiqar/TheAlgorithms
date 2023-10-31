package math

/*

The Euclidean algorithm is a way to find the greatest common divisor of two positive integers.

*/

func Euclidean[I Integer](a I, b I) I {
	var x I = 0
	var y I = 1

	var lastX I = 0
	var lastY I = 1

	for b != 0 {
		quotient := a / b
		a, b = b, a%b
		lastX, x = x, lastX-quotient*x
		lastY, y = y, lastY-quotient*y
	}

	return a
}
