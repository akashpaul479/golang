package test

import "fmt"

func Test3() {
	var a int = 1
	for a < 10 {
		if a == 4 {
			a = a + 1
			continue
		}
		fmt.Print("value of a is:", a, "\n")
		a++
	}

}
