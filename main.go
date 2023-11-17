package main

import "github.com/01-edu/z01"

func main() {
	str := "i am a girl "

	lastspace := len(str) - 1

	for lastspace >= 0 && str[lastspace] == ' ' {
		lastspace--
	}

	end := lastspace

	for lastspace >= 0 && str[lastspace] != ' ' {
		lastspace--
	}
	if lastspace >= 0 {
		for i := lastspace + 1; i <= end; i++ {
			z01.PrintRune(rune(str[i]))
		}
	}
}
