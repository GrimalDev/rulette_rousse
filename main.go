package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello test")
		time.Sleep(time.Second * 3)
	}
}
