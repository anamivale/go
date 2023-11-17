// package main

// func split(s string, sep string ) []string{

// }

package main

import (
	"fmt"
	"os"
)

func SplitWhiteSpaces(s string) []string {
	var result []string
	var word string
	for i, c := range s {
		if c != ' ' && c != '\t' && c != '\n' {
			word += string(c)
		}
		if ((c == 'H' || i == len(s)-1)) {
			result = append(result, word)
			word = ""
		}
	}
	return result
}

func main() {
	if len(os.Args)!=2{
		return
	}
	s := os.Args[1]
	x := SplitWhiteSpaces(s)
	fmt.Print(x)
}
