package main

import (
	"fmt"
)

type mamal struct {
	yearOffirstBreath int
}

type bird struct {
	yearOfhatching int
}

type human struct {
	yearOfBirth int
}

func (h human) getAge() int {
	return 2020 - h.yearOfBirth
}
func (b bird) getAge() int {
	return 2020 - b.yearOfhatching
}
func (m mamal) getAge() int {
	return 2020 - m.yearOffirstBreath
}

func (h human) makeSound() string {
	return "I am human"
}
func (b bird) makeSound() string {
	return "Cherp Cherp"
}
func (m mamal) makeSound() string {
	return "AAAGgggrrrrrrr!!!!"
}

type alive interface {
	getAge() int
	makeSound() string
}

func identify(a alive) {
	fmt.Println(a.getAge(), a.makeSound())
}

func main() {
	man := human{yearOfBirth: 1989}
	crow := bird{yearOfhatching: 2018}
	elephant := mamal{yearOffirstBreath: 2000}

	identify(man)
	identify(crow)
	identify(elephant)

}
