package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//var name string

	fmt.Print("\tWelcome\n\n")
	fmt.Println("What is your name?")
	fmt.Print("--> ")

	//------Input String with white spaces--------//

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Println()

	fmt.Printf("Hello, %s   :)\n\n", name) // USE Printf for formatting
	//fmt.Println("Hello, " + name + "  :)")
	//fmt.Println("\nHello, ", name)

}
