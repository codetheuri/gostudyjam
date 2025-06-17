package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on pointers")

	var ptr *int 
	x :=10
	ptr = &x // ptr now holds the address of x

}
