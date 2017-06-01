package main

import (
	"fmt"
	"huanyu0w0/model"
	"time"
)

type Tree struct {
	Salary int
	LTree  *Tree
	RTree  *Tree
}

func main() {
	a := &model.Article{
		Time: time.Now().Add(-34 * time.Minute),
	}

	fmt.Println(a.GetShowTime())
}
