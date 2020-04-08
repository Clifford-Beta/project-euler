package evenfibonnaci

func Fibonacci(n uint) uint64 {

	sequence := []uint64{}
	if n <= 1 {
		return uint64(n)
	}

	var n2, n1 uint64 = 0, 1
	sequence = append(sequence, n2)

	for i := uint(2); i < n; i++ {
		sequence = append(sequence, n1)
		n2, n1 = n1, n1+n2
	}

	return n2 + n1
}

func FibonacciSequence(n uint) []uint64 {

	sequence := []uint64{}
	if n <= 1 {
		return sequence
	}

	var n2, n1 uint64 = 0, 1
	sequence = append(sequence, n2)

	for i := uint(2); i < n; i++ {
		sequence = append(sequence, n1)
		n2, n1 = n1, n1+n2
	}

	return sequence
}

func sumEvenTermms(terms []uint64) uint64 {
	sum := uint64(0)
	for _, v := range terms {
		if v%2 == 0 && v < 4000000 {
			sum = sum + v
		} else if v >= 4000000 {
			break
		}

	}
	return sum
}
