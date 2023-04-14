package main

import "fmt"

func main() {
	// defer과 같이 선언된 코드는 바로 실행되지 않고, 마지막 brasket 전에 실행됨.
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("Hello")
	MyDefer() // function call first
}

func MyDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("value : ", i)
	}
}
