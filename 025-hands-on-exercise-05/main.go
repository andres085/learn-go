package main

import "fmt"

func main() {
	var zero int
	one, two, three, _, four := 1, 2, 3, "blank", 4
	var name string = "Andres"

	fmt.Println(zero)
	fmt.Println(one, two, three)
	fmt.Println(name)
	fmt.Println(one, two, three, four)
}
