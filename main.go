package main

import "fmt"

func main() {

	// 1 задача

	// var char1 int

	// for char1 = 1; char1 <= 50; char1++ {

	// 	if char1%3 == 0 && char1%5 == 0 {
	// 		fmt.Println("FizzBuzz")
	// 	} else if char1%3 == 0 {
	// 		fmt.Println("Fizz")
	// 	} else if char1%5 == 0 {
	// 		fmt.Println("Buzz")
	// 	} else {
	// 		fmt.Println(char1)
	// 	}

	// }

	// 2 задача

	var char1 int

	summ := 0
	for char1 = 1; char1 <= 100; char1++ {
		summ += char1
	}
	fmt.Println(summ)

}
