package main

import "fmt"

/*
  i/p: "abcde"
  o/p: "cdabe"

  i/p: "abc"
  o/p: "abc"

  i/p: "abcdefghi"
  o/p: "cdabghefi"
...
*/

func swapConsecutive(s string) string {
	if len(s) < 4 {
		return s
	}
	swapped := fmt.Sprintf("%s%s%s%s", string(s[2]), string(s[3]), string(s[0]), string(s[1]))
	small_output := swapConsecutive(s[4:])
	return fmt.Sprintf("%s%s", swapped, small_output)
}

func main() {
	var s string
	fmt.Print("Enter a string: ")
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Printf("Error!!! %v", err)
	}
	fmt.Println(string(swapConsecutive(s)))
}
