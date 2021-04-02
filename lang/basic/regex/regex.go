package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is 6022816661@qq.com
email1 is bbbb@mazi.com.cn
email2 is mazi0210@gmail.com
`

func main() {
	//re, err := regexp.Compile(text)
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)

	//if err != nil {
	//	panic(err)
	//}
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		for _, s := range m[1:] {
			fmt.Printf("%s ", s)
		}
		fmt.Println()
	}
}


