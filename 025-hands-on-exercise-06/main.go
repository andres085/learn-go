package main

import "fmt"

func main() {
	text, number, money := "Andres", 38, 9.9

	fmt.Printf("%v of type %T \n", text, text)
	fmt.Printf("%v of type %T \n", number, number)
	fmt.Printf("%v of type %T \n", money, money)
}
