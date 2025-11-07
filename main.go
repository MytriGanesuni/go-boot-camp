package main

import "fmt"
import "github.com/fatih/color"

func main() {

	type rating struct {
		rating   int
		comment  string
		customer string
	}

	ratings := rating{}

	fmt.Println("Enter your name:")
	fmt.Scanln(&ratings.customer)

	fmt.Println("Enter your rating:")
	fmt.Scanln(&ratings.rating)

	fmt.Println("Enter your comment:")
	fmt.Scanln(&ratings.comment)

	stars := ""
	for i := 0; i < ratings.rating; i++ {
		stars += "⭐"
	}

	if(ratings.rating >=3 ){
		color.Green("Thank you ",ratings.customer," for your rating", stars, "and your comment", ratings.comment)
	} else {
		color.Red("Sorry to hear about the rating ", stars ," ,we will improve our service.")
	}

}
