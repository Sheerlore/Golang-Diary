package main

import "fmt"

func YearsUntilRetirement(age int) {
	fmt.Println(100 - age)
}

func main() {
	age := new(int)
	*age = 26

	YearsUntilRetirement(age)
}