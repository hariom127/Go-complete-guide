package main

import "fmt"

func main() {

	factorial := getFactorial(3)

	fmt.Println(factorial)

}

func getFactorial(num int) int {

	if num == 0 {
		return 1
	}
	return num * getFactorial(num-1)

}
