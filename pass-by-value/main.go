package main

import "fmt"

func addA(a int, b int) int {
	return a + b
}

func addB(a, b int) int {
	c := a + b
	return c
}

func addC(a, b int) (c int) {
	c = a + b
	return
}

func Sum(values ...int) (sum int) {
	sum = 0
	for _, v := range values {
		sum += v
	}
	return
}

func main() {
	var a, b int = 1, 2

	// diff ways to call and define functions
	resA := addA(a, b)
	resB := addB(a, b)
	_ = addC(1, 2)

	// vardiac functions
	// a0 := Sum()
	// a1 := Sum(2)
	// a3 := Sum(2, 3, 5)

	fmt.Println(resA, resB) // 3
	//fmt.Println(a0, a1, a3) // // 0 2 10
}
