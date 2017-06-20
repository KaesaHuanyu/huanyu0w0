package main

import (
	"github.com/rainycape/memcache"
	"fmt"
)

func main() {
	mc,_ := memcache.New("192.168.1.42:60050")

	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	it, _ := mc.Get("foo")
	fmt.Printf("Get memcache %s successed", string(it.Value))
}