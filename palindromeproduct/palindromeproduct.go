package palindromeproduct

import (
	"strconv"
)

func findPalindromeProduct() int {
	max := 0
	maxi := 0
	maxj := 0
	for i := 1000; i >= 100; i-- {
		for j := 1000; j >= 100; j-- {
			product := i * j

			if isPalindrome(strconv.Itoa(product)) {
				if product > max {
					max = product
					maxi = i
					maxj = j
				} else if i < maxi && j < maxj {
					//this little optimization reduced the number of iterations
					//required to find the solution from 812702 to 566338
					break
				}
			}
		}
	}
	return max
}

func isPalindrome(str string) bool {
	return str == reverse(str)
}

func reverse(str string) string {
	reversed := ""
	for _, v := range str {
		reversed = string(v) + reversed
	}
	return reversed
}
