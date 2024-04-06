package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 1
	return &p
}

func main() {
	fmt.Println(person{"Zartosht", -2500})
	fmt.Println(person{name: "Anosh", age: 12})

	//zero will be printed
	fmt.Println(person{name: "Zar"})

	fmt.Println(*newPerson("Saad"))
	instance := person{"sample", 12}
	fmt.Println(instance.age)

	sp := &instance
	fmt.Println(sp.age)
}
