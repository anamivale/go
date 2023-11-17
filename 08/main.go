package main

import "fmt"

func main() {
	s := "abc"
	str := ""

	for _, c := range s {
		switch {
		case (c >= 'a' && c <= 'z'):
			c = ('z'-c)%26 + 'a'
			str += string(c)
		case (c >= 'A' && c <= 'Z'):
			c = ('Z'-c)%26 + 'A'
			str += string(c)
		}
	}
	fmt.Println(str)
}
