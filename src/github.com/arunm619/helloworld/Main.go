package main

import "fmt"

func main() {

	var s string = "This is a string"
	fmt.Printf("%v , %T\n", s, s)

	//prints 104 because string is array of bytes.
	fmt.Printf("%v , %T\n", s[1], s[1])

	//convert to string...
	fmt.Printf("%v , %T\n", string(s[1]), string(s[1]))

	//s[2] = 100
	//s[2] = 'a'
	//s[2] = "nothing can go here as it is immutable"
	fmt.Printf("%v , %T\n", s, s)

	var s2 = "Another String"

	fmt.Printf("%v , %T\n", s+s2, s+s2)

	b := []byte(s)
	fmt.Printf("%v , %T\n", b, b)

	c := 'a'
	fmt.Printf("%v , %T\n", c, c)

}
