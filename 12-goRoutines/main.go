package main

import (
	"fmt"
	"time"
)

func main() {
	// isDone := make([]chan bool, 3)
	isDone := make(chan bool)
	// isDone[0] = make(chan bool)
	go greet("Hello from first greet.", isDone)

	// isDone[1] = make(chan bool)
	go slowGreet("Hello..... from... slow....", isDone)

	// isDone[2] = make(chan bool)
	go greet("Hello from second greet.", isDone)

	// <-isDone
	// <-isDone
	// <-isDone

	for range isDone {
		<-isDone
	}
}

func greet(phrase string, isDone chan bool) {
	fmt.Println(phrase)
	isDone <- true
}

func slowGreet(phrase string, isDone chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)
	isDone <- true
	close(isDone)
}
