// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

/* block 3.3 task 1
func main() {
	var a int
	var b int
	fmt.Println("Only numbers!")
	fmt.Scan(&a)
	fmt.Scan(&b)
	summ := summNumbers(a, b)
	fmt.Println("Result:", summ)

}
func summNumbers(a, b int) int {

	result := a + b
	return result

}
*/

/*  block 3.3 task 2
func main() {
	var a int
	var b int
	var c int
	fmt.Println("Only numbers!")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	maxNum := maxNum(a, b, c)
	fmt.Println("Result:", maxNum)

}

func maxNum(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else {
		return c
	}
}
*/

// block 3.3 task 3
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	fmt.Println("Only string!")

	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	s = strings.TrimSpace(s) // убираем \n

	result := LenStrings(s)
	fmt.Println("Result:", result)
}

func LenStrings(s string) int {
	return len(s)
}

/* 	block 3.4 task 1

   	var char1 int

   	for char1 = 1; char1 <= 50; char1++ {

   		if char1%3 == 0 && char1%5 == 0 {
   			fmt.Println("FizzBuzz")
   		} else if char1%3 == 0 {
   			fmt.Println("Fizz")
   		} else if char1%5 == 0 {
   			fmt.Println("Buzz")
   		} else {
   			fmt.Println(char1)
   		}

   	}
*/

/*	block 3.4 task 2

	var char1 int

	summ := 0
	for char1 = 1; char1 <= 100; char1++ {
		summ += char1
	}
	fmt.Println(summ)
*/

/*	block 3.4 task 3

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

*/
