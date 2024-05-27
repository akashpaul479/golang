package first

import (
	"fmt"
	"os"
)

func First6() {
	var s, Arg string
	for i := 0; i < len(os.Args); i++ {
		s += Arg + os.Args[i] + ""
	}
	fmt.Println(s)
}
