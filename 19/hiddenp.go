package main

import "fmt"

func notdup(chr rune, str string) bool {
	for _, c := range str {
		if c == chr {
			return false
		}
	}
	return true
}

func main() {
	a := "abc"
	b := "2albtrcb53.sse"
	output := ""
	index := 0
	for _, c := range a {
		for i := index; i < len(b); i++ {
			if c == rune(b[i]) {
				output += string(c)
				index = i + 1
				break

			}
		}
	}
	// fmt.Print(output)
	if output == a {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
