package main

import (
	"encoding/json"
	"fmt"
)

// lowercase mean private access
type course struct {
	Name     string `json:"coursename"` // this mean Name paramenter convert to coursename in json
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // remove this field when convert to json
	Tag      []string `json:"tags,omitempty"` // mitempty mean if this field is null (nill) don't throw that field
}

func main() {
	DecodeJson()
}

func EncodeJson() {

	lcoCourse := []course{
		{"Reactjs", 299, "dev@dev.com", "pass", []string{"web", "dev", "js"}},
		{"Java Spring", 211, "java@dev.com", "pass1", []string{"web", "java"}},
		{"Python", 399, "python@dev.com", "pass2", []string{"web", "app", "python"}},
		{"Nill", 500, "python@dev.com", "pass2", nil},
	}

	//package this data as json

	// fJson, err := json.Marshal(lcoCourse)
	fJson, err := json.MarshalIndent(lcoCourse, "lco", "\t") // lco is prefix

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", fJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(
		`{
			        "coursename": "Reactjs",
			             "Price": 299,
		             "website": "dev@dev.com",
			            "tags": [
		                     "web",
			                     "dev",
			                     "js"]
			     }`)
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Print("JSON was valid\n")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) // pass reference
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	var myOnlineData map[string]interface{} // maybe data type is not string, using interface keyword
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
