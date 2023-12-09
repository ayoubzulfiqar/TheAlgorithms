package digit

import (
	"fmt"
	"math"
)

/*
Minimum Number of digits should be remove
to make a number perfect square

*/
func PerfectSquare(number string) int {
	var n int = len(number)
	var ans int = -1
	var num string

	for i := 1; i < (1 << uint(n)); i++ {
		var str string = ""

		for j := 0; j < n; j++ {
			if (i>>uint(j))&1 == 1 {
				str += string(number[j])
			}
		}

		if str[0] != '0' {
			var temp int = 0
			for j := 0; j < len(str); j++ {
				temp = temp*10 + int(str[j]-'0')
			}

			var k int = int(math.Sqrt(float64(temp)))

			if k*k == temp {
				if ans < len(str) {
					ans = len(str)
					num = str
				}
			}
		}
	}

	if ans == -1 {
		return ans
	}

	fmt.Println("Perfect Square:", num+" ")
	return n - ans
}

func IsPerfectSquare(num int) bool {
	root := int(math.Sqrt(float64(num)))
	return root*root == num
}
