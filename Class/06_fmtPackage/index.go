package main

import "fmt"

var y = 42

func main()  {


fmt.Println(y)
// TYPE
fmt.Printf("%T\n", y) 
// BINARY
fmt.Printf("%b\n", y)
// hexadecimal
fmt.Printf("%x\n", y)
// zero x
fmt.Printf("%#x\n", y)
y = 911
fmt.Printf("%#x\n", y) 

s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)

fmt.Println(s)

	
}

