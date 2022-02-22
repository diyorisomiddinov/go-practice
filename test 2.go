package main

import "fmt"

type Contact struct {
	name string
	age  int
}

func (x Contact) welcome() {
	fmt.Printf("Hello %v. Are you %v years old?", x.name, x.age)

}

func main() {
	x := Contact{"Muhammaddiyor", 20}
	x.welcome()
}
