package main

import (
	"fmt"
	"github.com/punzrr/progression/arithmetic"
	"github.com/punzrr/progression/geometric"
)

func main() {
	val, err := arithmetic.GetVal(10, 0, 5)
	fmt.Println(val, err)
	val, err = geometric.GetVal(-4, 4, 7)
	fmt.Println(val, err)
}
