package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const maxNumGuesses int = 5

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var (
		guess string
		n     int
		err   error
	)
	found := false
	numberToGuess := r1.Intn(100)

	fmt.Println("Pick a number between 0 and 100")
	fmt.Println("You get a maximum of", maxNumGuesses, "guesses")

	for i := 0; i < maxNumGuesses; i++ {
		fmt.Print("Guess #", i+1, " : ")
		n, err = fmt.Scan(&guess)
		if err == nil && n > 0 {
			g, _ := strconv.Atoi(guess)
			if g > numberToGuess {
				fmt.Println("You're guess was too high!")
			} else if g < numberToGuess {
				fmt.Println("You're guess was too low!")
			} else {
				fmt.Println("TIME TO CELEBRATE!!!! AWESOME JOB!!!! You correctly guessed the number.")
				found = true
				break
			}
		} else {
			fmt.Println("Huh, I do not recognize what you typed, goodbye")
			break
		}
	}

	if !found {
		fmt.Println("Too bad that you are terrible at guessing.")
		fmt.Println("The correct number was ", numberToGuess)
	}

}
