package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
}

/*** Declare generic type in multiple way ***/
// add[T interface{}](a, b T) T {
// add[T any](a, b T) T {

func add[T int | float64 | string](a, b T) T {
	return a + b
}
