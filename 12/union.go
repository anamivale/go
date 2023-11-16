package main

import "fmt"

func isNotDuplicate(char rune, str string) bool {
	for _, c := range str {
		if c == char {
			return false
		}
	}
	return true
}

func main() {
	a := "valeria"
	b := "muhembele"

	str := ""
	for _, c := range a {
		if isNotDuplicate(c, str) {
			str = str + string(c)
		}
	}
	for _, c := range b {
		if isNotDuplicate(c, str) {
			str = str + string(c)
		}
	}
	fmt.Println(str)
}
