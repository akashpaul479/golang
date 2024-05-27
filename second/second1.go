package second

import (
	"fmt"
	"strings"
)

func Second1() {
	str := "india"
	fmt.Println(strings.ToUpper(str))
}

func Second3() {
	str := "INDIA"
	fmt.Println(strings.ToLower(str))
}

func Second4() {
	s := "INDIA"
	fmt.Println(strings.HasPrefix(s, "IN"))
}

func Second5() {
	s := "INDIA"
	fmt.Println(strings.HasSuffix(s, "IA"))
}

func Second6() {
	var arr = []string{"a", "b", "c"}
	fmt.Println(strings.Join(arr, "*"))
}

func Second7() {
	var str = "NEW  "
	fmt.Println(strings.Repeat(str, 4))
}
func Second8() {
	str := "AKASH"
	fmt.Println(strings.Replace(str, "A", "a", 2))
}
