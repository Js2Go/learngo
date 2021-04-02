package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//var aa = 3
//var ss = "kkk"
//var bb = true

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, false, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter()  {
	a, b, c, s := 3, 4, false, "def"
	fmt.Println(a, b, c, s)
}

func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f %.3f\n",
		cmplx.Exp(1i * math.Pi) + 1,
		cmplx.Pow(math.E, 1i * math.Pi) + 1,
	)
}

func triangle() {
	var a, b = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}

func consts() {
	//常量的数值可以作为各种类型使用
	const (
		filename string = "abc.txt"
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename, c)
}

func enums() {
	//const (
	//	cpp = 0
	//	java = 1
	//	python = 2
	//	golang = 3
	//)
	const (
		cpp = iota + 1
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(aa, bb, ss)
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
	consts()
	enums()
}
