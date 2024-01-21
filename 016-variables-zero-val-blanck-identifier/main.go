package main

import "fmt"

func main() {
	// a:=38

	// fmt.Println(a)

	// b, c, d, _, f := 0, 1, 2, 3, "happiness"
	// fmt.Println(b, c, d, f)

	// // b, c, d := 1, 2, 3
	// // fmt.Println(b, c)

	// var integer int
	// var string string
	// var boolean bool

	// fmt.Println(integer, string, boolean)

	// integer = 37
	// fmt.Println(integer)
	// martinez:= 38

	// fmt.Printf("38 as binary %b \n", martinez)
	// fmt.Printf("38 as hexadecimal %x \n", martinez)

	// 	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	// 	const String = "%v as binary and hex \t %b \t %X \n"

	// 	fmt.Printf(String, a, a, a)
	// 	fmt.Printf(String, b, b, b)
	// 	fmt.Printf(String, c, c, c)
	// 	fmt.Printf(String, d, d, d)
	// 	fmt.Printf(String, e, e, e)
	// 	fmt.Printf(String, f, f, f)

	z := 42.0
	var m float32 = 43.742

	// no se puede asignar en el lugar de una variable tipo float32  para tapar una valor float 32
	// esto no funciona
	// z = m

	// esto si funciona
	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)
}
