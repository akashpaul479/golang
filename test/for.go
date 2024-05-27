package test

import "fmt"

func Test() {
	var a int
loop:
	fmt.Print("you are not eligible for vote\n")
	fmt.Print("Enter your age:")
	fmt.Scanln(&a)
	if a <= 18 {
		goto loop
	} else {
		fmt.Print("you are eligible")
	}

}
