package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Excuting my concurrent anonymous function")
	}()

	fmt.Scanln()
}