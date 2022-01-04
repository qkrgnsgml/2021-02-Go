package main

import "fmt"

type soldier struct {
	name string
	id   int
}

func main() {
	s1 := new(soldier)
	s1.name = "sda"
	fmt.Println(*s1)
}
