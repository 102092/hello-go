package main

import "fmt"

func main() {

	hello := User{"hello", "hello@gmail", true, 100}
	fmt.Println(hello)
	// hello details are : {Name:hello Email:hello@gmail Status:true Age:100}
	// +v syntax -> key:value ...
	fmt.Printf("hello details are : %+v\n", hello)
	fmt.Printf("Name is %v and email is %v", hello.Name, hello.Email)
	hello.GetStatus()
	hello.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@dev.go"
	fmt.Println("User Email is : ", u.Email)
}
