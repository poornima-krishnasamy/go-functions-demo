package main

import "fmt"

type PersonFactory struct{}

type Person struct {
	Name string
	Age  int
}

// `NewPerson` is now a method of a `PersonFactory` struct
func (p *PersonFactory) NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// `isAdult` is now a function where we're passing the `Person` as an argument
// instead of a receiver
func isAdult(p *Person) bool {
	return p.Age > 18
}

func modifyAdult(p *Person) int {
	p.Age = 25
	return p.Age
}

func main() {

	p1 := Person{}

	fmt.Println(isAdult(&p1))
	modifyAdult(&p1)
	fmt.Println(p1.Age)

	// object to create a new instance of person
	factory := &PersonFactory{}

	p2 := factory.NewPerson("John", 21)

	fmt.Println(isAdult(p2))
	modifyAdult(p2)
	fmt.Println(p2.Age)

}
