package vilse

import (
	"fmt"
)

func Vilse() {
	number := 50
	guess := 3

	if guess < 1 {
		fmt.Print("below 1")
	} else if guess > 100 {
		fmt.Println("more than 100")
	} else {
		if guess < number {
			fmt.Println("low")
		}
		if guess > number {
			fmt.Println("high")
		}
		if guess == number {
			fmt.Println("crct")
		}
	}
}
