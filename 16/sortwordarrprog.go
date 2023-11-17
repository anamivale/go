package main

import "fmt"

func main() {
	names := []string{"a", "b", "1", "A", "B"}
	l := len(names)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1; j++ {
			if names[j] > names[j+1] {
				names[j], names[j+1] = names[j+1], names[j]
			}
		}
	}
	fmt.Println(names)
}
