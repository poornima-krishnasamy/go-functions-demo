package main

import "fmt"

// Composite types
func modifySlice(a []string) []string {
	// declaring and initialing inside the func
	//a = make([]string, 0)
	a[1] = "mice"
	return a
}

// Primitive types
func ModifyBasicTypesPointer(name *string, age *int) {
	*name = "Golang"
	*age = 2022
	return
}

// Primitive types
func ModifyBasicTypes(name string, age int) {
	name = "Golang"
	age = 2022
	return
}

func main() {

	var name string = "ruby"
	var age int = 2021

	fmt.Println("Before function call: ", name, age)
	ModifyBasicTypes(name, age)
	fmt.Println("Original values: ", name, age)

	ModifyBasicTypesPointer(&name, &age)
	fmt.Println("Original values after function call: ", name, age)

	animals := []string{
		"cat", "dog", "fish",
	}
	// animals := make([]string, 3)
	// animals := make([]string, 0)

	fmt.Println("\nBefore function call: ", animals)

	fmt.Println("Function call:", modifySlice(animals))

	fmt.Println("After function call: ", animals)

}
