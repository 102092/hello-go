package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on pointers")

	// * -> pointer, storing integer values
	// var ptr *int

	// fmt.Println("Value of pointer is .", ptr)

	// init
	myNumber := 23

	// creating pointer and reference some memory (myNumber)
	// reference -> &
	var ptr = &myNumber

	fmt.Println("Value of pointer is ", ptr)
	// Value of pointer is . 0xc00001e0e8

	fmt.Println("Value of pointer is ", *ptr)
	// Value of pointer is . 23

	*ptr = *ptr + 2

	fmt.Println("New value of pointer is ", myNumber)
	// New value of pointer is  25
}
