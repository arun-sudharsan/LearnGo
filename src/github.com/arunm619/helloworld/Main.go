package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)
	var j float32 = 3.14
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var s string = "Arun"
	fmt.Printf("%v, %T\n", s, s)

	s = string(i) //string is just array of bytes.
	//converting 42 to string is ascii conversion which is *
	fmt.Printf("%v, %T\n", s, s)

	s = strconv.Itoa(i) //actual itoa method integer to string
	fmt.Printf("%v, %T\n", s, s)

}
