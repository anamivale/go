package main

import "fmt"

func main()  {
	a := 20
	b:=17

	for b!=0{
		a, b = b, a%b
	}
	fmt.Printf("GCD is %v\n", a)
}