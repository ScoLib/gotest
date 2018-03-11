package demo

import "fmt"

type Vert struct {
	Lat, Long float64
}

func Map() {
	var p map[string]int

	if p == nil {
		fmt.Print("map is nil\n")
		p = make(map[string]int)
		fmt.Print(p)
	}

	m := make(map[string]Vert)
	m["Bob"] = Vert{
		34.2452, 23.12,
	}

	fmt.Println(m["Bob"])
}
