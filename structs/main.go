package main

import (
	"fmt"
)

type pet struct {
	animal string
	age  int
}


type person struct {
	name string
	age  int
	pet  pet
}

func (p person) String() string {
	return p.name + " is " + fmt.Sprint(p.age) + " years old and has a pet " + p.pet.animal + " who is " + fmt.Sprint(p.pet.age) + " years old."
}

func (p person) updateName(newName string) {
	p.name = newName
}

func main() {
	person1 := person{name: "Alice", age: 30, pet: pet{animal: "dog", age: 5}}
	person2 := person{name: "Bob", age: 25, pet: pet{animal: "cat", age: 3}}
	
	fmt.Println(person1)
	fmt.Println(person2)

	person1.updateName("Alicia") // This does not change the name
	fmt.Println(person1)
}
