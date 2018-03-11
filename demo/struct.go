package demo

import "fmt"

type Vertex struct {
	X int
	Y int
}

func TStruct() {
	fmt.Print(Vertex{1, 2})
}