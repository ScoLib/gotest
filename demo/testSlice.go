package demo

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, 0)
	for i := 0; i < dy; i++ {
		t := make([]uint8, 0)
		for j := 0; j < dx; j++ {
			t = append(t, uint8((i+j)/3))
		}
		r = append(r, t)
	}
	return r
}

func PrintPic() {
	pic.Show(Pic)
}
