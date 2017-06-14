package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Printf("Memory address %v \n", b)
	fmt.Printf("Dereferenced Value %v \n", *b)

	*b = 42
	fmt.Println(a)
}
