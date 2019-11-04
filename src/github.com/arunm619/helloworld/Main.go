package main

import "fmt"

// incrementing counters..

const (
	c1 = iota
	c2
	c3
)

const (
	a2 = iota
)

/*const (
	c1 = iota
	c2 = iota
	c3 = iota
)*/

const (
	_ = iota
	okCoder
	superCoder
	legendaryCoder
)

func main() {

	const myConstant = 4
	//cannot modify constants
	//myConstant = 1
	fmt.Printf("%v, %T\n", myConstant, myConstant)

	//typed constants
	const a int = 1
	const b string = "b"
	const c float32 = 3.14
	const d bool = true

	fmt.Printf("%v, %T\n", c+myConstant, c+myConstant)

	fmt.Printf("%v, %T\n", c1, c1)
	fmt.Printf("%v, %T\n", c2, c2)
	fmt.Printf("%v, %T\n", c3, c3)

	fmt.Printf("%v, %T\n", a2, a2)

	var coder int = legendaryCoder
	fmt.Printf("%v", coder == legendaryCoder)

}

//typed constants
//untyped constants
//enumerated
