package main

import (
	"fmt"
)

var z = true

func main() {
	x := 16
	y := "strings"
	z := true

	_, erros := fmt.Println(x, y, z)
	fmt.Println(erros)
}
