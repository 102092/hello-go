package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	// [Apple Tomato Peach Mango]

	// colon syntax mean slice
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
	// [Tomato Peach Mango Banana]

	// start position 1 stop position in 3
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	// [Peach Mango]

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 456
	highScore[3] = 789
	// highScore[4] = 777 // panic: runtime error: index out of range [4] with length 4

	// re-allocate memory (highScore []int) and new value accomodate
	highScore = append(highScore, 777) // no error and print [234 945 456 789 777]

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)                     // [234 456 777 789 945]
	fmt.Println(sort.IntsAreSorted(highScore)) // true
}
