package main

import (
	"fmt"
	"time"
)


func main(){
	
	go hello("my boy")
	hello("my girl")

	fmt.Println("lets explore go routines")

	
}

func hello(s string) {
	// time.Sleep(1 * time.Second)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("hello "  + s)
}