package main

import (
	"fmt"
)

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Println("Hello, World!")
	foo()
	fmt.Println("Another println")

	//for i :=0; i <10; i++ {
	//	if i % 2 == 0 {
	//		fmt.Println(i)
	//	}
	//}
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}

func foo() {
	fmt.Println("I am in foo")
}
