package main

import "fmt"

func main() {
	// Create a new Person instance
	person1 := User{
		"John Doe",
		
		"jojn2@gmail.com",
		true,
		30,
	}
	fmt.Println("Person 1:", person1)
	fmt.Println("Name:", person1.Name)
	 person1.GetStatus()
	  person1.NewMail()
	  fmt.Println("Person 1:", person1)
}

type User struct {
	Name string

	Email  string
	Status bool
	Age    int
}
func (u User) GetStatus(){
 fmt.Println("is user Active:", u.Status)
}
func (u User)NewMail() {
u.Email = "hello@gmail.com"
fmt.Println("New Email:", u.Email)
}
