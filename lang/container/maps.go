package main

import "fmt"

func main() {
	m := map[string]string {
		"name": "mazi",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int // m3 == nil
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	// map的key在map中无序
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("error key")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

	// 除了 slice map function 的内建类型都可以作为 key
	// struct 类型不包括上述字段，也可以作为 key
}
