package main

import "fmt"

func main() {
	var name string = "Misha" // Задаем значение переменной 1-ым способом через "var (название переменной) (тип переменной)  = ..."
	fmt.Println(name)

	name = "Katya" // в Go существует возможность задавать новое значение переменной
	fmt.Println(name)

	secondName := "Danya" // Задаем значение переменной 2-ым споособом через "(название переменной) := ..."
	fmt.Println(secondName)

	secondName = "Kostya"
	fmt.Println(secondName)
}
