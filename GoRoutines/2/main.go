package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) printer() string {
	return fmt.Sprintf("%v is %v years old", (*p).name, (*p).age)
}

type human interface {
	printer() string
}

func main() {
	persons := []person{
		{"Rahul Sachan", 30},
		{"Isha Singh", 28},
	}

	for _, v := range persons {
		fmt.Println(v.printer())
	}

	for _, v := range persons {
		fmt.Println(human(&v).printer())
	}

	// for _, v := range persons {
	// 	//This will error out as methods attached to the pointer type are not part of the interface the type implements
	//  //This would work if this method was attached to the Type itself  p person not p *person
	// 	fmt.Println(human(v).printer())
	// }
}
