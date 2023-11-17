package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args[1]
	b := os.Args[2]

	if len(os.Args) != 3{
		return
	}
	if a <b{
		fmt.Print(-1)
	} else if a >b{
		fmt.Print(1)
	}else if a ==b {
		fmt.Print(0)
	}
}