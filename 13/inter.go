package main

import "fmt"

func isnotDuplicate(c rune, str string) bool {
	for _, char := range str {
		if c == char {
			return false
		}
	}
	return true
}

func main() {
	a := "padinton"
	b := "paqefwtdjetyiytjneytjoeyjnejeyj"
	str := ""

	for _, ac := range a {
		if isnotDuplicate(ac, str) {
			for _, bc := range b {
				if bc == ac {
					str += string(ac)
					break
				}
			}
		}
	}
	fmt.Println(str)
}
