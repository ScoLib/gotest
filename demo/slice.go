package demo

import (
	"fmt"
)

func Slice() {
	/*a := []int{3, 2, 3, 4}
	t, v := interface{}(a).([]int)
	fmt.Print(t, a, v)*/
	//s := []struct {
	//	i, v string
	//}{
	//	{"s", "afff"},
	//	{"add", "wweg"},
	//	{"aff", "gvrr"},
	//	{"era", "veeww"},
	//	{"vvawd", "afdfe"},
	//	{"vbee", "afeefb"},
	//}
	//fmt.Print(s)

	var b []int
	fmt.Println(b, len(b), cap(b))
	if b == nil {
		fmt.Print("b is nil\n")
	}

	c := make([]int, 0, 5)
	printSlice("c", c)

	c = c[:cap(c)]
	printSlice("c", c)

	c = c[1:]
	printSlice("c", c)
	c = append(c, 1)
	printSlice("c", c)

	pls := [][]string{
		{"C", "C++", "JAVA"},
		{"PHP", "Python", "JAVA"},
	}
	fmt.Printf("%s len=%d cap=%d %v\n", "pls", len(pls), cap(pls), pls)

	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	needc := countries[:len(countries)-2]
	needcpy := make([]string, len(needc))
	copy(needcpy, needc)
	fmt.Printf("%s len=%d cap=%d %v\n", "needcpy", len(needcpy), cap(needcpy), needcpy)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
