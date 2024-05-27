package demo

import "fmt"

func add(x, y int) int {
	return x + y
}
func Demo() {
	result := add(4, 7)
	fmt.Println("result:", result)
}
