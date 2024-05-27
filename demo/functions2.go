package demo

import "fmt"

type akash struct {
	fname string
	lname string
}

func (a akash) fullname() {
	fmt.Println(a.fname + "" + a.lname)
}
func Demo3() {
	result := akash{"akash", "paul"}
	result.fullname()
}
