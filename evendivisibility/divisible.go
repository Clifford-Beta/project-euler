package evendivisibility

//GCD computes the Greatest Common Divisor of two integers using Euclid's algorithm.
func GCD(a, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	}
	if a < b {
		a, b = b, a
	}
	return GCD(b, a%b)
}

//LCM computes the Least Common Multiple of two number.
func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

func computeLCM(nums ...int) int {
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		lcm = LCM(lcm, nums[i])
	}
	return lcm

}
