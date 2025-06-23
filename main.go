package main

import "fmt"

// import "fmt"

// func main() {
// 	var workArray [10]uint8

// 	for i := 0; i < 10; i++ {
// 		fmt.Scan(&workArray[i])
// 	}

// 	for k := 0; k < 3; k++ {
// 		var i, j uint8
// 		fmt.Scan(&i, &j)
// 		temp := workArray[i]
// 		workArray[i] = workArray[j]
// 		workArray[j] = temp
// 	}

// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%d ", workArray[i])
// 	}

// }

func main() {
	baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Базовый массив: %v\n", baseArray)

	baseSlice := baseArray[5:8]
	// fmt.Printf(
	// 	"Срез, основанный на базовом массиве длиной %d и емкостью %d: %v\n",
	// 	len(baseSlice),
	// 	cap(baseSlice),
	// 	baseSlice,
	// )
	pointer := fmt.Sprintf("%p", baseSlice)

	baseSlice = append(baseSlice, 10)
	fmt.Printf("Массив: %v\n", baseArray)
	fmt.Printf("Срез длиной %d и емкостью %d: %v\n", len(baseSlice), cap(baseSlice), baseSlice)
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))
}
