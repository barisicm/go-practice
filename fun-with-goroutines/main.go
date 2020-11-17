package main

import (
	"fmt"
	"time"
)

//EXAMPLE 2
func main() {
	n := 3

	// This is where we "make" the channel, which can be used
	// to move the `int` datatype
	out := make(chan int)

	// We still run this function as a goroutine, but this time,
	// the channel that we made is also provided
	go multiplyByTwo(n, out)

	// Once any output is received on this channel, print it to
	// the console and proceed
	fmt.Println(<-out)
}

// This function now accepts a channel as its second argument...
func multiplyByTwo(num int, out chan<- int) {
	result := num * 2

	//... and pipes the result into it
	out <- result
}

//EXAMPLE 1!
func main1() {
	n := 3

	// We want to run a goroutine to multiply n by 2
	go multiplyByTwo1(n)

	// We pause the program seo that the `multiplyByTwo` goroutine
	// can finish and print the output before the code exits
	time.Sleep(time.Second)
}

func multiplyByTwo1(num int) int {
	result := num * 2
	fmt.Println(result)
	return result
}
