package main

import "fmt"

func FirstRune(s string) rune {
	x := []rune(s)
	return x[0]
}

func main() {
	x := FirstRune("valeria")
	fmt.Println(x)
}
