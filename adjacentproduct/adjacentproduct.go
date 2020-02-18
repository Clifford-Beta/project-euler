package adjacentproduct

import (
	"strconv"
)

func adjacentNDigits(number string, n int) ([]int, int) {
	maxProduct := 0
	digits := []int{}
	noOfDigits := len(number)
	for i := 0; i < noOfDigits; i++ {
		j := i + n
		if j >= noOfDigits {
			break
		}
		dgts, product := calculateProduct(number[i:j])
		if product > maxProduct {
			digits, maxProduct = dgts, product
		}
		//calculate product
	}
	return digits, maxProduct
}

func calculateProduct(num string) ([]int, int) {
	numbers := []int{}
	product := 1
	for _, rune := range num {
		val, _ := strconv.Atoi(string(rune))
		product *= val
		numbers = append(numbers, val)
	}
	return numbers, product
}
