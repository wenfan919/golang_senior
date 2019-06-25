package file

import (
	"fmt"
	"os"
	"bufio"
)

var (
	firstName, lastName, s string
	i int
	f float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)

func Test() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s!\n", lastName) // Hi Chris Naegels
	fmt.Printf("Hi %s!\n", firstName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go
}

var inputReader *bufio.Reader
var in string
var err error

func Test2() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	in, err = inputReader.ReadString('o')
	if err == nil {
		fmt.Printf("The input was: %s\n", in)
	}
}


func Test3() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
	// For Unix: test with delimiter "\n", for Windows: test with "\r\n"
	switch input {
	case "Philip\r\n":  fmt.Println("Welcome Philip!")
	case "Chris\n":   fmt.Println("Welcome Chris!")
	case "Ivo\r\n":     fmt.Println("Welcome Ivo!")
	default: fmt.Printf("You are not welcome here! Goodbye!")
	}

	// version 2:
	switch input {
	case "Philip\r\n":  fallthrough
	case "Ivo\r\n":     fallthrough
	case "Chris\r\n":   fmt.Printf("Welcome %s\n", input)
	default: fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "Philip\r\n", "Ivo\r\n":   fmt.Printf("Welcome %s\n", input)
	default: fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}