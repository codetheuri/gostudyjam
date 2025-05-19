package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	//comma ok  //err o
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello",input)
	fmt.Printf("type of this inout is %T",input)
}
																							