package main

import "fmt"

func main() {

	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	// adding value
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println("JS shorts for :", languages["JS"])

	// delete value
	delete(languages, "RB")
	fmt.Println(languages)

	// how to loop through map?
	for key, value := range languages {
		fmt.Printf("For key %v, Value is %v\n", key, value)
	}

	// _ mean anything..
	for _, value := range languages {
		fmt.Printf("For key v, Value is %v\n", value)
	}
}
