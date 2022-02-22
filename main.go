package main

import (
	"fmt"
)

func main() {

	fmt.Print("Enter your name: ")
	var input string
	fmt.Scanf("%s", &input)

	i := input
	fizzbuzz(i)

}

//func fizzbuzz(i int) {
//
//	if i < 18 {
//		fmt.Printf("%v is not a proper age. You are not allowed to visit the park.", i)
//	} else if i >= 18 && i < 25 {
//		fmt.Println("You can enjoy visiting the park but entrance fee for you is $150")
//	} else if i >= 25 && i < 50 {
//		fmt.Println("You can enjoy visiting the park but the entrance fee for you is $250")
//	} else {
//		fmt.Println("You are too old to enter the park, Sorry")
//	}
//}

func fizzbuzz(i string) {
	length := len(i)
	if length > 10 {
		fmt.Println("You are a great person :)")
	} else if length < 10 {
		fmt.Println("You are an amazing person :)")
	}
}
