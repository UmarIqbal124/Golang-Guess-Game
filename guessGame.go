package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// for creating random number
	secret := getRandomNumber()

	for matching := false; !matching; {

		// for user input
		guess := getUserInput()
		// for matching user input and random number
		matching = isMatching(secret, guess)
		fmt.Println(matching)
	}

}

func isMatching(secret, guess int) bool {

	if guess > secret {

		fmt.Println("Your guess is too big ")
		return false
	} else if guess < secret {
		fmt.Println("You guess is too small")
		return false
	} else {
		fmt.Println("***Congratulation You got it***")
		fmt.Println("Game over!")
		time.Sleep(5 * time.Second)
		return true
	}

}

func getRandomNumber() int {

	// to generate random number
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}

func getUserInput() int {

	var input int
	fmt.Print("Please enter your Guess : ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Wrong Input!")
		getUserInput()
	} else {

		fmt.Println("You guessed : ", input)

	}
	return input

}
