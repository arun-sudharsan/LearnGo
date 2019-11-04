package main

import "fmt"

func callME() {
	fmt.Println("Hello world!")
}

func main() {
	var i int //declare but no init at this time
	i = 1     //initialisation happens after some time.
	fmt.Printf("%v %T\n", i, i)

	var j = 10. //start with some value
	fmt.Printf("%v %T\n", j, j)
	//either int or float64
	k := 100.
	fmt.Printf("%v %T\n", k, k)
}
