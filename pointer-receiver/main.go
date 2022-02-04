package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// This function returns a new instance of `Person`
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// The `Person` pointer type is the receiver of the `isAdult` method
func (p *Person) isAdult() bool {
	return p.Age > 18
}

func (p *Person) modifyAdult() int {
	p.Age = 25
	return p.Age
}

func main() {

	p1 := Person{}
	fmt.Println(p1.isAdult())
	fmt.Println(p1.modifyAdult())

	// to create a new instance of person
	p2 := NewPerson("John", 21)

	fmt.Println(p2.isAdult())
	p2.modifyAdult()
	fmt.Println(p2.Age)

}
