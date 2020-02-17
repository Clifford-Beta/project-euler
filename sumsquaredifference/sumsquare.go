package sumsquaredifference

func sumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += (i * i)
	}
	return sum
}

func squareOfSum(n int) int {
	square := 0
	for i := 1; i <= n; i++ {
		square += i
	}
	return square * square

}

func sumSquareDifference(n int) int {
	return squareOfSum(n) - sumOfSquares(n)
}
