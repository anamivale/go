package main

import "fmt"

func main() {
	s := "valeria"
	str := ""
	for _, c := range s {
		switch {
		case (c >= 'a' && c <= 'z'):
			c = (c-'a'+1)%26 + 'a'
			str += string(c)
		case (c >= 'A' && c <= 'Z'):
			c = (c-'A'+1)%26 + 'A'
			str += string(c)
		}
	}
	fmt.Println(str)
}
