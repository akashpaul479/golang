package test

import "fmt"

func Test2() {
	var a int = 1
	for a < 10 {
		fmt.Print("value of a is ", a, "\n")
		a++
		if a == 5 {
			break
		}
	}
}
