package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello welcome to slices in go")
	var fruitList = []string{"Apple", "Banana", "Cherry"}
	fmt.Printf("Type of fruitList: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Orange")
	fmt.Println("Fruit List:", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 789
	highScores[3] = 101
	highScores = append(highScores, 500) // Adding a new score
	// highScores[4] = 500
	fmt.Println("High Scores:", highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted High Scores:", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//remove a value from slices based on index

	var courses = []string{"React", "Angular", "Vue", "JavaScript", "Python"}
	fmt.Println("Courses:", courses)
	var index int = 2
	courses= append(courses[:index], courses[index+1:]...)
	 fmt.Println("Updated Courses:", courses)

}
