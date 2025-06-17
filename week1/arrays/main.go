package main

import "fmt"

func main() {
	fmt.Println("welcome to arrays in golang")
	var fruitList [4]string
	fruitList[0]= "Apple"
	fruitList[1]= "Banana"
	fruitList[3]= "Cherry"
	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Length of Fruit List:", len(fruitList))
	var vegList= [5] string{"Carrot", "Broccoli", "Spinach"}
	fmt.Println("Vegetable List:", len(vegList))
}
