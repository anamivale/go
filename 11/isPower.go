package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	x := 1
	flag := false
	for i := 1; i < 15; i++ {
		x = x * 2
		if x == num {
			flag = true
		}
	}
	if flag {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
