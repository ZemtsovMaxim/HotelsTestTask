package main

import (
	"fmt"
	"os"
)

func PrintMultiplicationTable(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}
}
func main() {

	var number int
	fmt.Print("Введите число: ")
	fmt.Fscan(os.Stdin, &number)

	PrintMultiplicationTable(number)

}
