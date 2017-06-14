package main

import "fmt"

func doStuff(x int) {

	fmt.Printf("Value of argument x: %v \n", x)
	fmt.Printf("Memory address of x: %v \n", &x)

	x = x * 5

	fmt.Printf("Value of argument x: %v \n", x)
	fmt.Printf("Memory address of x: %v \n", &x)
}

func doStuffWithPointer(x *int) {
	fmt.Printf("Value of argument x: %v \n", x)
	fmt.Printf("Memory address of x: %v \n", &x)
	fmt.Printf("Dereferenced value of x: %v \n", *x)

	*x = *x * 5

	fmt.Printf("Value of argument x: %v \n", x)
	fmt.Printf("Memory address of x: %v \n", &x)
	fmt.Printf("Dereferenced value of x: %v \n", *x)
}

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Printf("Memory address %v \n", b)
	fmt.Printf("Dereferenced Value %v \n", *b)

	*b = 42
	fmt.Println(a)

	doStuff(a)

	fmt.Println(a)

	doStuffWithPointer(&a)

	fmt.Println(a)
}
