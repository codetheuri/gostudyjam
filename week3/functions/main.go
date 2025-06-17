package main

import "fmt"

func main() {
	  greeter()
   fmt.Println("welcome to functions in Go")
   result:= adder(3,5)
  fmt.Println("result is",result)


  proRes,message:= proAdder(1,2,3,4,-78)  
  fmt.Println("pro result is", proRes)
   fmt.Println("message from proAdder is: ", message)
 
  
}
func adder(valone int, valtwo int) int {
	return valone + valtwo
}
func proAdder(values ...int)(int,string) {
   total := 0
   for _, value := range values {
      total += value
   }
   return total, "Hi from proAdder"
}
func greeter() {
   fmt.Println("Hello from the greeter function")
}
