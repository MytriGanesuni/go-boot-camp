package main

import (
	"fmt"

	"hello_world/ratings"
)

func main() {

	rating := &ratings.Rating{}
	err := rating.AddUserRating("Customer1", 8.0, "Great product")
	if err != nil {
		fmt.Println(err)
	}
	err = rating.AddUserRating("Customer2", 4.0, "okay")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rating)

}
