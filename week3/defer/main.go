package main

import "fmt"

func main() {
	 defer fmt.Println("world")
	 defer fmt.Println("one") 
	   defer fmt.Println("two")
	 
	  myDefer()
	  fmt.Println("hello")
	
}

func myDefer(){
	for i := 0; i < 5; i++ {
	defer fmt.Println(i)
  }
}