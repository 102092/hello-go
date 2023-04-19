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
	EncodeJson()
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
