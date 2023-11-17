package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str := os.Args[1]

	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	for _, cha := range str {
		switch {
		case (cha >= 'a' && cha <= 'z'):
			rep := int(cha-'a') + 1
			for i := 0; i < rep; i++ {
				z01.PrintRune(cha)
			}
		case (cha >= 'A' && cha <= 'Z'):
			rep := int(cha-'A') + 1
			for i := 0; i < rep; i++ {
				z01.PrintRune(cha)
			}
		}

	}
}
