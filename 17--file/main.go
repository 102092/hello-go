package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("go files")
	content := "test text"

	file, err := os.Create("./mylcogofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	//  databyte [116 101 115 116 32 116 101 120 116] -> byte code...
	fmt.Println("Text data inside in file \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		// panic means shutdown program and throw error
		panic(err)
	}
}
