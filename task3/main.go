package main

import (
	"fmt"
	"math"
)

func FindPrimeNumbers(min, max int) []int {
	if max < min {
		panic("Max should be bigger than min")
	}

	primeNumbers := []int{}
	for i := min; i <= max; i++ {
		if i < 2 {
			continue
		}
		isPrime := true
		limit := int(math.Sqrt(float64(i)))
		for j := 2; j <= limit; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers
}

func main() {

	fmt.Println(FindPrimeNumbers(11, 20))

}
