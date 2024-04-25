package main

import (
	"fmt"
	"math"
	"sort"
)

func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func FindCommonDivisors(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	gcd := nums[0]
	for _, num := range nums[1:] {
		gcd = findGCD(gcd, num)
	}

	divisors := []int{}
	for i := 1; i <= int(math.Sqrt(float64(gcd))); i++ {
		if gcd%i == 0 {
			divisors = append(divisors, i)
			if i != gcd/i {
				divisors = append(divisors, gcd/i)
			}
		}
	}

	sort.Ints(divisors)

	return divisors
}

func main() {
	nums := []int{42, 12, 18}
	commonDivisors := FindCommonDivisors(nums)
	fmt.Println("Общие делители:", commonDivisors) // Выводит: Общие делители: [1 2 3 6]
}
