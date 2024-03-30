package main

import (
	"fmt"
)

func main() {
	x := 100

	fmt.Printf("x: %v, %T\n", x, x)

	x, z := 20, 30

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
}
