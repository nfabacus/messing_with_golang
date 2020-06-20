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

	a := 0
	for {
		a++
		if a > 10 {
			break
		}
		if a%2 != 0 {
			continue
		}
		fmt.Println(a)

	}
	fmt.Println("Done!")

	// if else if
	value2 := 41
	if value2 < 40 {
		fmt.Println("value is less than 40")
	} else if value2 == 40 {
		fmt.Println("Value is 40")
	} else {
		fmt.Println("Value is more than 40")
	}

	// switch statement
	switch { // missing expression will default to true
	//case false:
	//	fmt.Println("print if false")
	case false:
		fmt.Println("print if false")

	case true:
		fmt.Println("Print if true")
	}

	myVal := 7
	switch myVal {
	case 7:
		fmt.Println("value is 7")
		fallthrough
	case 5: // in this case even value is not 5, it prints! because it did fallthrough from above.
		fmt.Println("value is 5")
		//fallthrough
	default:
		fmt.Println("default!")
	}

	userName := "John"
	switch userName {
	case "Tony", "John":
		fmt.Println("Your name is", userName)
	default:
		fmt.Println("default name Ben!")
	}

}
