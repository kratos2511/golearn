package main

import "fmt"

func main() {
	adder := bar("s")
	fmt.Printf("%T\n", adder)
	fmt.Println(bar("s")(2, 3))
	fmt.Println(bar("p")(2, 3))
	fmt.Println(bar("q")(3, 2))
	fmt.Println(bar("r")(3, 2))

}

func bar(s string) func(int, int) string {
	switch s {
	default:
		return func(x, y int) string {
			return fmt.Sprintf("%v + %v = %v", x, y, x+y)
		}
	case "p":
		return func(x, y int) string {
			return fmt.Sprintf("%v * %v = %v", x, y, x*y)
		}
	case "q":
		return func(x, y int) string {
			return fmt.Sprintf("%v / %v = %v", x, y, x/y)
		}
	case "r":
		return func(x, y int) string {
			return fmt.Sprintf("%v %% %v = %v", x, y, x%y)
		}
	}
}
