package main

import (
	"fmt"
	"os"
)

func SayItRight(number int) string {
	mod := number % 10

	switch mod {
	case 1:
		return fmt.Sprintf("%d компьютер", number)
	case 2, 3, 4:
		return fmt.Sprintf("%d компьютера", number)
	default:
		return fmt.Sprintf("%d компьютеров", number)
	}
}

func main() {
	var number int

	fmt.Print("Введите число: ")
	fmt.Fscan(os.Stdin, &number)

	result := SayItRight(number)
	fmt.Println(result)
}
