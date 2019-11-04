package main

import "fmt"

//global variable needs to be mentioned with type.
var global float32 = 32

var (
	name    string = "Arun"
	age     int    = 22
	worksAt string = "Freshworks"
)

func main() {
	//local scope takes preference. Name Shadowing.
	var name = "UknowWho"
	//j declared but not used error
	//j:=10
	fmt.Printf("%v, %T\n", global, global)
	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", age, age)
	fmt.Printf("%v, %T\n", worksAt, worksAt)

}
