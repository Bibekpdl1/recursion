package main

import (
	"fmt"

	"github.com/mukezhz/recursion/array"
	"github.com/mukezhz/recursion/toh"
)

func main() {
	// TOH is a function to solve the Tower of Hanoi problem
	toh.Toh(3, 'A', 'C', 'B')
	randomArr := []int{1, 3, 2, 4, 5}
	fmt.Println(randomArr)
	fmt.Println(array.CheckSorted(randomArr))
	fmt.Println(array.CheckSortedBetter(randomArr, 0))
}
