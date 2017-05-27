package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Tree struct {
	Salary int
	LTree  *Tree
	RTree  *Tree
}

func main() {
	a := bson.NewObjectId()
	fmt.Println(a, a.Hex())
}
