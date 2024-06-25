package main

import "fmt"

func main() {
	// Первый способ определения массивов в Go
	var numbersOne [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbersOne)

	// Второй способ определения массива
	numbersTwo := [5]int{6, 7, 8, 9, 9}
	fmt.Println(numbersTwo)

	example := [3]int{1, 2, 3}
	fmt.Println(example[0])

	// В Go индексы выступают как ключи к значениям. При этому числовые ключи необязательно располагать в порядке возрастания
	colors := [3]string{0: "red", 1: "blue", 2: "green"}
	fmt.Println(colors[0])
}
