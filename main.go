package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Rating struct {
	Rating   int
	Comment  string
	Customer string
}

func (r *Rating) GetStars() string {
	return strings.Repeat("⭐", r.Rating)
}

func (r *Rating) CollectInput() {
	fmt.Println("Enter your name:")
	fmt.Scanln(&r.Customer)
	fmt.Println("Enter your rating:")
	fmt.Scanln(&r.Rating)
	fmt.Println("Enter your comment:")
	fmt.Scanln(&r.Comment)
}

func (r *Rating) Display() {
	stars := r.GetStars()
	if r.Rating >= 3 {
		color.Green("Thank you %s for your rating %s and your comment: %s", r.Customer, stars, r.Comment)
	} else {
		color.Red("Sorry to hear about the rating %s, we will improve our service.", stars)
	}
}

func main() {
	rating := &Rating{}
	rating.CollectInput()

	if rating.Rating < 1 || rating.Rating > 5 {
		color.Red("Invalid rating. Please enter a rating between 1 and 5.")
		return
	}

	rating.Display()
}
