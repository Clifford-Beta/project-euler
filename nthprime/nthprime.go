package nthprime

func isPrime(num int) bool {

	for i:=2; i<num; i++{
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
			//fmt.Println("prime candidate ", primeCandidate, isPrime(primeCandidate), n)
		}

		if n == number {
			return primeCandidate
		}
		primeCandidate += 2
	}
}
