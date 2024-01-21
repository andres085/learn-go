package main

import "fmt"

func main() {
	decimal, float, hexadecimal := 747, 911, 90210

	fmt.Printf("%d \t\t %b \t\t %#X \n", decimal, decimal, decimal)
	fmt.Printf("%d \t\t %b \t\t %#X \n", float, float, float)
	fmt.Printf("%d \t\t %b \t %#X \n", hexadecimal, hexadecimal, hexadecimal)
}
