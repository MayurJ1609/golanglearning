package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// No inheritance in golang; No super or parent
	mayur := User{"Mayur", "mayur@go.dev", true, 31}
	fmt.Println(mayur)
	fmt.Printf("Mayur details are: %+v\n ", mayur)
	fmt.Printf("Name is %v and email is %v", mayur.Name, mayur.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
