package main

import (
	"os"
	"fmt"
)

type Tree struct {
	Salary int
	LTree *Tree
	RTree *Tree
}

func main() {
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}