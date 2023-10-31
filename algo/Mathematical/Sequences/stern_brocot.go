package sequence

/*


TODO: Needs Better Implementation


The Stern-Brocot sequence is a sequence of rational numbers that can be generated using the following steps:

1. Start with the two numbers 1 and 1.
2. To generate the next number in the sequence, take the average of the two previous numbers and then reduce to lowest terms.
3. Repeat step 2.

Here is a step-by-step example of how to generate the Stern-Brocot sequence without code:

```
Start with: 1, 1

Next number: (1 + 1) / 2 = 1
Next number: (1 + 1) / 2 = 1
Next number: (1 + 2) / 2 = 3/2
Next number: (3/2 + 2) / 2 = 5/4
Next number: (5/4 + 3/2) / 2 = 13/8
Next number: (13/8 + 5/4) / 2 = 7/5
```

And so on.

The Stern-Brocot sequence has many interesting properties, including:

* It is a fractal, meaning that it contains smaller copies of itself.
* It is a dense sequence, meaning that every rational number can be approximated by a number in the sequence.
* It has applications in music theory, computer science, and other fields.

Table: Stern-Brocot sequence:

| Term | Fraction |
|---|---|
| 1 | 1/1 |
| 2 | 2/1 |
| 3 | 1/2 |
| 4 | 3/2 |
| 5 | 5/4 |
| 6 | 13/8 |
| 7 | 7/5 |
| 8 | 17/11 |
| 9 | 41/26 |
| 10 | 23/15 |

// Why Return Type is []string not []int:

The Stern-Brocot Sequence is a sequence of fractions that cannot always be accurately represented as integers.T
he reason for returning []string instead of []int is that each element in the sequence is a fraction
in the form "numerator/denominator," and the fractions in this sequence do not necessarily have to have integer
values for their numerators and denominators.

For example: in the Stern-Brocot Sequence, you will find fractions like 1/2, 2/3, 3/4, and so on.
These fractions cannot be accurately represented as integers because they have non-integer values
for both the numerator and the denominator.

By returning []string, we can accurately represent and preserve the fractions in their exact form,
avoiding rounding errors that would occur if we were to convert them to integers.
This ensures that the sequence maintains its precision and accuracy.

*/
/*
func SternBrocot(number int) []string {
	// Root NODE
	var sequence []string = []string{"1/1"}
	for i := 1; i < number; i++ {
		newSeq := []string{}
		for j := 0; j < len(sequence); j++ {
			fraction := sequence[j]
			sep := strings.Index(fraction, "/")
			numerator, _ := strconv.Atoi(fraction[:sep])
			denominator, _ := strconv.Atoi(fraction[sep+1:])

			leftNumerator := numerator
			leftDenominator := numerator + denominator

			rightNumerator := numerator + denominator
			rightDenominator := denominator

			newSeq = append(newSeq, fmt.Sprintf("%d/%d", leftNumerator, leftDenominator))
			newSeq = append(newSeq, fmt.Sprintf("%d/%d", rightNumerator, rightDenominator))
		}

		sequence = newSeq

	}
	for idx, value := range sequence {

		fmt.Printf("Number: %v\nStern-Brocot Sequence:%v", idx, value)
	}
	return sequence
}
*/
func SternBrocot(n int) []float64 {
	if n <= 0 {
		return []float64{}
	}

	result := make([]float64, 0, n)
	computeSternBrocot(&result, 0, 1, 1, 0, n)
	return result
}

func computeSternBrocot(result *[]float64, a, b, c, d, n int) {
	if len(*result) >= n {
		return
	}

	// Compute the next term in the sequence
	x := float64(a+c) / float64(b+d)
	*result = append(*result, x)

	// Recursively compute the left and right branches
	computeSternBrocot(result, a, b, a+c, b+d, n)
	computeSternBrocot(result, a+c, b+d, c, d, n)
}
