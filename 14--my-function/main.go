package main

import "fmt"

func main() {
	greeter()

	// result := adder(3, 5)
	result, hello := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1)

	fmt.Println("Result is :", result)
	fmt.Println("my message is :", hello)

}

func adder(i1 int, i2 int) int {
	return i1 + i2
}

// all value type are integer
func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "hello"
}

func greeter() {
	fmt.Println("Namstey from golang")
}
