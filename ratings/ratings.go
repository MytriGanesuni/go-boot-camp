package ratings

import "time"
import "fmt"
import "errors"

type Rating struct {
	id          int
	userRatings []UserRatings
	avgRating   float64
}

type Comment struct {
	date time.Time
	text string
}

type UserRatings struct {
	customer string
	rating   float64
	comments []Comment
}

func (r *Rating) AddUserRating(customer string, rating float64, comment string) error {
	if rating < 0 || rating > 5 {
		return errors.New("rating must be between 0 and 5")
	}
	r.userRatings = append(r.userRatings, UserRatings{
		customer: customer,
		rating:   rating,
		comments: []Comment{
			{date: time.Now(),
				text: comment},
		},
	})
	r.avgRating = r.AverageRating()
	fmt.Println(r.String())
	return nil
}

func (receiver Comment) String() string {
	return fmt.Sprintf("Date : %s, Text: %s", receiver.date.Format("06-01-02 15:04:05"), receiver.text)
}
func (receiver UserRatings) String() string {
	return fmt.Sprintf("Customer: %s, Rating: %.2f, Comments: %v", receiver.customer, receiver.rating, receiver.comments)
}
func (receiver Rating) String() string {
	return fmt.Sprintf("User Ratings: %v, Average Rating: %.2f", receiver.userRatings, receiver.avgRating)
}

func (receiver Rating) AverageRating() float64 {

	totalRating := 0.0
	for _, userRating := range receiver.userRatings {
		totalRating += userRating.rating
	}
	return totalRating / float64(len(receiver.userRatings))
}
