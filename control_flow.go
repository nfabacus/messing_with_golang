package main

import "fmt"

func main() {
	//data := 14
	//fmt.Printf("Variable string %d content\n", data)

	for i := 0; i <= 10; i++ {
		response := fmt.Sprintf("****** Outer %d *******", i)
		fmt.Println(response)
		for j := 0; j <= 3; j++ {
			fmt.Printf("The outer loop %d\t: The inner loop: %d\n", i, j)
		}
	}

	// below is the Go equivalent of while statement
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")

	z := 1
	for {
		if z > 9 {
			break
		}
		fmt.Println(z)
		z++
	}
	fmt.Println("done.")

	value1 := 85 / float64(80)
	fmt.Println(value1)

}
