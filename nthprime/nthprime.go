package nthprime

import "math"

func isPrime(num int) bool {

	for i:=2; i<=int(math.Sqrt(float64(num))); i++{
		if (num%i == 0){
			return false
		}
	}
	return true
}


func nthPrime(number int) int {
	n := 1
	primeCandidate := 3
	for   {
		if  isPrime(primeCandidate){
			n += 1
		}

		if n == number {
			return primeCandidate
		}
		primeCandidate += 2
	}
}
