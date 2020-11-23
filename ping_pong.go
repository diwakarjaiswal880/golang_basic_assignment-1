// go concurrency
package main

import (
	"fmt"
	"time"
)

func main() {
	go pingpong("Ping")
	pingpong("pong")
}
func pingpong(t string) {
	for i := 1; i < 20; i++ {
		fmt.Println(t)
		time.Sleep(time.Millisecond * 500)
	}
}
