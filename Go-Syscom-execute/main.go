package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute() {
	//here we perform the ks command.
	//we can use store the output of this in our out variabel
	//and catch any errors in err
	out, err := exec.Command("ls").Output()
	if err != nil {
		fmt.Println(err)
	}

	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)

	out, err = exec.Command("pwd").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Command Successfully Executed")
	output = string(out[:])
	fmt.Println(output)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}