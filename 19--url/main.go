package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=123"

func main() {
	fmt.Println(myurl)

	//parsing url..
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // lco.dev:3000
	fmt.Println(result.Port())   // 3000
	fmt.Println(result.Path)     // /learn
	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid=123

	qparam := result.Query()
	fmt.Printf("type of query param are %T\n", qparam) // url.Value
	fmt.Println(qparam["coursename"])                  // reactjs

	for _, val := range qparam {
		fmt.Println("val is ", val)
	}

	// passing reference
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
