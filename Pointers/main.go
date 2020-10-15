package main

import "fmt"

type person struct {
	name string
	age  int
}

func changeMe(p *person) {
	p.name = "Rahul"
	(*p).age = 30
}

func main() {
	p := person{
		name: "Piyush",
		age:  29,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
