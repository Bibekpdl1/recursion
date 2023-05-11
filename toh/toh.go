package toh

import "fmt"

// TOH is a function to solve the Tower of Hanoi problem
func Toh(n int, src rune, des rune, aux rune) {
	if n == 1 {
		fmt.Printf("Move peg 1 from %c to %c\n", src, des)
		return
	}
	Toh(n-1, src, aux, des)
	fmt.Printf("Move peg %d from %c to %c\n", n, src, aux)
	Toh(n-1, aux, des, src)
}
