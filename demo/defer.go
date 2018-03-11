package demo

import "fmt"

func TDefer() {
	defer fmt.Print("world\n")

	fmt.Print("hello ")


	fmt.Print("counting\n")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("%v\n", i)
	}

	fmt.Print("done\n")
}
