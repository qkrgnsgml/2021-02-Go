package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 20; i++ {
		fmt.Print("a")
		time.Sleep(time.Second)
	}
}

func b() {
	for i := 0; i < 20; i++ {
		fmt.Print("b")
		time.Sleep(time.Second)
	}
}

func c() {
	fmt.Println("c")
	time.Sleep(time.Second * 2)
}

func main() {
	go a()
	go b()
	defer c()
	time.Sleep(time.Second * 2)
	fmt.Println("\n메인함수끝")
}
