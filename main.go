package main

import "fmt"

func main() {
	a := 18

	hexChars := "0123456789abcdef"
	var changed []byte
	if a == 0 {
		fmt.Println(0)
	}

	for a > 0 {
		remainder := a % 16
		changed = append([]byte{hexChars[remainder]}, changed...)
		a = a / 16
	}
	fmt.Println(string(changed))
}
