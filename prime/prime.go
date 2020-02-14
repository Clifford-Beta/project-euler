package prime

func findLargestPrimeFactor(prime int) int {
	for {
		if prime%2 > 0 {
			break
		} else {
			prime /= 2
		}
	}
	divisor := 3
	for {
		if prime%divisor == 0 {
			prime /= divisor
		} else {
			divisor += 2
		}
		if divisor == prime {
			break
		}
	}
	return prime
}
