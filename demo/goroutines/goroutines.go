package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		fmt.Printf("%c done! \n", r)
	}

	for i := 0; i < len(data); i++ {
		go capIt(data[i])
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Capitalized: %c\n", capitalized)

}
