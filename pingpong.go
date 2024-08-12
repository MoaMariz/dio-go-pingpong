package main

import (
	"fmt"
	"time"
)

func ping(canal chan string) {
	for {
		canal <- "ping"
	}
}

func pong(canal chan string) {
	for {
		canal <- "pong"
	}
}

func print(canal chan string) {
	for {
		msg := <-canal
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	var canal chan string = make(chan string)

	go ping(canal)
	go print(canal)
	go pong(canal)

	var enter string
	fmt.Scanln(&enter)

}
