package main

import (
	"fmt"
)

type polvilho int

var x polvilho

var y int

func main() {

	fmt.Printf("%v, %T\n", x, x)
	var x polvilho = 42
	fmt.Printf("%v\n", x)

	y = int(x)
	fmt.Printf("%v, %T", y, y)

}
