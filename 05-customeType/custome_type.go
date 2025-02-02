package main

import "fmt"

type str string

// add method to this str custom type

func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str = "Hariom"
	name.log()
}
