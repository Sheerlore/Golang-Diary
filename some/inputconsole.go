package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Simple shell")
	//fmt.Println("---------------------------")

	//for {
	//	fmt.Println("->")
	//	text, _ := reader.ReadString('\n')

	//	//convert CRLF to LF
	//	text = strings.Replace(text, "\n" , "", -1)

	//	if strings.Compare("hi", text) == 0 {
	//		fmt.Println("Hello Yourself")
	//	}
	//}
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(char)

	switch char {
	case 'A':
		fmt.Println("A key pressed")
		break
	case 'a':
		fmt.Println("a key pressed")
		break
	}
}