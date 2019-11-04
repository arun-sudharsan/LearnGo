package main

import "fmt"

func main() {
	var i int = 4
	fmt.Printf("%v, %T\n", i, i)
	var j float32 = 3.14
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

}
