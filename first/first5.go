package first

import (
	"fmt"
)

func First5() {
	slice1 := []int{2, 3, 4, 5, 6, 7}
	slice2 := slice1[:3]
	fmt.Println(slice2)
	slice3 := slice1[0:4]
	fmt.Println(slice3)
	slice4 := slice1[:5]
	fmt.Println(slice4)
	fmt.Println(slice1)
}
