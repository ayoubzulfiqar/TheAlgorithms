package math

/*



*/


func FindGCD[I Integer](arr []I) I {
	var gcd I = arr[0]
	for i := 1; i < len(arr); i++ {
		gcd = getGDC(gcd, arr[i])
	}
	return gcd
}
