package functionAsValue

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 3, 5, 6}

	//**** Function is a fist class value in go
	// so we can pass a function in another function as a value *****

	doubleNum := transformNumbers(&numbers, double)
	tripleNum := transformNumbers(&numbers, triple)

	fmt.Println(doubleNum)
	fmt.Println(tripleNum)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}
