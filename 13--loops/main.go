package main

import "fmt"

func main() {

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// fmt.Println(days[d])
	// }

	// 위 for 문과 결과는 같음
	// for i := range days {
	// fmt.Println(days[i])
	// }

	// comma ok syntax
	for index, day := range days {
		// index 0 and value Sunday
		// index 1 and value Tuesday
		// index 2 and value Wednesday
		// index 3 and value Friday
		// index 4 and value Saturday
		fmt.Printf("index %v and value %v\n", index, day)
	}

	rougueValue := 1

	// if rougueValue is 10, will stop
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		// as soon as match this value is 7, for loop wil stop
		if rougueValue == 7 {
			break
		}

		fmt.Println("value is : ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("hello rougueValue is 2")

}
