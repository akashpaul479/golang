package demo

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
func Demo6() {
	num := 7
	result := factorial(num)
	fmt.Printf("factorial of %d is %d\n", num, result)
}
