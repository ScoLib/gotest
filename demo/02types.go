package demo

import (
	"fmt"
	"unsafe"
)

func PrintTypes() {
	// bool
	a, b := true, false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	// 整型
	var i8 int8 = -128 // -128～127
	fmt.Printf("type %T var %d size %d", i8, i8, unsafe.Sizeof(i8))
	var i16 int16 = -32768 // -32768～32767
	fmt.Printf("\ntype %T var %d size %d", i16, i16, unsafe.Sizeof(i16))

	// rune 是 int32 的别名
	var i32 int32 = -2147483648 // -2147483648～2147483647
	fmt.Printf("\ntype %T var %d size %d", i32, i32, unsafe.Sizeof(i32))
	var i64 int64 = -9223372036854775808 // -9223372036854775808～9223372036854775807
	fmt.Printf("\ntype %T var %d size %d", i64, i64, unsafe.Sizeof(i64))

	// int：根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型
	i := 2147483648
	fmt.Printf("\ntype %T var %d size %d", i, i, unsafe.Sizeof(i))
	i = -9223372036854775808
	fmt.Printf("\ntype %T var %d size %d", i, i, unsafe.Sizeof(i))

	// 无符号整型
	// uint8： 0～255
	// byte 是 uint8 的别名

	// uint16： 0～65535
	// uint32： 0～4294967295
	// uint64： 0～18446744073709551615
	// uint：根据不同的底层平台，表示 32 或 64 位无符号整型。

	f1, f2 := 5.67, 8.97
	fmt.Printf("\ntype of f1 %T f2 %T\n", f1, f2)

	// 复数类型
	// complex64：实部和虚部都是 float32 类型的的复数。
	// complex128：实部和虚部都是 float64 类型的的复数。
	c1 := complex(5.6, 7.4)
	fmt.Printf("\ntype of c1 %T var %v\n", c1, c1)
	c2 := 8 + 9i
	fmt.Printf("type of c2 %T var %v\n", c2, c2)
	cadd := c1 + c2
	fmt.Println("cadd is", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	// string 类型
	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("My name is", name)

	// 类型转换
	j := 5
	k := 6.1
	h := float64(j) + k
	fmt.Println("h is", h)
}
