package division

import "fmt"

// Generic Version to Work With

type Integer interface {
	~int | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func IsDivisible[I Integer](number I, by I) bool {
	if number%by == 0 {
		fmt.Printf("Number: %v Divided by %v is %v\nYES: It's Divisible\n", number, by, number/by)
		return true
	}
	fmt.Printf("Number: %v Can Not be Divided by %v\n", number, by)
	return false
}
