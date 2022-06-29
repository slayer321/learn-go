package main

import (
	"fmt"
	"time"
)

func Alphabet(start, end int) {

	allAlphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	requiredAlphabet := allAlphabet[start-1 : end]

	fmt.Println(requiredAlphabet)
}

func main() {

	go Alphabet(1, 5)
	go Alphabet(4, 8)
	go Alphabet(2, 3)

	time.Sleep(time.Second)
}

// The reason alphabets are not in order is because all the 3 goroutines are not
// in sync and any one that is getting finish first is getting printed
