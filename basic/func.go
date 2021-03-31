package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := div(a, b)
		return q, nil
	default:
		//return 0, errors.New("unsupported operation: " + op)
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	//return a / b, a % b
	q, r = a / b, a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling func %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for idx, val := range numbers {
		fmt.Println(idx, val)
		s += val
	}
	return s
}

//func swap(a, b *int) {
//	*b, *a = *a, *b
//}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	if result, err := eval(12, 4, "//"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(func (a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
