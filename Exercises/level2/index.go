package main

import "fmt"

// 1. Use var to Declare three variables. The variables should have package level scope. Do not assign Values to the variables. Use the followinf IDENTIFIERS for The variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE)
// a. identifier "x" type int
// b. identifier "y" type string
// c. identifier "z" type bool

// 2. in func main
// 	a. print out the values for each identifier
// 	b. The compiler assigned values to the variables. What are these values called?

var x int
var y string
var z bool

func main()  {

	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", x)

	fmt.Printf("%T\n", y)
	fmt.Printf("%v\n", y)

	fmt.Printf("%T\n", z)
	fmt.Printf("%v\n", z)

	//all values is 0
}