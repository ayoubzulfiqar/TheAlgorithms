package sequence

import (
	"fmt"
	"math"
)

func SeriesSum(number int) float64 {
	var seriesSum float64 = 0.02469 * ((10*math.Pow(10, float64(number)) - 1) - (9 * float64(number)))
	fmt.Printf("Number: %v\nSeriesSum: %v\n", number, seriesSum)
	return seriesSum
}
