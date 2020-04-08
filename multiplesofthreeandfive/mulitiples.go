package multiples

import "fmt"

func findSumOfAllMultiplesNumbers(numbers []int, upperLimit int) int {

	sum := 0
	for i := 0; i < upperLimit; i++ {
		//previouslyAdded := 0
		for _, v := range numbers {
			if i%v == 0 {
				sum = sum + i
				break //break out of the inner loop as an early exit strategy, this simulates and OR condition.

			}
		}
	}

	return sum
}

func main() {
	upperLimit := 1000
	numbers := []int{3, 5}
	fmt.Println(findSumOfAllMultiplesNumbers(numbers, upperLimit))

}
