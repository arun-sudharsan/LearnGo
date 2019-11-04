package main

import "fmt"

func main() {

	a := 10.4
	b := 3.2

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	fmt.Printf("%v\n", a+b)
	fmt.Printf("%v\n", a-b)
	fmt.Printf("%v\n", a*b)
	fmt.Printf("%v\n", a/b)
	// % not defined for float
	//fmt.Printf("%v\n", a%b)

	var n complex64 = 1 + 2i
	fmt.Printf("%v , %T\n", n, n)
	fmt.Printf("%v , %T\n", real(n), real(n))
	fmt.Printf("%v , %T\n", imag(n), imag(n))

	var s string = "This is a string"
	fmt.Printf("%v , %T\n", s, s)

}
