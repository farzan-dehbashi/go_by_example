package main

import (
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
)

func main() {
	const s = "فرزان دهباشی هستم"
	fmt.Println(s, len(s))

	for i:=0; i<len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println()
	fmt.Printf("%x", s[0])
	fmt.Println()
	var x rune = '1'
	fmt.Println(x)
	fmt.Println(utf8.RuneCountInString("123"))
	runes := []rune{'H', 'e', 'l', 'l', 'o'}
	fmt.Println(runes)
	fmt.Println(utf16.Encode(runes))


    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %v\n", runeValue, idx)
    }

	var y rune = '1'
	if y == '1' {
		fmt.Println("x is equal")
	}

}
