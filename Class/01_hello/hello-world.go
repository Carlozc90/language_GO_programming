package main

import "fmt"
func main()  {
	fmt.Println("hello everyone, kasjdkasj")
	foo()
	fmt.Println("hello everyone3")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			
			fmt.Println(i)
		}
	}
}

func foo(){
	fmt.Println("hello everyone2")
}