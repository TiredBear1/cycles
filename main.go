package main

import "fmt"

func main() {

	// task 1

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

	// task 2

	// var char1 int

	// summ := 0
	// for char1 = 1; char1 <= 100; char1++ {
	// 	summ += char1
	// }
	// fmt.Println(summ)

	// task 3

	var num1 int
	fmt.Println("Введите число дня недели (1-7):")
	fmt.Scan(&num1)

	switch num1 {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6:
		fmt.Println("Суббота")
	case 7:
		fmt.Println("Воскресенье")
	default:
		fmt.Println("Такой цифры дня недели нет!")
	}

}
