package main

import "fmt"

func main() {

	var normalIntUnspecified int = 42 // unspecified represents platform dependent
	fmt.Printf("%v, %T\n", normalIntUnspecified, normalIntUnspecified)

	var smallInt int8 = 42 // 8 bit
	fmt.Printf("%v, %T\n", smallInt, smallInt)

	var mediumInt int16 = 42 // 16bit
	fmt.Printf("%v, %T\n", mediumInt, mediumInt)

	var regular int32 = 42 // 32 bit
	fmt.Printf("%v, %T\n", regular, regular)

	//bigger than this? go for big package from math library

	// we also have unsigned ints. add prefix u.

	var UnormalIntUnspecified uint = 42 // unspecified represents platform dependent
	fmt.Printf("%v, %T\n", UnormalIntUnspecified, UnormalIntUnspecified)

	var UsmallInt uint8 = 42 // 8 bit
	fmt.Printf("%v, %T\n", UsmallInt, UsmallInt)

	var UmediumInt uint16 = 42 // 16bit
	fmt.Printf("%v, %T\n", UmediumInt, UmediumInt)

	var Uregular uint32 = 42 // 32 bit
	fmt.Printf("%v, %T\n", Uregular, Uregular)
}
