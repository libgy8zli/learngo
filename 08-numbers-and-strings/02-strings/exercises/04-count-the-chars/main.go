// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Count the Chars
//
//  1. Change the following program to work with unicode
//     characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func main() {
	// Currently it returns 7
	// Because, it counts the bytes...
	// It should count the runes (codepoints) instead.
	//
	// When you run it with "İNANÇ", it should return 5 not 7.
	var a = 'a'
	fmt.Printf("%d is type %[1]T\n", a)
	var b byte = 'b'
	fmt.Printf("%d is type %[1]T\n", b)
	fmt.Printf("%c ascii code is %d\n", b, b)
	fmt.Printf("%c unicode is %U\n", b, b)
	length_of_name := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length_of_name)
	fmt.Println(len(os.Args[1]))

}

