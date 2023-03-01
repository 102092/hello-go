package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println("present time: ", presentTime)

	// 특정 포맷으로 알아서 바꿔준다.
	fmt.Println("present time: ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.April, 1, 12, 23, 30, 11, time.UTC)
	fmt.Println("present time: ", createdDate)
	fmt.Println("present time: ", createdDate.Format("01-02-2006 Monday"))
}
