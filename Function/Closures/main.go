package main

import "fmt"

func main() {
	peopleCounter := getCounter()
	animalCounter := getCounter()

	fmt.Println("People Count", peopleCounter())
	fmt.Println("People Count", peopleCounter())
	fmt.Println("Animal Count", animalCounter())
	fmt.Println("Animal Count", animalCounter())
	fmt.Println("Animal Count", animalCounter())
	fmt.Println("Animal Count", animalCounter())
	fmt.Println("People Count", peopleCounter())
	fmt.Println("People Count", peopleCounter())
	fmt.Println("People Count", peopleCounter())
	fmt.Println("People Count", peopleCounter())
	fmt.Println("People Count", peopleCounter())
	fmt.Println("People Count", peopleCounter())
	fmt.Println("Animal Count", animalCounter())
	fmt.Println("Animal Count", animalCounter())
	fmt.Println("Animal Count", animalCounter())
}

func getCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
