package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// create dice number from random
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
		fallthrough
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("except 1 or 2")
	}

}
