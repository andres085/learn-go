package main

import "fmt"

func main() {
	var integer8 int8 = 127;
	var uinteger8 uint8 = 255;

	fmt.Printf("%v of type %T \n", integer8, integer8)
	fmt.Printf("%v of type %T \n", uinteger8, uinteger8)
}
