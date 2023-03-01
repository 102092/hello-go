package main

import "fmt"

func main() {

	var frutiList [4]string

	frutiList[0] = "Apple"
	frutiList[1] = "Tomato"
	frutiList[3] = "Peach"

	fmt.Println("Fruit list is : ", frutiList)
	// Fruit list is :  [Apple Tomato  Peach]
	// fruitList[2] 가 비어있어서, Tomato, Peach 사이의 스페이스가 좀 크다

	fmt.Println("Fruit list is : ", len(frutiList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("vegy list is ", vegList)
}
