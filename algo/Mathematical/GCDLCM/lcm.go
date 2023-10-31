package math

func FindLCM[I Integer](arr []I) I {
	var lcm I = arr[0]
	for i := 1; i < len(arr); i++ {
		lcm = getLCM(lcm, arr[i])
	}
	return lcm
}

func LCM[I Integer](arr []I) I {
	var lcm I = 1
	var divisor I = 2

	for {
		var counter int = 0
		divisible := false

		for i := 0; i < len(arr); i++ {
			if arr[i] == 0 {
				return 0
			} else if arr[i] < 0 {
				arr[i] = arr[i] * 1
			}
			if arr[i] == 1 {
				counter++
			}
			if arr[i]%divisor == 0 {
				divisible = true
				arr[i] = arr[i] / divisor
			}
		}

		if divisible {
			lcm = lcm * divisor
		} else {
			divisor++
		}

		if counter == len(arr) {
			return lcm
		}
	}
}
