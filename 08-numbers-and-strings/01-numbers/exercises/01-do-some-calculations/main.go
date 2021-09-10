// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Do Some Calculations
//
//  1. Print the sum of 50 and 25
//  2. Print the difference of 50 and 15.5
//  3. Print the product of 50 and 0.5
//  4. Print the quotient of 50 and 0.5
//  5. Print the remainder of 25 and 3
//  6. Print the negation of `5 + 2`
//
// EXPECTED OUTPUT
//  75
//  34.5
//  25
//  100
//  1
//  -7
// ---------------------------------------------------------

func main() {
	fmt.Printf("The sum of 50 and 25 is %d.\n", 50+25)
	fmt.Printf("The difference of 50 and 15.5 is %0.1f.\n", 50.0-15.5)
	fmt.Printf("The product of 50 and 0.5 is %0.1f.\n", 50 * 0.5)
	fmt.Printf("The quotient of 50 and 0.5 is %0.1f.\n", 50 / 0.5)
	fmt.Printf("The remainder of 25 and 3 is %d.\n", 25 % 3)
	fmt.Printf("The negation of `5+2` is %d.\n", -(5+2))
}
