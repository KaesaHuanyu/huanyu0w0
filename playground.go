package main

import (
	"fmt"
)

type Tree struct {
	Salary int
	LTree  *Tree
	RTree  *Tree
}

func main() {
	var a interface{}
	a = 2
	v, ok := a.(string)
	fmt.Println(v, ok)
}
