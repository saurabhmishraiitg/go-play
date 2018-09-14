package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello Go.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Who Are You ?")
	text, _ := reader.ReadString('\n')
	fmt.Println("Welcome:", text)
}
