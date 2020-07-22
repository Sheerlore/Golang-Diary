package main

import (
	"fmt"
	"time"
)

func backgroundTask() {
	ticker := time.NewTicker(1 * time.Second)

	for _ = range ticker.C {
		fmt.Println("tock")
	}
}

func main() {
	fmt.Println("Go ticker")

	go backgroundTask()

	fmt.Println("The rest of my application can continue")

	select {}
}