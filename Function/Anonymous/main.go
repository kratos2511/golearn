package main

import "fmt"

func main() {
	x := func(x, y int) string {
		return fmt.Sprintf("%v + %v = %v", x, y, x+y)
	}(2, 3)

	fmt.Println(x)
}
