package main

import "fmt"

func StrLen(str string) int {
	// count := 0
	// for i := 0; i < len(str); i++ {
	// 	count++
	// }
	return len(str)
}
func main() {
	str := "Hello World!"
	nb := StrLen(str)
	fmt.Println(nb)
}
