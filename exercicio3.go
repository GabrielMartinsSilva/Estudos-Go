package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

func main() {

	s := fmt.Sprint(x, y, z)

	//fmt.Println(x)
	fmt.Println(s)
}
