package main

import (
	"fmt"
	"os"
	"strconv"
)

func ConvertToHex(num int) string {
	const hexChars = "0123456789abcdef"

	// Handle the case when the number is 0 separately
	if num == 0 {
		return "0"
	}

	// Create a slice to store hexadecimal digits
	var hexDigits []byte

	// Convert the number to hexadecimal
	for num > 0 {
		remainder := num % 16
		hexDigits = append([]byte{hexChars[remainder]}, hexDigits...)
		num /= 16
	}

	return string(hexDigits)
}

func main() {
	// Check if the number of command-line arguments is not equal to 2
	if len(os.Args) != 2 {
		fmt.Println() // Print a newline and exit
		return
	}

	// Get the number from the command-line argument
	inputNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert the number to base 16
	hexNumber := ConvertToHex(inputNumber)

	// Display the result in base 16 followed by a newline
	fmt.Println(hexNumber)
}
