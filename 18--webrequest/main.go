package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response type of  %T \n ", response) // *http.Response

	defer response.Body.Close() // close the connection

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// html template
	fmt.Println("body is : ", string(data))
}
