package first

import "fmt"

func First() {
	var a [5]int
	var i, j int
	for i = 0; i < 5; i++ {
		a[i] = i + 10
	}
	for j = 0; j < 5; j++ {
		fmt.Printf("elements:[%d]=%d\n", j, a[j])
	}
}
