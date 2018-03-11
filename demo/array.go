package demo

import "fmt"

func Array() {
	/*var a [2]string
	a[0] = "hello"
	a[1] = "world"

	fmt.Print(a[0], a[1])
	fmt.Print(a)

	p := [...]int{2, 3, 2, 4}
	fmt.Print(p)

	s := p[1:3]
	fmt.Print(s)*/

	names := [4]string{
		"jack",
		"lucy",
		"Bob",
		"John",
	}
	/*aa := names[1:3]
	b := names[2:]
	fmt.Print(aa, b)

	b[0] = "xxxxx"

	fmt.Print(aa, b)
	fmt.Print(names)*/
	for i, v := range names {
		fmt.Printf("%d of %s\n", i, v)
	}

}
