package main

import "fmt"

var jwtToken = 100 // private variable
// public variable
const JwtToken2 string = "effrfrwf" // public variable

func main() {
	var username string = "Gopher"
	var isloggedin bool = true
	var smallval int = 256
	var smallfloat float64 = 255.5324432435

	fmt.Println("hello " + username)
	fmt.Println(isloggedin)
	fmt.Println(smallval)
	fmt.Println(smallfloat)
	fmt.Printf("variable is type : %T \n", username)
	fmt.Printf("variable is type : %T \n", isloggedin)
	fmt.Printf("variable is type : %T \n", smallval)
	fmt.Printf("variable is type : %T \n", smallfloat)

	//default values and some aliases
	var a int
	fmt.Println(a)                             // 0
	fmt.Printf("variable is of type :%T\n", a) // int

	//imlicit type
	var website = "https://www.golang.org"
	fmt.Println(website) // https://www.golang.org

	//no var style
	numberOfUser := 100
	fmt.Println(numberOfUser) // 100

	fmt.Println("jwtToken is ", JwtToken2)
}
