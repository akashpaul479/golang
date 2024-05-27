package first

import "fmt"

func First3() {
	names := [4]string{
		"akash",
		"abhi",
		"kushal",
		"roshan",
	}
	fmt.Println(names)
	slice1 := names[1:3]
	slice2 := names[0:2]
	fmt.Println(slice1, slice2)
	slice2[0] = "bunty"
	fmt.Println(slice1, slice2)
	fmt.Println(names)
}
