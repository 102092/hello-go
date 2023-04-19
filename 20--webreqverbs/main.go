package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// PerformGetRequest()
	PerFromPostJsonRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	// have to close connection
	defer response.Body.Close()

	fmt.Println("Status code is : ", response.StatusCode)
	fmt.Println("Content length is  : ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(content) // byte format...
	fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("byte count is ", byteCount) // exactly same content length
	fmt.Println(responseString.String())
}

func PerFromPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`{
		"name" : "hello",
		"age" : 123
	}`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
