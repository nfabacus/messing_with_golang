package main

import (
	"fmt"
	"runtime"
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

	fmt.Println("\n##################")

	value1 := 42
	value2 := 24.356
	fmt.Println(value1)
	fmt.Printf("%T\n", value1)
	fmt.Println(value2)
	fmt.Printf("%T\n", value2)

	fmt.Println("\n##################")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Println("\n##################")
	s := "Hello, world!!!"
	s2 := `Yo,
		friends!
	`
	s = "Hi, Bob!"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)

	bs := []byte(s) // convert string to bytestring

	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	fmt.Println("\n##################")
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		//fmt.Println(s[i])
		fmt.Printf("%#U ", s[i])
	}

}

func foo() {
	fmt.Println("I am in foo")
}
