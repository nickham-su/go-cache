package main

import (
	"fmt"
	"github.com/nickham-su/go-cache"
)

func main() {
	c := cache.New()
	c.Set("test", 1)
	fmt.Println(c.GetFloat("test"))
	c.Delete("test")
	fmt.Println(c.GetFloat("test"))

	c.Set("testStr", "hello world")
	fmt.Println(c.GetString("testStr"))
}
