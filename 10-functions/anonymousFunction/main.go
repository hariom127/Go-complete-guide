package anonymousFunction

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// pass anonymous function to transformNumbers function
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)

	double := createTransformer(2)
	triple := createTransformer(3)
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// **** Closure function ****
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
