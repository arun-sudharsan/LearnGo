package main

import "fmt"

func main() {

	var isLearning bool = true
	fmt.Printf("%v, %T\n", isLearning, isLearning)
	//false
	var whatIsDefaultBoolean bool
	fmt.Printf("%v, %T\n", whatIsDefaultBoolean, whatIsDefaultBoolean)
	isSame := 1 == 1
	fmt.Printf("%v, %T\n", isSame, isSame)

}
