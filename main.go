package main

import (
	"fmt"
)

type User struct {
	Id  int
	Pwd int
}

func init() {
	fmt.Println("init被调用了")
}
func main() {
	fmt.Println("main被调用了")
	var a int = 22
	var user *int = &a
	fmt.Println("user:", *user)
}
