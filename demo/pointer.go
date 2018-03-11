package demo

import "fmt"

func Pointer() {
	i, j := 40, 3244
	p := &i
	fmt.Print(*p)
	fmt.Print("\n")
	*p = 21
	fmt.Print(i)
	fmt.Print("\n")

	p = &j
	*p = *p / 2
	fmt.Print(j)
}
