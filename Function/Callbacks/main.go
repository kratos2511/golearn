package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Odd Sum", oddSum(adder, numbers...))
	fmt.Println("Even Sum", evenSum(adder, numbers...))
}

func adder(xi ...int) int {
	sum := 0
	for _, number := range xi {
		sum += number
	}
	return sum
}

func oddSum(f func(...int) int, numbers ...int) int {
	oddNumbers := make([]int, 0, len(numbers))
	for _, number := range numbers {
		if number%2 != 0 {
			oddNumbers = append(oddNumbers, number)
		}
	}
	return f(oddNumbers...)
}

func evenSum(f func(...int) int, numbers ...int) int {
	evenNumbers := make([]int, 0, len(numbers))
	for _, number := range numbers {
		if number%2 == 0 {
			evenNumbers = append(evenNumbers, number)
		}
	}
	return f(evenNumbers...)
}
