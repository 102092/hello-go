package main

import "fmt"

const LoginToken string = "hello" // Public

func main() {
	var username string = "han"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n ", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n ", isLoggedIn)

	// var smallVal uint8 = 256 // overflow
	var smallVal uint8 = 255 // overflow
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n ", smallVal)

	// var smallFloat float32 = 255.12391231411 // 255.12392, five value for decimal
	var smallFloat float64 = 255.12391231411 // 255.12391231411
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n ", smallFloat)

	// default value and some aliases
	var aVariable int // default value is 0
	fmt.Println(aVariable)
	fmt.Printf("Variable is of type: %T \n ", aVariable) // float64

	// implicit type
	var website = "website"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n ", website) // string

	// no var style
	numberOfUser := 300_000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n ", numberOfUser) // string

	// cons
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n ", LoginToken) // string
}
