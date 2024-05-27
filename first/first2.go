package first

import (
	"fmt"
	"reflect"
)

func First2() {
	var a string = "HELLO WORLD"
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
}
