package main

import "fmt"

func Ints(vs ...int) int {
	return ints(vs)
}

func ints(vs []int) int {
	if len(vs) == 0 {
		return 0
	}

	return ints(vs[1:]) + vs[0]
}

func main() {
	out := Ints(1, 2, 3)
	fmt.Println(out)
}
