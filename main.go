package main

import (
	"bufio"
	"fmt"
	"hello_world/ratings"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	productToReviewMap := make(map[string]*ratings.Rating)

	for {
		fmt.Println("Enter 1 to add product & review ,2 to view all products , 3 to exit:")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "1" {
			fmt.Println("Enter the product name:")
			productName, _ := reader.ReadString('\n')
			productName = strings.TrimSpace(productName)

			if productToReviewMap[productName] == nil {
				productToReviewMap[productName] = &ratings.Rating{}
			}

			// Sub-menu for adding reviews
			for {
				fmt.Println("Enter 1 to add review, 2 to main menu:")
				subChoice, _ := reader.ReadString('\n')
				subChoice = strings.TrimSpace(subChoice)

				if subChoice == "1" {
					fmt.Println("Enter the Customer name:")
					name, _ := reader.ReadString('\n')
					name = strings.TrimSpace(name)

					fmt.Println("Enter the rating (1-5):")
					ratingValue, _ := reader.ReadString('\n')
					ratingValue = strings.TrimSpace(ratingValue)

					fmt.Println("Enter the comment:")
					comment, _ := reader.ReadString('\n')
					comment = strings.TrimSpace(comment)

					ratingValueFloat, err := strconv.ParseFloat(ratingValue, 64)
					if err != nil {
						fmt.Println("Invalid rating")
						continue
					}

					err = productToReviewMap[productName].AddUserRating(name, ratingValueFloat, comment)
					if err != nil {
						fmt.Println(err)
					}
				} else if subChoice == "2" {
					break
				}
			}
		} else if choice == "2" {
			for product, rating := range productToReviewMap {
				fmt.Println("Product: ", product)
				fmt.Println("Rating: ", rating.String())
			}

		}else if choice == "3" {
			break
		}
	}
}
