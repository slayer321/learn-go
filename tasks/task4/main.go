package main

import (
	"fmt"
)

//func numericFree(alphaNumeric)

func glitchRemover(name string, c chan string) {
	var out string
	//fmt.Printf("Coming : %s\n", name)
	for _, r := range name {

		//fmt.Println(string(r))
		if r > 58 {
			out += string(r)
		}

	}
	c <- out
	//fmt.Printf("Finish %s\n", out)

}

func main() {

	//val := "1cita2del3"
	arr := []string{"gopher123", "alpha99beta", "1cita2del3"}
	c := make(chan string, len(arr))

	for _, val := range arr {

		//fmt.Printf("In goroutine: %s\n", val)
		go glitchRemover(val, c)

	}

	// var res []string

	// val1 := <-c
	// val2 := <-c
	// val3 := <-c

	// for v := range c {
	// 	res = append(res, v)
	// }
	// res = append(res, val1, val2, val3)

	// fmt.Println(res)
	//close(c)
	for val := range c {
		fmt.Println(val)

	}

}
