package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	nb, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	for i := 1; i <= 9; i++ {
		var result int
		result = i * nb
		fmt.Printf("%v x %v = %v \n", i, nb, result)
	}

	fmt.Print(nb)

}
