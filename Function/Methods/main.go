package main

import "fmt"

type person struct {
	first, last string
	age         int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}
func main() {
	p1 := person{
		first: "Rahul",
		last:  "Sachan",
		age:   30,
	}
	p1.speak()
}
