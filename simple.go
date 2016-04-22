package main

import "fmt"

// main runs the program.
func main() {
	myfunc(2)
	fmt.Println("Hello world")
}

// myfunc is a function that does nothing.
func myfunc(y int) {
	x := y
	fmt.Println(x)
}
