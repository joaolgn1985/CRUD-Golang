package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Joao"
	fmt.Println("variable name is:", reflect.TypeOf(name))
	version := 1.1

	fmt.Println("Hello, Mr.", name)
	fmt.Println("This is program is in version", version)

	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	var comand int
	fmt.Scanf("%d", &comand)
	fmt.Println("Value of the command variable is:", comand)

	switch comand {
	case 1:
		fmt.Println("Starting monitoring...")
	case 2:
		fmt.Println("Opening logs...")
	case 0:
		fmt.Println("Exiting of the program...")
	default:
		fmt.Println("Please choose one the options above...")
	}
}
