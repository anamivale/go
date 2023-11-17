package main

import (
	"os"

	"github.com/01-edu/z01"
)

func mainx() {
	str1 := os.Args[1]
	str2 := os.Args[2]
	index := 0
	rewritten := ""
	for _, c := range str1 {
		for i := index; i < len(str2); i++ {
			if string(c) == string(str2[i]) {
				rewritten += string(c)
				index = i + 1
				break
			}
		}
	}
	for _, c := range rewritten {
		z01.PrintRune(c)
	}
}
