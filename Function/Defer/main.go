package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("main ran")
}

func foo() {
	defer func() {
		fmt.Println("Foo defer ran")
	}()
	fmt.Println("Foo ran")
}
