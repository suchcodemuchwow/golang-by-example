package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("function call direct")

	go f("create goroutine 1")
	go f("create goroutine 2")
	go func(msg string) { fmt.Println(msg) }("create goroutine inline")

	time.Sleep(time.Second / 10)
}

/*
order of goroutine stdouts change in each run
$ go run goroutines.go
function call direct : 0
function call direct : 1
function call direct : 2
create goroutine 1 : 0
create goroutine 1 : 1
create goroutine 1 : 2
create goroutine inline
create goroutine 2 : 0
create goroutine 2 : 1
create goroutine 2 : 2
*/
