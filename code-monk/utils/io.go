package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
ReadFromStdin This function reads data from STDIN and prints to the STDOUT
*/
func ReadFromStdin() {
	var i int
	var x string
	fmt.Println("Print an integer to be scanned, then after enter Print a text to be scanned:")
	_, err := fmt.Scanf("%d\n%s", &i, &x)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Integer entered : ", i)
	fmt.Println("String entered : ", x)
}

/*
ReadFromStdinAlt This function reads data from STDIN using bufio library and prints to STDOUT
*/
func ReadFromStdinAlt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Who Are You ?")
	text, _ := reader.ReadString('\n')
	fmt.Println("Welcome:", text)
}
