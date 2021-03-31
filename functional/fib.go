package main

import (
	"bufio"
	"fmt"
	"io"
	"learngo/functional/fib"
	"strings"
)

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	s := fmt.Sprintf("%d\n", next)
	if next > 10000 {
		return 0, io.EOF
	}

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()

	printFileContents(intGen(f))

	//fmt.Println(f()) // 1
	//fmt.Println(f()) // 1
	//fmt.Println(f()) // 2
	//fmt.Println(f()) // 3
	//fmt.Println(f()) // 5
	//fmt.Println(f()) // 8
	//fmt.Println(f()) // 13
	//fmt.Println(f()) // 21
	//fmt.Println(f()) // 55
}
