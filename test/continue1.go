package test

import "fmt"

func Test4() {
	//var a int = 1
	var b int = 1
	for a := 1; a < 3; a++ {
		for b := 1; b < 3; b++ {
			if a == 2 && b == 2 {
				continue
			}
			fmt.Printf("value of i and j is %d  %d\n", a, b)
		}
		fmt.Printf("value of i and j is %d  %d\n", a, b)

	}
}
