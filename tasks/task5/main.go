package main

import (
	"fmt"
	"sync"
)

func glitchRemover(name string, wg *sync.WaitGroup) {
	var out string
	//fmt.Printf("Coming : %s\n", name)
	for _, r := range name {

		//fmt.Println(string(r))
		if r > 58 {
			out += string(r)
		}

	}

	fmt.Printf("%s,", out)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	arr := []string{"gopher123", "alpha99beta", "1cita2del3"}

	for _, val := range arr {

		wg.Add(1)
		go glitchRemover(val, &wg)
		fmt.Printf("After goroutine: %s\n", val)

	}

	wg.Wait()

	fmt.Printf("\n")

}
