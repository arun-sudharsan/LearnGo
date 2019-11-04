package main

import "fmt"

func main() {

	a := 10
	b := 3

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	fmt.Printf("%v\n", a+b)
	fmt.Printf("%v\n", a-b)
	fmt.Printf("%v\n", a*b)
	fmt.Printf("%v\n", a/b)
	fmt.Printf("%v\n", a%b)

	var normalInt int = 5
	var abnormalInt uint16 = 5
	//complaints
	//fmt.Printf("%v\n", normalInt+abnormalInt)
	//mismatched types.

	fmt.Printf("%v\n", normalInt+int(abnormalInt))

	/*
		a = 10  binary = 1 0 1 0
		b =  3  bianry = 0 0 1 1
				 a & b = 0 0 1 0
				 a | b = 1 0 1 1
				 a ^ b = 1 0 0 1

	*/
	fmt.Printf("%v\n", a&b)
	fmt.Printf("%v\n", a|b)
	fmt.Printf("%v\n", a^b)

	a = 8
	//left shift
	fmt.Printf("%v\n", a<<3) //multiply by 2^3
	//right shift
	fmt.Printf("%v\n", a>>3) // divide by 2^3

	n := 3.14
	n = 1e10
	n = 0.4E14
	fmt.Printf("%v\n", n)

}
