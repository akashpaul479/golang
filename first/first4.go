package first

import "fmt"

func First4() {
	s := []struct {
		i int
		b bool
	}{
		{1, true}, {2, true}, {3, false}, {4, true}, {5, false}, {6, true},
	}
	fmt.Println(s)
}
