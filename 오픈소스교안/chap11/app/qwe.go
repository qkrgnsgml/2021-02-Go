package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Id    int
	Score float64
	Class string
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Score < s[j].Score } //부등호 바뀌면
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := Student{"홍길동", 12345678, 3.91, "A"}
	fmt.Println(s)
}
