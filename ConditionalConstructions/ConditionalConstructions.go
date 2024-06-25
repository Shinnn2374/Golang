package main

import "fmt"

func main() {
	firstNum := 5
	secondNum := 5

	if firstNum < secondNum {
		fmt.Println("See again")
	} else if firstNum > secondNum {
		fmt.Println(":(")
	} else {
		fmt.Println("first number = second number")
	}

	/////////////////////////////////////////////////////

	x := 3

	switch x {
	case 1:
		fmt.Println("x = 1")
	case 2:
		fmt.Println("x = 2")
	case 3:
		fmt.Println("x = 3")
	default:
		fmt.Println("Privetik")
	}

	////////////////////////////////////////////////////

	mishaAge := 27
	vasyaAge := 25
	katyaAge := 21

	high := 1
	middle := 0
	low := -1

	if mishaAge > vasyaAge && mishaAge > katyaAge && vasyaAge > katyaAge {
		high = mishaAge
		middle = vasyaAge
		low = katyaAge
		fmt.Println("hight - ", high)
		fmt.Println("middle - ", middle)
		fmt.Println("low - ", low)
	} else if mishaAge > vasyaAge && mishaAge > katyaAge && vasyaAge < katyaAge {
		high = mishaAge
		middle = katyaAge
		low = vasyaAge
		fmt.Println("hight - ", high)
		fmt.Println("middle - ", middle)
		fmt.Println("low - ", low)
	} else if vasyaAge > mishaAge && vasyaAge > katyaAge && mishaAge > katyaAge {
		high = vasyaAge
		middle = mishaAge
		low = katyaAge
		fmt.Println("hight - ", high)
		fmt.Println("middle - ", middle)
		fmt.Println("low - ", low)
	} else if vasyaAge > mishaAge && vasyaAge > katyaAge && mishaAge < katyaAge {
		high = vasyaAge
		middle = katyaAge
		low = mishaAge
		fmt.Println("hight - ", high)
		fmt.Println("middle - ", middle)
		fmt.Println("low - ", low)
	} else if katyaAge > mishaAge && katyaAge > vasyaAge && mishaAge > vasyaAge {
		high = katyaAge
		middle = mishaAge
		low = vasyaAge
		fmt.Println("hight - ", high)
		fmt.Println("middle - ", middle)
		fmt.Println("low - ", low)
	} else {
		high = katyaAge
		middle = vasyaAge
		low = mishaAge
		fmt.Println("hight - ", high)
		fmt.Println("middle - ", middle)
		fmt.Println("low - ", low)
	}
}
