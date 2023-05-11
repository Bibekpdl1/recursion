package main

import (
	"fmt"

	"github.com/mukezhz/recursion/array"
	"github.com/mukezhz/recursion/sort"
	str "github.com/mukezhz/recursion/string"
	"github.com/mukezhz/recursion/toh"
)

func main() {
	// TOH is a function to solve the Tower of Hanoi problem
	toh.Toh(3, 'A', 'C', 'B')
	randomArr := []int{1, 3, 2, 4, 5}
	fmt.Println(randomArr)
	fmt.Println(array.CheckSorted(randomArr))
	fmt.Println(array.CheckSortedBetter(randomArr, 0))
	sortedArr := sort.MergeSort(randomArr)
	fmt.Println(sortedArr)
	randomArr = []int{1, 3, 2, 8, 5}
	fmt.Println(randomArr)
	sortedArr = sort.QuickSort(randomArr, 0, 5)
	fmt.Println(sortedArr)
	fmt.Println(array.CheckSorted(randomArr))
	fmt.Println(array.CheckSortedBetter(randomArr, 0))
	fmt.Println(str.ReplaceString("xpixpiyush"))
}
