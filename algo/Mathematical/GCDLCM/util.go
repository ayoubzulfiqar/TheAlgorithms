package math

// Generic
type Integer interface {
	~int | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// GCD: Greatest Common Divisor - Using Euclidean algorithm
func getGDC[I Integer](a I, b I) I {

	for a > 0 && b > 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	if a == 0 {
		return b
	}
	return a
}

// LCM: Least Common Multiple
// Formula: LCM(a, b) = (a * b) / GCD(a, b)

func getLCM[I Integer](a I, b I) I {
	return (a * b) / getGDC(a, b)
}

// GCD for Float
func FGcd(a, b float64) float64 {
	if a < b {
		return FGcd(b, a)
	}

	// base case
	if abs(b) < 0.001 {
		return a
	}

	return FGcd(b, a-a/b*b)
}
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
