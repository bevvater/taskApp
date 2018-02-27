package main

import (
	"fmt"
)

//var	i int = 10


type Sample struct {
	Name   string `default:"John Smith"`
	Age    int    `default:"27"`
}

func main() {

//	fmt.Println(i)

	v := &Sample{}
	fmt.Println(v)
}