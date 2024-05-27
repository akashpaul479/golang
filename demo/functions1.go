package demo

import "fmt"

func sub(x, y int) int {
	return x - y
}
func Demo1() {
	a := sub(1, 1)
	fmt.Println("Result:", a)
}
