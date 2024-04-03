package main

import (
	"fmt"
)

type polvilho int

var x polvilho

func main() {

	fmt.Printf("%v, %T\n", x, x)
	var x polvilho = 42
	fmt.Printf("%v", x)

}
